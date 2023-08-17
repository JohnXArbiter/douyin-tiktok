package mq

import (
	"douyin-tiktok/common/utils"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func (l *RabbitMQLogic) FavoriteUpdatePublisher(videoId, userId int64, isCancel int8) {
	var ticker = time.NewTicker(time.Second * 10)
	msg := &utils.VFMessage{
		VideoId:  videoId,
		UserId:   userId,
		Time:     time.Now().Local(),
		IsCancel: isCancel,
	}
	body, _ := l.svcCtx.Json.Marshal(msg)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++", l.svcCtx.RmqCore.Channel, l.svcCtx.RmqCore.Conn)
	publisher := utils.NewRabbitMQ(utils.VideoFavoriteQueue, utils.VideoFavoriteExchange, "cc", l.svcCtx.RmqCore.Conn, l.svcCtx.RmqCore.Channel)
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		logx.Infof("[RABBITMQ ERROR] FavoriteUpdatePublisher 发送点赞消息失败 %v\n", err)
		select {
		case <-ticker.C:
			if isCancel == 0 {
				go l.FavoriteCheck(msg)
			} else {
				go l.FavoriteCancelCheck(msg)
			}
		}
	}
}
