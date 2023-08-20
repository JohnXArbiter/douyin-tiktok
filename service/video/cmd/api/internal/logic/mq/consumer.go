package mq

import (
	"douyin-tiktok/common/utils"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

func StartVFConsumer(conn *amqp.Connection) {
	channel, err := conn.Channel()
	if err != nil {
		panic("[RABBITMQ ERROR] StartVFConsumer 初始化失败 " + err.Error())
	}
	r := utils.NewRabbitMQ(utils.VideoFavoriteDeadQueue, utils.VideoFavoriteDeadExchange, utils.VideoFavoriteDeadRoutingKey, conn, channel)

	exchangeName := r.Exchange
	queueName := r.QueueName
	key := r.Key
	// 声明死信交换机
	err = r.Channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 声明有死信队列
	_, err = r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 将死信交换机和死信队列绑定
	err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
	if err != nil {
		panic(err)
	}
	// 开始监听
	msgs, err := r.Channel.Consume(utils.VideoFavoriteDeadQueue, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	forever := make(chan int, 0)
	for msg := range msgs {
		var vfMsg = new(utils.VFMessage)
		if err = Json.Unmarshal(msg.Body, vfMsg); err != nil {
			logx.Infof("[RABBITMQ ERROR] StartVFConsumer 反序列化消息失败 %v\n", err)
			msg.Nack(false, false)
			continue
		}
		fmt.Printf("%+v \n dhsdfdjash", vfMsg)
		if vfMsg.IsCancel == 0 {
			RabbitMQ.FavoriteCheck(vfMsg)
		} else {
			RabbitMQ.FavoriteCancelCheck(vfMsg)
		}
		msg.Ack(false)
	}
	<-forever
}
