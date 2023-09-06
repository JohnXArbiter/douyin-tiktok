package user

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/logic/relation"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"douyin-tiktok/service/user/model"
	__video "douyin-tiktok/service/video/cmd/rpc/types"
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
		rpcChan      = make(chan *__video.GetFavoriteAndFavoritedCntResp)
	)

	go func() {
		var res *__video.GetFavoriteAndFavoritedCntResp
		defer func() {
			rpcChan <- res
		}()

		res, err := l.svcCtx.VideoRpc.GetFavoriteAndFavoritedCnt(l.ctx, &__video.GetFavoriteAndFavoritedCntReq{UserId: userId})
		if err != nil || res.Code != 0 {
			logx.Errorf("[RPC ERROR] GetInfo 调用rpc获取获赞数和点赞数失败 %v\n", err)
		}
	}()

	userInfo := &model.UserInfo{Id: targetUserId}
	has, err := l.svcCtx.UserInfo.Omit("`username`", "`password`").Get(userInfo)
	if !has {
		if err != nil {
			logx.Errorf("[DB ERROR] GetInfo 查询用户信息失败 %v\n", err)
		}
		return nil, errors.New("找不到该用户")
	}

	if userId != targetUserId {
		userInfo.IsFollow = l.isFollowed(userId, targetUserId)
	}

	userInfo.FollowCount, userInfo.FollowerCount = l.getRelationCnt(userId)

	rpcResp := <-rpcChan
	if rpcResp != nil {
		userInfo.TotalFavorited = rpcResp.TotalFavorited
		userInfo.FavoriteCount = rpcResp.FavoriteCount
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
	for _, follow := range userRelation.Followers {
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

func (l *GetInfoLogic) getRelationCnt(userId int64) (int64, int64) {
	userIdStr := strconv.FormatInt(userId, 10)
	relationCommonLogic := relation.NewRelationCommonLogic(l.ctx, l.svcCtx)
	followCnt := relationCommonLogic.LoadIdsAndStore(userId, 1, utils.UserFollow+userIdStr)
	fansCnt := relationCommonLogic.LoadIdsAndStore(userId, 0, utils.UserFan+userIdStr)
	return int64(len(followCnt)), int64(len(fansCnt))
}
