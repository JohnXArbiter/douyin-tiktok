package mq

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"github.com/streadway/amqp"
)

func InitRabbitMQ(svcCtx *svc.ServiceContext) {
	RabbitMQ = NewRabbitMQLogic(context.Background(), svcCtx)

	go InitVFPublisher(svcCtx.RmqCore)

	go StartVFConsumer(svcCtx.RmqCore.Conn)
	go StartVLConsumer(svcCtx.RmqCore.Conn)
}

func InitVFPublisher(core *utils.RabbitmqCore) {
	// 获取connection
	r := utils.NewRabbitMQ(utils.VideoFavoriteQueue, utils.VideoFavoriteExchange, "cc", core.Conn, core.Channel)
	if r == nil {
		panic("[RABBITMQ ERROR] InitVFPublisher 初始化 cmdty collect publisher 错误！")
	}
	// 延迟队列配置
	delaySeconds := 1000
	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明ttl队列的交换机
	err := r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic("[RABBITMQ ERROR] InitVFPublisher ExchangeDeclare 错误 : " + err.Error())
		return
	}
	args := amqp.Table{
		"x-dead-letter-exchange":    utils.VideoFavoriteDeadExchange,
		"x-dead-letter-routing-key": "cc",
		"x-message-ttl":             int32(delaySeconds * 30),
	}
	// 声明带有ttl的队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		panic("[RABBITMQ ERROR] QueueDeclare error : " + err.Error())
		return
	}
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		panic("[RABBITMQ ERROR] QueueBinding error : " + err.Error())
		return
	}
}
