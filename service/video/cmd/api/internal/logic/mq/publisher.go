package mq

import (
	"douyin-tiktok/common/utils"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func CollectUpdateFavoritePublisher(redisKey string, member int64, isCollect bool, mqLogic *RabbitMQLogic) {
	ticker := time.NewTicker(time.Second * 30)
	msg := &utils.VFMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      time.Now().Local(),
		IsCollect: isCollect,
	}
	body, _ := Json.Marshal(msg)
	publisher := utils.NewRabbitMQ(utils.VideoFavoriteQueue, utils.VideoFavoriteExchange, "cc", mqLogic.svcCtx.RmqCore.Conn, mqLogic.svcCtx.RmqCore.Channel)
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] CollectUpdatePublisher 发送收藏消息失败 %v\n", err)
		select {
		case <-ticker.C:
			if isCollect {
				go mqLogic.FavoriteCheck(msg)
			} else {
				go mqLogic.FavoriteCancelCheck(msg)
			}
		}
	}
}

func LikeCheckUpdate(redisKey string, member int64, mqLogic *RabbitMQLogic) {
	ticker := time.NewTicker(time.Second * 30)
	msg := &utils.VLMessage{
		RedisKey: redisKey,
		Time:     time.Now().Local(),
		UserId:   member,
	}
	body, _ := Json.Marshal(msg)
	publisher := utils.NewRabbitMQ(utils.VideoLikeQueue, utils.VideoLikeExchange, "cl", mqLogic.svcCtx.RmqCore.Conn, mqLogic.svcCtx.RmqCore.Channel)
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] LikeCheckUpdate 发送收藏消息失败 %v\n", err)
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(msg)
		}
	}
}
