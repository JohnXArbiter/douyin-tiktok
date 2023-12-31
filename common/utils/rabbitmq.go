package utils

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

var MQUrl string

type (
	// RabbitMQConf rabbitMQ配置
	RabbitMQConf struct {
		RmqUrl string
	}

	RabbitmqCore struct {
		Conn    *amqp.Connection
		Channel *amqp.Channel
	}

	// RabbitMQ rabbitMQ结构体
	RabbitMQ struct {
		Conn      *amqp.Connection
		Channel   *amqp.Channel
		QueueName string // 队列名称
		Exchange  string // 交换机名称
		Key       string // bind Key 名称
		MQUrl     string // 连接信息
	}
)

func InitRabbitMQ(rc RabbitMQConf) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(rc.RmqUrl)
	logx.Infof("[RABBITMQ CONNECTING] InitRabbitMQ RmqUrl: %v\n", rc.RmqUrl)
	if err != nil {
		panic("[RABBITMQ ERROR] InitRabbitMQ 连接不到rabbitmq" + err.Error())
	}
	channel, err := conn.Channel()
	if err != nil {
		panic("[RABBITMQ ERROR] InitRabbitMQ 获取rabbitmq通道失败 " + err.Error())
	}
	return conn, channel
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string, conn *amqp.Connection, channel *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{
		Conn:      conn,
		Channel:   channel,
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MQUrl:     MQUrl,
	}
}

// PublishTopic 话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) error {
	// 发送消息
	err := r.Channel.Publish(r.Exchange, r.Key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return err
	}
	return nil
}

const (
	VideoFavoriteExchange       = "douyin_video_favorite_exchange"
	VideoFavoriteDeadExchange   = "douyin_video_favorite_exchange_dead"
	VideoFavoriteQueue          = "douyin_video_favorite"
	VideoFavoriteDeadQueue      = "douyin_video_favorite_dead"
	VideoFavoriteRoutingKey     = "douyin_vf_routing"
	VideoFavoriteDeadRoutingKey = "douyin_vf_dead_routing"
	VideoFavoriteTTL            = 10000
)

type (
	VFMessage struct {
		Time     int64 `json:"time"`
		UserId   int64 `json:"userId"`
		VideoId  int64 `json:"videoId"`
		IsCancel int8  `json:"isCancel"`
	}
)
