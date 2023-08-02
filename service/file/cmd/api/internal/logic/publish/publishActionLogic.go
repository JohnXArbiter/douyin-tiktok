package publish

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/file/cmd/api/internal/logic/oss"
	"douyin-tiktok/service/file/model"
	__video "douyin-tiktok/service/video/cmd/rpc/types"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"douyin-tiktok/service/file/cmd/api/internal/types"

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
		userIdStr = strconv.FormatInt(userId, 10)
		respChan  = make(chan *__video.SaveVideoResp)
	)

	ossLogic := oss.NewOssLogic(l.ctx, l.svcCtx)
	url, objectName, err := ossLogic.Upload(header, userIdStr)
	if err != nil {
		return errors.New("视频上传失败！")
	}

	// 请求视频（作品）rpc 服务更新作品
	videoRpcReq := &__video.SaveVideoReq{
		UserId: userId,
		Title:  req.Title,
		Url:    url,
	}
	go l.saveVideoInfo(respChan, videoRpcReq)

	videoRpcResp := <-respChan // 阻塞等待
	if videoRpcResp.Code == -1 {
		go ossLogic.Delete(objectName) // 保证原子性，删除 oss 文件
		return errors.New("视频上传失败！")
	}

	// 保存到 file_video 表
	videoId := videoRpcResp.VideoId
	fv := &model.FileVideo{
		UserId:     userId,
		VideoId:    videoId,
		Objectname: objectName,
		Url:        url,
		UploadAt:   time.Now().Local(),
	}
	_, err = l.svcCtx.FileVideo.Insert(fv)
	if err != nil {
		go ossLogic.Delete(objectName) // 保证原子性，删除 oss 文件
		return errors.New("视频保存失败！")
	}
	return nil
}

func (l *PublishActionLogic) saveVideoInfo(dataChan chan *__video.SaveVideoResp, req *__video.SaveVideoReq) {
	var resp, err = l.svcCtx.VideoRpc.SaveVideo(l.ctx, req)
	if err != nil {
		logx.Errorf("[RPC ERROR] PublishAction 入新增视频失败 %v\n", err)
		resp = &__video.SaveVideoResp{Code: -1}
	}

	dataChan <- resp
}
