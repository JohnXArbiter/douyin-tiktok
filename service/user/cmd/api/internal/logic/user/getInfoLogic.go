package user

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"douyin-tiktok/service/user/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
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

func (l *GetInfoLogic) GetInfo(req *types.UserIdReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		userId       = loggedUser.Id
		targetUserId = req.UserId
	)

	userInfo := &model.UserInfo{Id: targetUserId}
	has, err := l.svcCtx.UserInfo.Omit("`username`", "`password`").Get(userInfo)
	if err != nil || !has {
		return nil, errors.New("找不到该用户")
	}

	if userId != targetUserId {
		userInfo.IsFollow = l.isFollowed(userId, targetUserId)
	}

	resp := utils.GenOkResp()
	resp["user"] = userInfo
	return resp, nil
}

func (l *GetInfoLogic) isFollowed(userId, targetUserId int64) bool {
	var (
		userIdStr       = strconv.FormatInt(userId, 10)
		targetUserIdStr = strconv.FormatInt(targetUserId, 10)
		key             = utils.UserFollow + userIdStr
		userRelation    model.UserRelation
	)

	if exist := l.svcCtx.Redis.Exists(l.ctx, key).Val(); exist == 1 {
		score := l.svcCtx.Redis.ZScore(l.ctx, key, targetUserIdStr).Val()
		if score != 0 {
			return true
		}
	}
	filter := bson.M{"_id": userId}
	projection := bson.D{{"follows", 1}}
	err := l.svcCtx.UserRelation.FindOne(l.ctx, filter, options.FindOne().SetProjection(projection)).Decode(&userRelation)
	if err != nil {
		logx.Errorf("[MONGO ERROR] GetInfo->isFollowed 查询关注文档失败 %v\n", err)
		return false
	}

	flag, zs := false, make([]redis.Z, 0)
	for _, follow := range userRelation.Follows {
		z := redis.Z{Score: float64(follow.Time), Member: follow.UserId}
		zs = append(zs, z)
		if follow.UserId == targetUserId {
			flag = true
		}
	}

	pipeline := l.svcCtx.Redis.Pipeline()
	pipeline.ZAdd(l.ctx, key, zs...)
	pipeline.Expire(l.ctx, key, 5*time.Minute)
	go pipeline.Exec(l.ctx)
	return flag
}
