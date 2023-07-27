package relation

import (
	"context"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowedUserByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowedUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowedUserByUserIdLogic {
	return &ListFollowedUserByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowedUserByUserIdLogic) ListFollowedUserByUserId(req *types.UserIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
