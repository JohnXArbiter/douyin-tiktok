package relation

import (
	"context"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
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

func (l *RelationCommonLogic) ListFollowedUsersOrFans(userId, isFollow int64, key string) []*model.UserInfo {
	var (
		ids       []int64
		userInfos []*model.UserInfo
		zs        []redis.Z
	)

	if isFollow == 1 {
		zs = l.LoadIdsAndStore(userId, 1, key)
	} else {
		zs = l.LoadIdsAndStore(userId, 0, key)
	}

	if zs == nil || len(zs) == 0 {
		return make([]*model.UserInfo, 0)
	}

	for _, z := range zs {
		id, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		ids = append(ids, id)
	}
	if err := l.svcCtx.UserInfo.Cols("`id`, `name`, `avatar`").
		In("`id`", ids).Find(&userInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFollowedUserByUserId 批量查询userInfo失败 %v\n", err)
		return nil
	}

	if len(userInfos) < 1 {
		return make([]*model.UserInfo, 0)
	}

	uiMap := make(map[int64]*model.UserInfo)
	for i := 0; i < len(userInfos); i++ {
		uiMap[userInfos[i].Id] = userInfos[i]
	}
	fmt.Println(len(ids))
	fmt.Println(ids)
	fmt.Println(uiMap)
	for i := 0; i < len(ids); i++ {
		userInfos[i] = uiMap[ids[i]]
	}

	return userInfos
}

func (l *RelationCommonLogic) LoadIdsAndStore(userId, isFollow int64, key string) []redis.Z {
	var zs, err = l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] LoadIdsAndStore sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 { //
		var userRelation, err = l.LoadIdsFromMongo(userId, isFollow)
		if (userRelation == nil && err == nil) || err != nil {
			return make([]redis.Z, 0)
		}

		if isFollow == 1 {
			zs, _ = l.StoreRelatedUsers2Redis(userRelation.Followers, key)
		} else {
			zs, _ = l.StoreRelatedUsers2Redis(userRelation.Fans, key)
		}
	}
	return zs
}

// LoadIdsFromMongo 从 mongo 中取 follows 或 fans 字段
func (l *RelationCommonLogic) LoadIdsFromMongo(id, isFollow int64) (*model.UserRelation, error) {
	var userRelation model.UserRelation

	filter := bson.M{"_id": id}
	projection := bson.D{{"followers", 0}}
	if isFollow == 1 {
		projection = bson.D{{"fans", 0}}
	}

	err := l.svcCtx.UserRelation.FindOne(l.ctx, filter, options.FindOne().SetProjection(projection)).Decode(&userRelation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		logx.Errorf("[MONGO ERROR] LoadIdsFromMongo 查询关注文档失败 %v\n", err)
		return nil, err
	}
	return &userRelation, nil
}

func (l *RelationCommonLogic) StoreRelatedUsers2Redis(relatedUsers model.RelatedUsers, key string) ([]redis.Z, error) {
	var zs []redis.Z
	for _, user := range relatedUsers {
		z := redis.Z{
			Score:  float64(user.Time),
			Member: strconv.FormatInt(user.UserId, 10),
		}
		zs = append(zs, z)
	}
	if err := l.svcCtx.Redis.ZAdd(l.ctx, key, zs...).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] StoreRelatedUsers2Redis 关注列表存储redis失败 %v\n", err)
		return zs, err
	}
	return zs, nil
}
