package utils

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	VideoFavoriteExchange     = "douyin_video_favorite_exchange"
	VideoFavoriteDeadExchange = "douyin_video_favorite_exchange_dead"
	VideoFavoriteQueue        = "douyin_video_favorite"
	VideoFavoriteDeadQueue    = "douyin_video_favorite_dead"
	VideoLikeExchange         = "douyin_video_like_exchange"
	VideoLikeDeadExchange     = "douyin_video_like_exchange_dead"
	VideoLikeQueue            = "douyin_video_like"
	VideoLikeDeadQueue        = "douyin_video_like_dead"
)

type (
	VFMessage struct {
		RedisKey  string
		Time      time.Time
		UserId    int64
		IsCollect bool
	}

	VLMessage struct {
		RedisKey string
		Time     time.Time
		UserId   int64
	}
)
