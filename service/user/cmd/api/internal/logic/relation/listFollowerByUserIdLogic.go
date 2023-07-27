package relation

import (
	"context"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowerByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowerByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowerByUserIdLogic {
	return &ListFollowerByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowerByUserIdLogic) ListFollowerByUserId(req *types.UserIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
