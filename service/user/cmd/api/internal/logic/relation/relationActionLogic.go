package relation

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"douyin-tiktok/service/user/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionReq, loggedUser *utils.JwtUser) error {
	var (
		userId   = loggedUser.Id
		toUserId = req.ToUserId
	)

	if req.ActionType == 1 {
		if err := l.follow(userId, toUserId); err != nil {
			logx.Errorf("[DB ERROR] RelationAction 关注失败 %v\n", err)
			return errors.New("关注失败")
		}
	} else {
		if err := l.unFollow(userId, toUserId); err != nil {
			logx.Errorf("[DB ERROR] RelationAction 取消关注失败 %v\n", err)
			return errors.New("取消关注失败")
		}
	}
	return nil
}

func (l *RelationActionLogic) follow(userId, toUserId int64) error {
	var (
		userIdStr = strconv.FormatInt(userId, 10)
		key       = utils.UserFollow + userIdStr
		now       = time.Now().Unix()
	)

	member := redis.Z{Score: float64(now), Member: toUserId}
	if res, err := l.svcCtx.Redis.ZAdd(l.ctx, key, member).Result(); err != nil {
		logx.Errorf("[REDIS ERROR] RelationAction->follow sth wrong with redis %v\n", err)
	} else if res == 0 {
		return nil // zset 中已经有了
	}

	relatedUser := model.RelatedUsers{
		UserId: toUserId,
		Time:   now,
	}
	updateOpt := options.Update().SetUpsert(true)
	filter := bson.M{"_id": userId}
	followedUser := bson.M{"$addToSet": bson.M{
		"follows": relatedUser,
	}}
	_, err := l.svcCtx.UserRelation.UpdateOne(l.ctx, filter, followedUser, updateOpt)
	if err != nil {
		return err
	}

	filter = bson.M{"_id": toUserId}
	relatedUser.UserId = userId
	fan := bson.M{"$addToSet": bson.M{
		"fans": relatedUser,
	}}
	_, err = l.svcCtx.UserRelation.UpdateOne(l.ctx, filter, fan, updateOpt)
	return err
}

func (l *RelationActionLogic) unFollow(userId, toUserId int64) error {
	var (
		userIdStr = strconv.FormatInt(userId, 10)
		key       = utils.UserFollow + userIdStr
	)

	if err := l.svcCtx.Redis.ZRem(l.ctx, key, toUserId).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] RelationAction->unfollow zrem 失败 %v\n", err)
	}

	filter := bson.M{"_id": userId}
	targetUser := bson.M{"$pull": bson.M{"follows": bson.M{"user_id": toUserId}}}
	_, err := l.svcCtx.UserRelation.UpdateOne(context.Background(), filter, targetUser)
	if err != nil {
		return err
	}

	filter = bson.M{"_id": toUserId}
	targetUser = bson.M{"$pull": bson.M{"fans": bson.M{"user_id": userId}}}
	_, err = l.svcCtx.UserRelation.UpdateOne(context.Background(), filter, targetUser)
	return err
}
