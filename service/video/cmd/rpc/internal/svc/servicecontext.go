package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/rpc/internal/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm         *xorm.Engine
	VideoInfo    *xorm.Session
	VideoComment *xorm.Session

	Mongo         *mongo.Database
	VideoFavorite *mongo.Collection

	Redis *redis.Client

	Json jsoniter.API
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	options := idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)

	mdb := utils.InitMongo(c.Mongo).Database("douyin_video")

	return &ServiceContext{
		Config:        c,
		Xorm:          engine,
		VideoInfo:     engine.Table("video_info"),
		VideoComment:  engine.Table("video_comment"),
		Mongo:         mdb,
		VideoFavorite: mdb.Collection("video_favorite"),
		Redis:         utils.InitRedis(c.Redis2),
		Json:          jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}
