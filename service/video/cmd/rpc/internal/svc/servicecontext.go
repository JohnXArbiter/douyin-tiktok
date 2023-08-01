package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/rpc/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm          *xorm.Engine
	VideoInfo     *xorm.Session
	VideoFavorite *xorm.Session
	VideoComment  *xorm.Session

	//Redis *redis.Client

}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	return &ServiceContext{
		Config:        c,
		Xorm:          engine,
		VideoInfo:     engine.Table("video_info"),
		VideoFavorite: engine.Table("video_favorite"),
		VideoComment:  engine.Table("video_comment"),
		//Redis:        utils.InitRedis(c.Redis),
	}
}
