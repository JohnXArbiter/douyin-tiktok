package logic

import (
	"context"

	"douyin-tiktok/service/user/cmd/rpc/internal/svc"
	"douyin-tiktok/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoByUserIdLogic {
	return &GetInfoByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoByUserIdLogic) GetInfoByUserId(in *__.UserIdReq) (*__.GetInfoByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetInfoByUserIdResp{}, nil
}
