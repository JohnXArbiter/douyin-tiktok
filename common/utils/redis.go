package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var UserServiceRedis *redis.Client // 供每个服务使用的

type RedisConf struct {
	Addr     string
	Password string
	Db       int
}

const (
	UserLogged        = "user:logged:"
	UserFollow        = "user:follow:"
	UserFan           = "user:fan:"
	VideoFavorite     = "video:favorite:"
	VideoFavoriteCnt  = "video:favorite:count:"
	VideoFavoriteLock = "video:favorite:lock:"
)

func UserServiceInit(ctx context.Context, client *redis.Client) {
	UserServiceRedis = client
	err := client.Ping(ctx).Err()
	if err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
}

func InitRedis(rc RedisConf) *redis.Client {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Password,
		DB:       rc.Db,
	})
	logx.Infof("[REDIS CONNECTING] InitRedis client: %v\n", client)

	err := client.Ping(ctx).Err()
	if err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
	UserServiceInit(ctx, client)
	return client
}

type DistributedLock struct {
	ctx context.Context
	rc  *redis.Client
	Key string
}

func NewDistributedLock(ctx context.Context, rc *redis.Client, Key string) *DistributedLock {
	return &DistributedLock{
		ctx: ctx,
		rc:  rc,
		Key: Key,
	}
}

func (l *DistributedLock) AcquireLock(ttl time.Duration) (bool, error) {
	success, err := l.rc.SetNX(l.ctx, l.Key, "", ttl).Result()
	if err != nil {
		logx.Errorf("[REDIS ERROR] AcquireLock 设置锁失败 %v\n", err)
	}
	return success, err
}

func (l *DistributedLock) ReleaseLock() error {
	_, err := l.rc.Del(l.ctx, l.Key).Result()
	if err != nil {
		logx.Errorf("[REDIS ERROR] ReleaseLock 释放锁失败 %v\n", err)
	}
	return nil
}
