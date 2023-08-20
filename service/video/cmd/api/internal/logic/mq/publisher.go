package mq

import (
	"douyin-tiktok/common/utils"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func (l *RabbitMQLogic) FavoriteUpdatePublisher(videoId, userId int64, isCancel int8) {
	var ticker = time.NewTicker(time.Second * 10)
	msg := &utils.VFMessage{
		VideoId:  videoId,
		UserId:   userId,
		Time:     time.Now().Unix(),
		IsCancel: isCancel,
	}
	body, _ := l.svcCtx.Json.Marshal(msg)
	publisher := utils.NewRabbitMQ(utils.VideoFavoriteQueue, utils.VideoFavoriteExchange,
		utils.VideoFavoriteRoutingKey, l.svcCtx.RmqCore.Conn, l.svcCtx.RmqCore.Channel)

	publishing := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	}
	if err := publisher.Channel.Publish(publisher.Exchange, publisher.Key,
		false, false, publishing); err != nil {
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
