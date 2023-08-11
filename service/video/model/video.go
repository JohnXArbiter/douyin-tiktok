package model

type VideoInfo struct {
	Id              int64       `json:"id"`
	UserId          int64       `json:"user_id"`        // 用户信息外键
	PlayUrl         string      `json:"play_url"`       // 视频播放地址
	CoverUrl        string      `json:"cover_url"`      // 视频封面地址
	FavoriteCount   int64       `json:"favorite_count"` // 视频的点赞总数
	CommentCount    int64       `json:"comment_count"`  // 视频的点赞总数
	Title           string      `json:"title"`          // 视频标题
	PublishAt       int64       `json:"-"`
	VideoObjectName string      `json:"video_object_name"`              // 视频文件object名
	IsFavorite      bool        `json:"is_favorite,omitempty" xorm:"-"` // true-已点赞，false-未点赞
	Author          interface{} `json:"author,omitempty"`
}

type VideoFavorite struct {
	UserId         int64            `json:"_id" bson:"user_id"` // mongo 主键存 userId
	FavoriteVideos []FavoriteVideos `json:"favorite_videos" bson:"favorite_videos"`
}

type FavoriteVideos struct {
	VideoId int64 `json:"video_id" bson:"videoId"`
	Time    int64 `json:"time" bson:"time"`
}

type VideoComment struct {
	Id         int64       `json:"id"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date" xorm:"-"`
	CreateAt   int64       `json:"-"`
	UserId     int64       `json:"user_id"`
	VideoId    int64       `json:"video_id"`
	User       interface{} `json:"user,omitempty"`
}
