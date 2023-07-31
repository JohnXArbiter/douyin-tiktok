package utils

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Url string
}

func InitMongo(mc Mongo) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mc.Url) // 设置客户端连接配置
	client, err := mongo.NewClient(clientOptions)      // 创建客户端
	logx.Infof("[MONGO CONNECTING] Init Mongo URL: %v", mc.Url)
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败 " + err.Error())
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败 " + err.Error())
	}
	return client
}
