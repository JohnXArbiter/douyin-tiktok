package publish

import (
	"context"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPublishedUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPublishedUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPublishedUserIdLogic {
	return &ListPublishedUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPublishedUserIdLogic) ListPublishedUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	// todo: add your logic here and delete this line

	return nil, nil
}
