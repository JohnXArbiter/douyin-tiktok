package comment

import (
	"context"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsByVideoIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentsByVideoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsByVideoIdLogic {
	return &ListCommentsByVideoIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentsByVideoIdLogic) ListCommentsByVideoId(req *types.VideoIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
