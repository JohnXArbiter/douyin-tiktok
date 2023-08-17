package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/file/cmd/rpc/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

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
		Oss: &Oss{
			Client:     client,
			BaseUrl:    "https://" + bucketName + "." + endPoint + "/",
			BucketName: bucketName,
		},
	}
}
