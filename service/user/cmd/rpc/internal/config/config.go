package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	Mysql  utils.MysqlConf
	Redis1 utils.RedisConf
	Mongo  utils.MongoConf
}
