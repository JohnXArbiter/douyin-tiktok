package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

var UserServiceRedis *redis.Client // 供每个服务使用的

type Redis struct {
	Addr     string
	Password string
	Db       int
}

const (
	UserLogged = "user:logged:"

	UserFollow = "user:follow:"
)

func UserServiceInit(ctx context.Context, client *redis.Client) {
	UserServiceRedis = client
	err := client.Ping(ctx).Err()
	if err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
}

func InitRedis(rc Redis) *redis.Client {
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
