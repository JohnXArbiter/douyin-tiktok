package relation

import (
	"context"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

type RelationCommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationCommonLogic {
	return &RelationCommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationCommonLogic) ListFollowedUsersOrFans(userId, isFollow int64, key string) []model.UserInfo {
	var (
		ids       []int64
		userInfos []model.UserInfo
	)

	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] ListFollowedUsersOrFans sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 { //
		var users, emptyFlag = l.loadIdsFromMongo(userId, isFollow)
		if err != nil {
			return nil
		} else if emptyFlag {
			return make([]model.UserInfo, 0)
		}

		for _, user := range users {
			z := redis.Z{
				Score:  float64(user.Time),
				Member: strconv.FormatInt(user.UserId, 10),
			}
			zs = append(zs, z)
		}
		fmt.Println(zs)
		if err = l.svcCtx.Redis.ZAdd(l.ctx, key, zs...).Err(); err != nil {
			logx.Errorf("[REDIS ERROR] ListFollowedUsersOrFans 关注列表存储redis失败 %v\n", err)
		}
	}

	for _, z := range zs {
		id, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		ids = append(ids, id)
	}
	if err = l.svcCtx.UserInfo.In("`id`", ids).Find(&userInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFollowedUserByUserId 批量查询userInfo失败 %v\n", err)
		return nil
	}
	fmt.Println("asdasopdjasod", err, userInfos)

	uiMap := make(map[int64]model.UserInfo)
	for _, info := range userInfos {
		uiMap[info.Id] = info
	}
	for i, id := range ids {
		userInfos[i] = uiMap[id]
	}

	return userInfos
}

func (l *RelationCommonLogic) loadIdsFromMongo(id, isFollow int64) ([]model.RelatedUsers, bool) {
	var (
		userRelation model.UserRelation
		projection   = bson.D{{"followers", 0}}
	)

	filter := bson.M{"_id": id}
	if isFollow == 1 {
		projection = bson.D{{"fans", 0}}
	}

	err := l.svcCtx.UserRelation.FindOne(l.ctx, filter, options.FindOne().SetProjection(projection)).Decode(&userRelation)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, true
		}
		logx.Errorf("[MONGO ERROR] ListFollowedUsersOrFans->loadIdsFromMongo 查询关注文档失败 %v\n", err)
		return nil, false
	}

	res, emptyFlag := l.reverse(userRelation, isFollow)
	return res, emptyFlag
}

func (l *RelationCommonLogic) reverse(userRelation model.UserRelation, isFollow int64) ([]model.RelatedUsers, bool) {
	var users, res []model.RelatedUsers

	if isFollow == 1 {
		users = userRelation.Follows
	} else {
		users = userRelation.Fans
	}

	length := len(users) - 1
	if length <= 0 {
		return nil, true
	}
	for i := length; i >= 0; i-- {
		res = append(res, users[i])
	}
	return res, false
}
