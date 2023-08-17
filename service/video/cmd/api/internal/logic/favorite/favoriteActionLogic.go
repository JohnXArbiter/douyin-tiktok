package favorite

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/api/internal/logic/mq"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionReq, loggedUser *utils.JwtUser) error {
	var (
		userId        = loggedUser.Id
		videoId       = req.VideoId
		userIdStr     = strconv.FormatInt(userId, 10)
		videoIdStr    = strconv.FormatInt(videoId, 10)
		actionType    = req.ActionType
		key           = utils.VideoFavorite + userIdStr
		rabbitMQLogic = mq.NewRabbitMQLogic(l.ctx, l.svcCtx)
	)

	if actionType == 1 {
		score, err := l.svcCtx.Redis.ZScore(l.ctx, key, videoIdStr).Result()
		if err != nil && err != redis.Nil {
			logx.Errorf("[REDIS ERROR] FavoriteAction %v\n", err)
			return errors.New("出错啦")
		} else if score != 0 {
			return errors.New("你已经点过赞了哦😊")
		}

		z := redis.Z{Score: float64(time.Now().Unix()), Member: videoId}
		if err = l.svcCtx.Redis.ZAdd(l.ctx, key, z).Err(); err != nil {
			logx.Errorf("[REDIS ERROR] FavoriteAction 添加 zset member 失败 %v\n", err)
			return errors.New("出错啦")
		}
		rabbitMQLogic.FavoriteUpdatePublisher(videoId, userId, 0)
	} else if actionType == 2 {
		if exists := l.svcCtx.Redis.Exists(l.ctx, key).Val(); exists == 1 {
			if res, err := l.svcCtx.Redis.ZRem(l.ctx, key, videoId).Result(); err != nil {
				return errors.New("出错啦")
			} else if res == 0 {
				return errors.New("你本来就没有赞人家嘛😥")
			}
			rabbitMQLogic.FavoriteUpdatePublisher(videoId, userId, 1)
		}
		if err := l.FavoriteCancelStrategy2(userId, videoId, userIdStr, key); err != nil {
			return err
		}
	}
	return nil
}

func (l *FavoriteActionLogic) FavoriteCancelStrategy2(userId, videoId int64, userIdStr, key string) error {
	lockKey := utils.VideoFavoriteLock + userIdStr
	lock := utils.NewDistributedLock(l.ctx, l.svcCtx.Redis, lockKey)
	isLocked, err := lock.AcquireLock(time.Second * 5)
	if err != nil {
		return errors.New("出错啦")
	} else if !isLocked {
		return errors.New("你操作地太快啦，请稍后🥵")
	}
	defer lock.ReleaseLock()
	filter := bson.M{"_id": userId}
	targetVideoId := bson.M{"$pull": bson.M{"favorite_videos": bson.M{"video_id": videoId}}}
	_, err = l.svcCtx.VideoFavorite.UpdateOne(l.ctx, filter, targetVideoId)
	if err != nil {
		logx.Errorf("[MONGO ERROR] FavoriteAction 删除文章收藏记录失败 %v\n", err)
	}
	favoriteCommonLogic := NewFavoriteCommonLogic(l.ctx, l.svcCtx)
	_, err = favoriteCommonLogic.LoadIdsAndStore(key, userId)
	return nil
}
