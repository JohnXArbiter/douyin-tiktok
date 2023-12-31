// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"douyin-tiktok/service/user/cmd/rpc/internal/logic"
	"douyin-tiktok/service/user/cmd/rpc/internal/svc"
	"douyin-tiktok/service/user/cmd/rpc/types"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) GetInfoById(ctx context.Context, in *__.GetInfoByIdReq) (*__.GetInfoByIdResp, error) {
	l := logic.NewGetInfoByIdLogic(ctx, s.svcCtx)
	return l.GetInfoById(in)
}

func (s *UserServiceServer) GetInfoList(ctx context.Context, in *__.GetInfoListReq) (*__.GetInfoListResp, error) {
	l := logic.NewGetInfoListLogic(ctx, s.svcCtx)
	return l.GetInfoList(in)
}

func (s *UserServiceServer) UpdateWorkCnt(ctx context.Context, in *__.UpdateWorkCntReq) (*__.CodeResp, error) {
	l := logic.NewUpdateWorkCntLogic(ctx, s.svcCtx)
	return l.UpdateWorkCnt(in)
}
