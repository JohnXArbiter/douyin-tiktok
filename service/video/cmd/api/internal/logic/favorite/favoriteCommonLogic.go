package favorite

import (
	"context"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
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

func (l *FavoriteCommonLogic) LoadIdsAndStore(key string, userId int64) ([]redis.Z, error) {
	var zs, err = l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] ListFavoriteVideos sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 {
		var videoFavorite, err = l.LoadIdsFromMongo(userId)
		if (videoFavorite == nil && err == nil) || len(videoFavorite.FavoriteVideos) == 0 {
			return make([]redis.Z, 0), nil
		} else if err != nil {
			return nil, errors.New("出错啦")
		}

		favoriteVideos := l.reverse(videoFavorite.FavoriteVideos)
		zs, _ = l.StoreFavoriteVideos2Redis(favoriteVideos, key)
	}
	return zs, nil
}

// LoadIdsFromMongo 从 mongo 中取 favorite_videos 字段
func (l *FavoriteCommonLogic) LoadIdsFromMongo(id int64) (*model.VideoFavorite, error) {
	var videoFavorite model.VideoFavorite

	filter := bson.M{"_id": id}
	err := l.svcCtx.VideoFavorite.FindOne(l.ctx, filter).Decode(&videoFavorite)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, nil
		}
		logx.Errorf("[MONGO ERROR] LoadIdsFromMongo 查询关注文档失败 %v\n", err)
		return nil, err
	}
	return &videoFavorite, nil
}

// 倒序排
func (l *FavoriteCommonLogic) reverse(videos []model.FavoriteVideos) []model.FavoriteVideos {
	var res []model.FavoriteVideos

	length := len(videos) - 1
	for i := length; i >= 0; i-- {
		res = append(res, videos[i])
	}
	return res
}

func (l *FavoriteCommonLogic) StoreFavoriteVideos2Redis(videos []model.FavoriteVideos, key string) ([]redis.Z, error) {
	var zs []redis.Z
	for _, video := range videos {
		z := redis.Z{
			Score:  float64(video.Time),
			Member: strconv.FormatInt(video.VideoId, 10),
		}
		zs = append(zs, z)
	}
	if err := l.svcCtx.Redis.ZAdd(l.ctx, key, zs...).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] StoreFavoriteVideos2Redis 点赞列表存储redis失败 %v\n", err)
		return zs, err
	}
	l.svcCtx.Redis.Expire(l.ctx, key, time.Minute*10)
	return zs, nil
}
