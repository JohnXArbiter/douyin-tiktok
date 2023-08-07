package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	VideoRpc zrpc.RpcClientConf

	Postgresql struct {
		Dsn string
	}

	Oss struct {
		AccessKeyId     string
		AccessKeySecret string
		EndPoint        string
		BucketName      string
	}
}