package relation

import (
	"context"
	"douyin-tiktok/common/utils"
	"errors"
	"strconv"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowedUsersByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowedUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowedUsersByUserIdLogic {
	return &ListFollowedUsersByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowedUsersByUserIdLogic) ListFollowedUsersByUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId    = req.UserId
		userIdStr = strconv.FormatInt(userId, 10)
		key       = utils.UserFollow + userIdStr
	)

	relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
	userInfos := relationCommonLogic.ListFollowedUsersOrFans(userId, 1, key)
	if userInfos == nil {
		return nil, errors.New("出错啦")
	}

	resp := utils.GenOkResp()
	resp["user_list"] = userInfos
	return resp, nil
}
