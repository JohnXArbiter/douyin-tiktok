package mq

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/model"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"strconv"
)

type RabbitMQLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRabbitMQLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitMQLogic {
	return &RabbitMQLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	RabbitMQ *RabbitMQLogic
	Json     = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (l *RabbitMQLogic) FavoriteCheck(vfMsg *utils.VFMessage) {
	var (
		userId  = vfMsg.UserId
		videoId = vfMsg.VideoId
		key     = utils.VideoFavorite + strconv.FormatInt(userId, 10)
	)
	score, err := l.svcCtx.Redis.ZScore(l.ctx, key, strconv.FormatInt(videoId, 10)).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] FavoriteCheck 获取 zset member 失败 %v\n", err)
		return
	} else if score == 0 {
		return
	}

	filter := bson.M{"_id": userId}
	option := options.Update().SetUpsert(true)
	favoriteVideo := bson.M{"$addToSet": bson.M{
		"favorite_videos": model.FavoriteVideo{
			VideoId: videoId,
			Time:    int64(score),
		},
	}}
	_, err = l.svcCtx.VideoFavorite.UpdateOne(l.ctx, filter, favoriteVideo, option)
	if err != nil {
		logx.Errorf("[MONGO ERROR] FavoriteCheck 插入文章收藏记录失败 %v\n", err)
	}
	fmt.Println("记录一下 视频点赞", err)
}

func (l *RabbitMQLogic) FavoriteCancelCheck(vfMsg *utils.VFMessage) {
	var (
		userId  = vfMsg.UserId
		videoId = vfMsg.VideoId
		key     = utils.VideoFavorite + strconv.FormatInt(userId, 10)
	)

	score, err := l.svcCtx.Redis.ZScore(l.ctx, key, strconv.FormatInt(videoId, 10)).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] FavoriteCancelCheck 获取 zset member 失败 %v\n", err)
		return
	} else if score != 0 {
		return
	}

	filter := bson.M{"_id": userId}
	targetVideoId := bson.M{"$pull": bson.M{"favorite_videos": bson.M{"video_id": videoId}}}
	_, err = l.svcCtx.VideoFavorite.UpdateOne(l.ctx, filter, targetVideoId)
	if err != nil {
		logx.Errorf("[MONGO ERROR] FavoriteCancelCheck 删除文章收藏记录失败 %v\n", err)
	}
	fmt.Println("记录一下 视频取消点赞", err)
}
