package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Dsn string
	}

	Oss struct {
		AccessKeyId     string
		AccessKeySecret string
		EndPoint        string
		BucketName      string
	}
}
