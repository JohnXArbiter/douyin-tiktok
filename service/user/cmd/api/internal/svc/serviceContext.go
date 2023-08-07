package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/config"
	"github.com/yitter/idgenerator-go/idgen"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm         *xorm.Engine
	UserInfo     *xorm.Session
	UserRelation *xorm.Session
	UserMessage  *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	options := idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)

	engine := utils.InitXorm("mysql", c.Mysql)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		UserInfo:     engine.Table("user_info"),
		UserRelation: engine.Table("user_relation"),
		UserMessage:  engine.Table("user_message"),
	}
}
