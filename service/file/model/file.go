package model

import "time"

type FileVideo struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	VideoId    int64     `json:"video_id"` // 作品视频id
	Objectname string    `json:"objectname"`
	Url        string    `json:"url"`
	UploadAt   time.Time `json:"upload_at"`
}

type FileUser struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	Objectname string    `json:"objectname"`
	Type       int8      `json:"type"` // 头像or背景图
	Url        string    `json:"url"`
	UploadAt   time.Time `json:"upload_at"`
}

type FileCover struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	FileVideoId int64     `json:"file_video_id"` // 作品id
	Objectname  string    `json:"objectname"`
	Url         string    `json:"url"`
	UploadAt    time.Time `json:"upload_at"`
}
