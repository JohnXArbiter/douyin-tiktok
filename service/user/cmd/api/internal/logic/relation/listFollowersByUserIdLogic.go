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

type ListFollowersByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowerByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowersByUserIdLogic {
	return &ListFollowersByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowersByUserIdLogic) ListFollowersByUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId    = req.UserId
		userIdStr = strconv.FormatInt(userId, 10)
		key       = utils.UserFan + userIdStr
	)

	relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
	userInfos := relationCommonLogic.ListFollowedUsersOrFans(userId, 2, key)
	if userInfos == nil {
		return nil, errors.New("出错啦")
	}

	for _, userInfo := range userInfos {
		userInfo.IsFollow = true
	}

	resp := utils.GenOkResp()
	resp["user_list"] = userInfos
	return resp, nil
}
