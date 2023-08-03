package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &__.UploadVideoResp{}, nil
}
