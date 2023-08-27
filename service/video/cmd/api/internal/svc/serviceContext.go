package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/rpc/userservice"
	"douyin-tiktok/service/video/cmd/api/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

	Xorm         *xorm.Engine
	VideoInfo    *xorm.Session
	VideoComment *xorm.Session

	Mongo         *mongo.Database
	VideoFavorite *mongo.Collection

	Redis *redis.Client

	RmqCore *utils.RabbitmqCore

	Json jsoniter.API
	Oss  *Oss
}

type Oss struct {
	Client     *oss.Client
	BaseUrl    string
	BucketName string
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	options := idgen.NewIdGeneratorOptions(20)
	idgen.SetIdGenerator(options)

	mdb := utils.InitMongo(c.Mongo).Database("douyin_video")
	rc, channel := utils.InitRabbitMQ(c.RabbitMQ)
	endPoint := c.Oss.EndPoint
	accessKeyId := c.Oss.AccessKeyId
	accessKeySecret := c.Oss.AccessKeySecret
	bucketName := c.Oss.BucketName

	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic("[OSS ERROR] NewServiceContext 获取OSS连接错误" + err.Error())
	}

	return &ServiceContext{
		Config:        c,
		Xorm:          engine,
		VideoInfo:     engine.Table("video_info"),
		VideoComment:  engine.Table("video_comment"),
		Mongo:         mdb,
		VideoFavorite: mdb.Collection("video_favorite"),
		UserRpc:       userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		Redis:         utils.InitRedis(c.Redis),
		Json:          jsoniter.ConfigCompatibleWithStandardLibrary,
		RmqCore: &utils.RabbitmqCore{
			Conn:    rc,
			Channel: channel,
		},
		Oss: &Oss{
			Client:     client,
			BaseUrl:    "https://" + bucketName + "." + endPoint + "/",
			BucketName: bucketName,
		},
	}
}
