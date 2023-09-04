package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/config"
	"douyin-tiktok/service/video/cmd/rpc/videoservice"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/zrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc videoservice.VideoService

	Xorm        *xorm.Engine
	UserInfo    *xorm.Session
	UserMessage *xorm.Session

	Mongo        *mongo.Client
	UserRelation *mongo.Collection

	Redis *redis.Client

	BgUrl string
}

func NewServiceContext(c config.Config) *ServiceContext {
	options := idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)

	engine := utils.InitXorm("mysql", c.Mysql)

	mc := utils.InitMongo(c.Mongo)
	fmt.Printf("%+v\n", c.VideoRpc)

	return &ServiceContext{
		Config:       c,
		VideoRpc:     videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
		Xorm:         engine,
		UserInfo:     engine.Table("user_info"),
		UserMessage:  engine.Table("user_message"),
		Mongo:        mc,
		UserRelation: mc.Database("douyin_user").Collection("user_relation"),
		Redis:        utils.InitRedis(c.Redis),
		BgUrl:        c.Bg.Url,
	}
}
