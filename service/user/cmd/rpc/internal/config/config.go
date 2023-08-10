package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql  utils.MysqlConf
	Redis1 utils.RedisConf
	Mongo  utils.MongoConf
}
