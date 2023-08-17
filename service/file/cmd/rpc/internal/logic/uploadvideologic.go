package logic

import (
	"bytes"
	"context"
	"douyin-tiktok/service/file/cmd/rpc/internal/logic/oss"
	"douyin-tiktok/service/file/model"
	"strconv"
	"time"

	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
	"douyin-tiktok/service/file/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadVideoLogic) UploadVideo(in *__.UploadVideoReq) (*__.UploadVideoResp, error) {
	var (
		resp      = &__.UploadVideoResp{Code: -1}
		userId    = in.GetUserId()
		userIdStr = strconv.FormatInt(userId, 10)
		videoFile = bytes.NewReader(in.GetVideoFile())
		ossLogic  = oss.NewOssLogic(l.ctx, l.svcCtx)
	)

	url, objectName, err := ossLogic.UploadRaw(videoFile, in.GetVideoName(), userIdStr)
	if err != nil {
		return resp, err
	}

	fileVideo := &model.FileVideo{
		UserId:     userId,
		VideoId:    in.GetVideoId(),
		ObjectName: objectName,
		Url:        url,
		UploadAt:   time.Now().Unix(),
	}
	if _, err = l.svcCtx.FileVideo.Insert(fileVideo); err != nil {
		go ossLogic.Delete(objectName)
		logx.Errorf("[DB ERROR] UploadVideo 插入视频文件记录失败 %v\n", err)
		return resp, err
	}

	resp.Code = 0
	resp.ObjectName = objectName
	resp.PlayUrl = url
	return resp, nil
}
