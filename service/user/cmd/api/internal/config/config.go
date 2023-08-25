package config

import (
	"douyin-tiktok/common/utils"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mongo utils.MongoConf
	Mysql utils.MysqlConf
	Redis utils.RedisConf
	Bg    struct {
		Url string `json:"Url"`
	} `json:"Bg"`
}
