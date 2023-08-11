package logic

import (
	"context"
	"douyin-tiktok/service/video/model"
	"time"

	"douyin-tiktok/service/video/cmd/rpc/internal/svc"
	"douyin-tiktok/service/video/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveVideoLogic {
	return &SaveVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveVideoLogic) SaveVideo(in *__.SaveVideoReq) (*__.SaveVideoResp, error) {
	var videoInfo = &model.VideoInfo{
		UserId:    in.UserId,
		PlayUrl:   in.Url,
		Title:     in.Title,
		PublishAt: time.Now().Unix(),
	}
	id, err := l.svcCtx.VideoInfo.Insert(videoInfo)
	if err != nil {
		logx.Errorf("[RPC ERROR] 视频服务：插入新增视频失败 %v\n", err)
		return &__.SaveVideoResp{Code: -1}, err
	}

	resp := &__.SaveVideoResp{VideoId: id}
	return resp, nil
}
