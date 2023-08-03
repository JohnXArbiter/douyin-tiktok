package user

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId       = l.ctx.Value("user").(utils.JwtUser).Id
		targetUserId = req.UserId
	)

	userInfo := &model.UserInfo{Id: targetUserId}
	if has, err := l.svcCtx.UserInfo.Get(userInfo); err != nil || !has {
		return nil, errors.New("找不到该用户")
	}

	if userId != targetUserId {
		isFollow, err := l.svcCtx.UserRelation.Where("`user_id` = ? AND `to_user_id` = ?", userId, targetUserId).Exist()
		if err != nil {
			logx.Errorf("[DB ERROR] GetInfo 查询关注记录失败 %v\n", err)
		}
		userInfo.IsFollow = isFollow
	}

	resp := utils.GenOkResp()
	resp["user"] = userInfo
	return resp, nil
}
