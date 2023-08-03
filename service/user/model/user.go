package model

import "time"

type UserInfo struct {
	Id              int64  `json:"id"`               // id
	Username        string `json:"username"`         // 用户名（账号）
	Password        string `json:"password"`         // 密码
	Name            string `json:"name"`             // 用户名称
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 点赞数量
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数量
	IsFollow        bool   `json:"is_follow,omitempty"`
}

type UserRelation struct {
	Id       int64     `json:"id"`
	UserId   int64     `json:"user_id"`
	ToUserId int64     `json:"to_user_id"` // 关注的用户
	CreateAt time.Time `json:"create_at"`
}

type UserMessage struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`    // 发送者
	ToUserId int64  `json:"to_user_id"` // 对方
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}
