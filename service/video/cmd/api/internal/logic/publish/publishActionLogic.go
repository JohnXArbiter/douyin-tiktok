package publish

import (
	"context"
	"douyin-tiktok/common/utils"
	__file "douyin-tiktok/service/file/cmd/rpc/types"
	"douyin-tiktok/service/video/model"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"io"
	"mime/multipart"
	"time"
	"xorm.io/xorm"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq, header *multipart.FileHeader) error {
	var (
		userId    = l.ctx.Value("user").(utils.JwtUser).Id
		respChan  = make(chan *__file.UploadVideoResp)
		videoId   = idgen.NextId()
		videoName = header.Filename
	)

	videoFile, err := header.Open()
	defer videoFile.Close()
	if err != nil {
		logx.Errorf("[SYS ERROR] PublishAction header open 失败 %v\n", err)
		return errors.New("视频读取失败！")
	}

	videoBytes, err := io.ReadAll(videoFile)
	if err != nil {
		logx.Errorf("[SYS ERROR] PublishAction 读取二进制数据失败 %v\n", err)
		return errors.New("视频读取失败！")
	}

	_, err = l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		var tx = session.Table("video_info")

		// 保存到 video_info 表
		videoInfo := &model.VideoInfo{
			Id:        videoId,
			UserId:    userId,
			CoverUrl:  "", // TODO 设置一个默认值
			Title:     req.Title,
			PublishAt: time.Now().Local(),
		}
		if _, err = tx.Insert(videoInfo); err != nil {
			return nil, errors.New("视频保存失败！")
		}

		// 请求文件服务 rpc 更新作品信息
		uploadVideoReq := &__file.UploadVideoReq{
			UserId:    userId,
			VideoName: videoName,
			VideoFile: videoBytes,
			VideoId:   videoId,
		}
		go l.saveVideoInfo(respChan, uploadVideoReq)

		uploadVideoResp := <-respChan // 阻塞等待
		if uploadVideoResp.GetCode() == -1 {
			return nil, errors.New("视频上传失败！")
		}

		fileVideoId := uploadVideoResp.GetFileVideoId()
		videoInfo.FileVideoId = fileVideoId
		videoInfo.PlayUrl = uploadVideoResp.GetPlayUrl()
		_, err = tx.Cols("file_video_id", "play_url").Update(videoInfo)
		if err == nil {
			return nil, nil
		}

		removeVideoReq := &__file.RemoveVideoReq{
			FileVideoId: fileVideoId,
			ObjectName:  uploadVideoResp.GetObjectName(),
		}
		go l.svcCtx.FileRpc.RemoveVideo(l.ctx, removeVideoReq) // 保证原子性，删除 oss 文件
		return nil, errors.New("视频上传失败！")
	})

	return err
}

func (l *PublishActionLogic) saveVideoInfo(dataChan chan *__file.UploadVideoResp, req *__file.UploadVideoReq) {
	var data = &__file.UploadVideoResp{Code: -1}
	defer func() {
		dataChan <- data
	}()

	resp, err := l.svcCtx.FileRpc.UploadVideo(l.ctx, req)
	if err != nil {
		logx.Errorf("[RPC ERROR] PublishAction 入新增视频失败 %v\n", err)
	} else {
		data = resp
	}
}
