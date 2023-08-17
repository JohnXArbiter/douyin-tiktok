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
	Oss    Oss
}

type Oss struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}
