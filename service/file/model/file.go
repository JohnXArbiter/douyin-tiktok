package model

type FileVideo struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"` // 作品视频id
	ObjectName string `json:"object_name"`
	Url        string `json:"url"`
	UploadAt   int64  `json:"upload_at"`
}

type FileUser struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	ObjectName string `json:"object_name"`
	Type       int8   `json:"type"` // 头像or背景图
	Url        string `json:"url"`
	UploadAt   int64  `json:"upload_at"`
}

type FileCover struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"file_video_id"` // 作品id
	ObjectName string `json:"object_name"`
	Url        string `json:"url"`
	UploadAt   int64  `json:"upload_at"`
}
