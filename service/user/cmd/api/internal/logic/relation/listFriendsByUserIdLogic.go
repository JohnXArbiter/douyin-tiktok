package relation

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFriendsByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFriendsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFriendsByUserIdLogic {
	return &ListFriendsByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFriendsByUserIdLogic) ListFriendsByUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId    = req.UserId
		userIdStr = strconv.FormatInt(userId, 10)
		followKey = utils.UserFollow + userIdStr
		fanKey    = utils.UserFan + userIdStr
		resp      = utils.GenOkResp()
	)

	// 1.找关注列表
	if exist := l.svcCtx.Redis.Exists(l.ctx, followKey).Val(); exist == 0 {
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 1)
		if userRelation == nil && err == nil {
			resp["user_list"] = make([]struct{}, 0)
			return resp, nil
		}
		_, err = relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Followers, followKey)
		if err != nil {
			return nil, errors.New("出错啦")
		}
	}

	// 2.找粉丝列表
	if exist := l.svcCtx.Redis.Exists(l.ctx, fanKey).Val(); exist == 0 {
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 2)
		if userRelation == nil && err == nil {
			resp["user_list"] = make([]struct{}, 0)
			return resp, nil
		}
		_, err = relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Fans, fanKey)
		if err != nil {
			return nil, errors.New("出错啦")
		}
	}

	// 3.求交集
	store := &redis.ZStore{Keys: []string{followKey, fanKey}, Weights: []float64{1, 1}}
	idStrs, err := l.svcCtx.Redis.ZInter(l.ctx, store).Result()
	if err != nil {
		logx.Errorf("[REDIS ERROR] ListFriendsByUserId 获取用户:%v，关注粉丝交集失败 %v\n", userId, err)
		return nil, errors.New("出错啦")
	}

	// 4.取用户信息
	ids, length := make([]int64, 0), len(idStrs)
	for i := length - 1; i >= 0; i-- {
		id, _ := strconv.ParseInt(idStrs[i], 10, 64)
		ids = append(ids, id)
	}

	userInfos, uisMap := make([]model.UserInfo, 0), map[int64]model.UserInfo{}
	err = l.svcCtx.UserInfo.In("`id`", ids).Cols("`id`, `name`, `avatar`").Find(&userInfos)
	if err != nil {
		logx.Errorf("[DB ERROR] ListFriendsByUserId 批量查询userInfo失败 %v\n", err)
		return nil, errors.New("出错啦")
	}

	for _, ui := range userInfos {
		uisMap[ui.Id] = ui
	}
	for i, id := range ids {
		userInfos[i] = uisMap[id]
	}
	resp["user_list"] = userInfos
	return resp, nil
}
