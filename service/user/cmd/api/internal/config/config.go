package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	VideoRpc zrpc.RpcClientConf
	Mongo    utils.MongoConf
	Mysql    utils.MysqlConf
	Redis    utils.RedisConf
	Bg       struct {
		Url string `json:"Url"`
	} `json:"Bg"`
}
