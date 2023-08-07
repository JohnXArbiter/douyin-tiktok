package relation

import (
	"context"
	"douyin-tiktok/common/utils"
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

	if exists := l.svcCtx.Redis.Exists(l.ctx, followKey).Val(); exists == 0 {
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 1)
		if userRelation == nil && err == nil {
			resp["user_list"] = make([]struct{}, 0)
			return nil, errors.New("出错啦")
		}

		_, err = relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Follows, followKey)
		if err != nil {
			return nil, errors.New("出错啦")
		}
	}

	if exists := l.svcCtx.Redis.Exists(l.ctx, fanKey).Val(); exists == 0 {
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 1)
		if userRelation == nil && err == nil {
			resp["user_list"] = make([]struct{}, 0)
			return nil, errors.New("出错啦")
		}

		_, err = relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Follows, fanKey)
		if err != nil {
			return nil, errors.New("出错啦")
		}
	}

	store := &redis.ZStore{Keys: []string{followKey, fanKey}, Weights: []float64{1, 1}}
	_, err := l.svcCtx.Redis.ZInter(l.ctx, store).Result()
	if err != nil {
		logx.Errorf("[REDIS ERROR] ListFriendsByUserId 获取用户:%v，关注粉丝交集失败 %v\n", err)
		return nil, errors.New("出错啦")
	}
	// todo
	//for i, s := range result {
	//
	//}

	return nil, nil
}
