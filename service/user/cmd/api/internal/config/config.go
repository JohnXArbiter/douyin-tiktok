package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mongo utils.Mongo
	Mysql utils.Mysql
}
