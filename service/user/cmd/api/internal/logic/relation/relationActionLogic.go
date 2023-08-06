package relation

import (
	"context"
	"douyin-tiktok/common/utils"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

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
	var filter = bson.M{"_id": userId}
	followedUser := bson.M{"$addToSet": bson.M{
		"follow": bson.M{
			"user_id": toUserId,
			"time":    time.Now().Unix(),
		}},
	}

	_, err := l.svcCtx.UserRelation.UpdateOne(l.ctx, filter, followedUser)
	return err
}

func (l *RelationActionLogic) unFollow(userId, toUserId int64) error {
	var filter = bson.M{"_id": userId}
	targetUser := bson.M{"$addToSet": bson.M{"follow": bson.M{"user_id": toUserId}}}

	// 执行更新操作
	_, err := l.svcCtx.UserRelation.UpdateOne(context.Background(), filter, targetUser)
	return err
}
