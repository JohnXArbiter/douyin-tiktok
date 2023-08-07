package model

import "time"

type VideoInfo struct {
	Id              int64       `json:"id"`
	UserId          int64       `json:"user_id"`        // 用户信息外键
	PlayUrl         string      `json:"play_url"`       // 视频播放地址
	CoverUrl        string      `json:"cover_url"`      // 视频封面地址
	FavoriteCount   int64       `json:"favorite_count"` // 视频的点赞总数
	CommentCount    int64       `json:"comment_count"`  // 视频的点赞总数
	Title           string      `json:"title"`          // 视频标题
	PublishAt       time.Time   `json:"publish_at"`
	VideoObjectName string      `json:"video_object_name"` // 视频文件object名
	Author          interface{} `json:"author,omitempty"`
}

type VideoFavorite struct {
	Id       int64     `json:"id"`
	UserId   int64     `json:"user_id"`
	VideoId  int64     `json:"video_id"` // 喜欢的视频
	CreateAt time.Time `json:"create_at"`
}
