package logic

import (
	"context"
	"douyin-tiktok/service/video/cmd/rpc/internal/svc"
	"douyin-tiktok/service/video/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
	"strconv"
	"time"
)

type FavoriteCommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteCommonLogic {
	return &FavoriteCommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteCommonLogic) LoadIdsAndStore(key string, userId int64) ([]int64, error) {
	var ids []int64

	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] ListFavoriteVideos sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 {
		var videoFavorite, err = l.LoadIdsFromMongo(userId)
		if (videoFavorite == nil && err == nil) || (videoFavorite != nil && len(videoFavorite.FavoriteVideos) == 0) {
			return ids, nil
		} else if err != nil {
			return nil, errors.New("出错啦")
		}
		return l.StoreFavoriteVideos2Redis(videoFavorite.FavoriteVideos, key)
	}
	for _, z := range zs {
		videoId, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		ids = append(ids, videoId)
	}
	return ids, nil
}

// LoadIdsFromMongo 从 mongo 中取 favorite_videos 字段
func (l *FavoriteCommonLogic) LoadIdsFromMongo(id int64) (*model.VideoFavorite, error) {
	var videoFavorite model.VideoFavorite

	filter := bson.M{"_id": id}
	err := l.svcCtx.VideoFavorite.FindOne(l.ctx, filter).Decode(&videoFavorite)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		logx.Errorf("[MONGO ERROR] LoadIdsFromMongo 查询关注文档失败 %v\n", err)
		return nil, err
	}
	return &videoFavorite, nil
}

func (l *FavoriteCommonLogic) StoreFavoriteVideos2Redis(videos model.FavoriteVideos, key string) ([]int64, error) {
	var zs, ids = make([]redis.Z, 0), make([]int64, 0)
	sort.Sort(videos)
	for _, video := range videos {
		z := redis.Z{
			Score:  float64(video.Time),
			Member: strconv.FormatInt(video.VideoId, 10),
		}
		zs = append(zs, z)
		ids = append(ids, video.VideoId)
	}
	if err := l.svcCtx.Redis.ZAdd(l.ctx, key, zs...).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] StoreFavoriteVideos2Redis 点赞列表存储redis失败 %v\n", err)
		return ids, err
	}
	l.svcCtx.Redis.Expire(l.ctx, key, time.Minute*10)
	return ids, nil
}
