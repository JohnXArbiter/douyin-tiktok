package publish

import (
	"context"
	__user "douyin-tiktok/service/user/cmd/rpc/types"
	"errors"
	"sync"

	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"douyin-tiktok/service/file/cmd/api/internal/types"

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

// deprecated
func (l *ListPublishedUserIdLogic) ListPublishedUserId(req *types.UserIdReq) error {
	var wg sync.WaitGroup

	// TODO
	userIdStr := "123"
	var userId int64 = 123

	var userRpcResp *__user.GetInfoByUserIdResp
	wg.Add(1)
	go func() {
		wg.Done()

		req := &__user.UserIdReq{UserId: 0}
		userRpcResp, err := l.svcCtx.UserRpc.GetInfoByUserId(l.ctx, req)
		if err != nil {
			logx.Errorf("[RPC ERROR] ListPublishedUserId 获取用户信息失败 %v\n", err)
			userRpcResp.Code = -1
		}
	}()
	wg.Wait()
	if userRpcResp.Code == -1 {
		return errors.New("数据获取失败！")
	}

	return nil
}
