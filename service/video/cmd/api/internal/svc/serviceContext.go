package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/file/cmd/rpc/fileservice"
	"douyin-tiktok/service/user/cmd/rpc/userservice"
	"douyin-tiktok/service/video/cmd/api/internal/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/zrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userservice.UserService
	FileRpc fileservice.FileService

	Xorm         *xorm.Engine
	VideoInfo    *xorm.Session
	VideoComment *xorm.Session

	Mongo         *mongo.Client
	VideoFavorite *mongo.Collection

	Redis *redis.Client

	RmqCore *utils.RabbitmqCore

	Json jsoniter.API
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	options := idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)

	mc := utils.InitMongo(c.Mongo)
	rc, channel := utils.InitRabbitMQ(c.RabbitMQ)

	return &ServiceContext{
		Config:        c,
		Xorm:          engine,
		VideoInfo:     engine.Table("video_info"),
		VideoComment:  engine.Table("video_comment"),
		Mongo:         mc,
		VideoFavorite: mc.Database("douyin_user").Collection("user_relation"),
		UserRpc:       userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		FileRpc:       fileservice.NewFileService(zrpc.MustNewClient(c.FileRpc)),
		Redis:         utils.InitRedis(c.Redis),
		Json:          jsoniter.ConfigCompatibleWithStandardLibrary,
		RmqCore: &utils.RabbitmqCore{
			Conn:    rc,
			Channel: channel,
		},
	}
}
