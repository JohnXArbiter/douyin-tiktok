package publish

import (
	"bytes"
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/api/internal/logic/oss"
	"douyin-tiktok/service/video/model"
	"errors"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/yitter/idgenerator-go/idgen"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
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

const (
	VideoPlayUrl = iota
	VideoObjectName
	CoverUrl
	CoverObjectName
)

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq, header *multipart.FileHeader, loggedUser *utils.JwtUser) error {
	var (
		userId    = loggedUser.Id
		userIdStr = strconv.FormatInt(userId, 10)
		videoId   = idgen.NextId()
	)

	mi, err := l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		var tx = session.Table("video_info")

		m, err := l.saveVideo(header, userIdStr)
		if err != nil {
			logx.Errorf("[DB ERROR] PublishAction 视频上传失败 %v\n", err)
			return nil, errors.New("视频上传失败！")
		}

		// 保存到 video_info 表
		videoInfo := &model.VideoInfo{
			Id:        videoId,
			UserId:    userId,
			PlayUrl:   m[VideoPlayUrl],
			CoverUrl:  m[CoverUrl],
			Title:     req.Title,
			PublishAt: time.Now().Unix(),
		}
		if _, err = tx.Insert(videoInfo); err != nil {
			logx.Errorf("[DB ERROR] PublishAction 插入视频（作品）失败 %v\n", err)
			return nil, errors.New("视频保存失败！")
		}
		return m, nil
	})

	if err != nil {
		m := mi.(map[int]string)
		go l.removeVideoAndCover(m) // 保证原子性，删除 oss 文件
	}
	return err
}

func (l *PublishActionLogic) saveVideo(file *multipart.FileHeader, userId string) (map[int]string, error) {
	res := make(map[int]string)
	ossLogic := oss.NewOssLogic(l.ctx, l.svcCtx)
	videoPlayUrl, videoObjectName, err := ossLogic.Upload(file, userId)
	if err != nil {
		return res, err
	}
	res[VideoPlayUrl] = videoPlayUrl
	res[VideoObjectName] = videoObjectName

	coverFrame, err := l.getFrameByFfmpeg(videoPlayUrl)
	if err != nil {
		return res, err
	}

	videoFileName := file.Filename
	coverUrl, coverObjectName, err := ossLogic.UploadRaw(coverFrame, videoFileName[:strings.LastIndex(videoFileName, ".")]+".png", userId)
	if err != nil {
		return res, err
	}
	res[CoverUrl] = coverUrl
	res[CoverObjectName] = coverObjectName

	return res, nil
}

func (l *PublishActionLogic) getFrameByFfmpeg(videoUrl string) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		logx.Errorf("[FFMPEG ERROR] getFrameByFfmpeg 生成视频封面图 %v\n", err)
		return nil, err
	}
	return buf, nil
}

func (l *PublishActionLogic) removeVideoAndCover(m map[int]string) {
	ossLogic := oss.NewOssLogic(l.ctx, l.svcCtx)
	if v := m[VideoObjectName]; v != "" {
		objectName, _ := strings.CutPrefix(v, l.svcCtx.Oss.BaseUrl)
		ossLogic.Delete(objectName)
	}
	if v := m[CoverObjectName]; v != "" {
		objectName, _ := strings.CutPrefix(v, l.svcCtx.Oss.BaseUrl)
		ossLogic.Delete(objectName)
	}
}
