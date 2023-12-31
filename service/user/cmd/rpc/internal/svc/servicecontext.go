package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/rpc/internal/config"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm        *xorm.Engine
	UserInfo    *xorm.Session
	UserMessage *xorm.Session

	Mongo        *mongo.Client
	UserRelation *mongo.Collection

	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	mc := utils.InitMongo(c.Mongo)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		UserInfo:     engine.Table("user_info"),
		UserMessage:  engine.Table("user_message"),
		Mongo:        mc,
		UserRelation: mc.Database("douyin_user").Collection("user_relation"),
		Redis:        utils.InitRedis(c.Redis1),
	}
}
