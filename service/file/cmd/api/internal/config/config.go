package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

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
