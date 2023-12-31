// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"douyin-tiktok/service/video/cmd/rpc/internal/logic"
	"douyin-tiktok/service/video/cmd/rpc/internal/svc"
	"douyin-tiktok/service/video/cmd/rpc/types"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServiceServer) GetFavoriteAndFavoritedCnt(ctx context.Context, in *__.GetFavoriteAndFavoritedCntReq) (*__.GetFavoriteAndFavoritedCntResp, error) {
	l := logic.NewGetFavoriteAndFavoritedCntLogic(ctx, s.svcCtx)
	return l.GetFavoriteAndFavoritedCnt(in)
}
