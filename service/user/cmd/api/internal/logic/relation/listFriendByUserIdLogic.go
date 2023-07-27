package relation

import (
	"context"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFriendByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendByUserIdLogic {
	return &ListFriendByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFriendByUserIdLogic) ListFriendByUserId(req *types.UserIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
