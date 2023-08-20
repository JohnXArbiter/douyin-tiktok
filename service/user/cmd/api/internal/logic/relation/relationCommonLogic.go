package relation

import (
	"context"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/model"
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

func (l *RelationCommonLogic) ListFollowedUsersOrFans(userId, isFollow int64, key string) []model.UserInfo {
	var (
		ids       []int64
		userInfos []model.UserInfo
	)

	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] ListFollowedUsersOrFans sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 { //
		var userRelation, err = l.LoadIdsFromMongo(userId, isFollow)
		if (userRelation == nil && err == nil) || (len(userRelation.Fans) == 0 && len(userRelation.Followers) == 0) {
			return make([]model.UserInfo, 0)
		} else if err != nil {
			return nil
		}

		if isFollow == 1 {
			zs, _ = l.StoreRelatedUsers2Redis(userRelation.Followers, key)
		} else {
			zs, _ = l.StoreRelatedUsers2Redis(userRelation.Fans, key)
		}
	}

	for _, z := range zs {
		id, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		ids = append(ids, id)
	}
	if err = l.svcCtx.UserInfo.In("`id`", ids).Cols("`id`, `name`, `avatar`").Find(&userInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFollowedUserByUserId 批量查询userInfo失败 %v\n", err)
		return nil
	}

	uiMap := make(map[int64]model.UserInfo)
	for _, info := range userInfos {
		uiMap[info.Id] = info
	}
	for i, id := range ids {
		userInfos[i] = uiMap[id]
	}

	return userInfos
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
