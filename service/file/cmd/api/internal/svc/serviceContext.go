package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/file/cmd/api/internal/config"
	"douyin-tiktok/service/user/cmd/rpc/userservice"
	"douyin-tiktok/service/video/cmd/rpc/videoservice"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	UserRpc  userservice.UserService
	VideoRpc videoservice.VideoService

	Xorm      *xorm.Engine
	FileVideo *xorm.Session
	FileUser  *xorm.Session
	FileCover *xorm.Session

	Oss *Oss
}

type Oss struct {
	Client     *oss.Client
	BaseUrl    string
	BucketName string
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	endPoint := c.Oss.EndPoint
	accessKeyId := c.Oss.AccessKeyId
	accessKeySecret := c.Oss.AccessKeySecret
	bucketName := c.Oss.BucketName
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic("[OSS ERROR] NewServiceContext 获取OSS连接错误" + err.Error())
	}

	return &ServiceContext{
		Config:    c,
		Xorm:      engine,
		FileVideo: engine.Table("file_video"),
		FileUser:  engine.Table("file_user"),
		FileCover: engine.Table("file_cover"),
		UserRpc:   userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		VideoRpc:  videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
		Oss: &Oss{
			Client:     client,
			BaseUrl:    "https://" + bucketName + "." + endPoint + "/",
			BucketName: bucketName,
		},
	}
}
