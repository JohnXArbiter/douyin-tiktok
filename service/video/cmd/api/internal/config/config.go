package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	Mysql    utils.MysqlConf
	Redis    utils.RedisConf
	Mongo    utils.MongoConf
	RabbitMQ utils.RabbitMQConf
	Oss      Oss
}

type Oss struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}
