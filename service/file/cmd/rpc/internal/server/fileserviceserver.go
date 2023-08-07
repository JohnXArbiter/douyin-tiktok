// Code generated by goctl. DO NOT EDIT.
// Source: FILE.proto

package server

import (
	"context"

	"douyin-tiktok/service/file/cmd/rpc/internal/logic"
	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
	"douyin-tiktok/service/file/cmd/rpc/types"
)

type FileServiceServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedFileServiceServer
}

func NewFileServiceServer(svcCtx *svc.ServiceContext) *FileServiceServer {
	return &FileServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServiceServer) UploadVideo(ctx context.Context, in *__.UploadVideoReq) (*__.UploadVideoResp, error) {
	l := logic.NewUploadVideoLogic(ctx, s.svcCtx)
	return l.UploadVideo(in)
}

func (s *FileServiceServer) RemoveVideo(ctx context.Context, in *__.RemoveVideoReq) (*__.CodeResp, error) {
	l := logic.NewRemoveVideoLogic(ctx, s.svcCtx)
	return l.RemoveVideo(in)
}