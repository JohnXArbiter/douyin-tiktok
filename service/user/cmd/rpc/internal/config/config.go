package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql utils.Mysql
	Redis utils.Redis
	Mongo utils.Mongo
}
