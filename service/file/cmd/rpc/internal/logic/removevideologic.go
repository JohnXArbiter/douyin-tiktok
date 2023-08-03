package logic

import (
	"context"

	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
	"douyin-tiktok/service/file/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveVideoLogic {
	return &RemoveVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveVideoLogic) RemoveVideo(in *__.RemoveVideoReq) (*__.RemoveVideoResp, error) {
	// todo: add your logic here and delete this line

	return &__.RemoveVideoResp{}, nil
}
