package model

import (
	"time"
)

type UserInfo struct {
	Id              int64  `json:"id" xorm:"id"`                             // id
	Username        string `json:"username" xorm:"username"`                 // 用户名（账号）
	Password        string `json:"password" xorm:"password"`                 // 密码
	Name            string `json:"name" xorm:"name"`                         // 用户名称
	Avatar          string `json:"avatar" xorm:"avatar"`                     // 用户头像
	BackgroundImage string `json:"background_image" xorm:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count" xorm:"favorite_count"`     // 点赞数量
	FollowCount     int64  `json:"follow_count" xorm:"follow_count"`         // 关注总数
	FollowerCount   int64  `json:"follower_count" xorm:"follower_count"`     // 粉丝总数
	Signature       string `json:"signature" xorm:"signature"`               // 个人简介
	TotalFavorited  int64  `json:"total_favorited" xorm:"total_favorited"`   // 获赞数量
	WorkCount       int64  `json:"work_count" xorm:"work_count"`             // 作品数量
	IsFollow        bool   `json:"is_follow" xorm:"-"`
}

type UserRelation struct {
	UserId  int64          `json:"user_id" bson:"_id"`
	Follows []RelatedUsers `json:"follows" bson:"follows"`
	Fans    []RelatedUsers `json:"fans" bson:"fans"`
}

type RelatedUsers struct {
	UserId int64 `json:"user_id" bson:"user_id"`
	Time   int64 `json:"time" bson:"time"`
}

type UserMessage struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`    // 发送者
	ToUserId   int64     `json:"to_user_id"` // 对方
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
