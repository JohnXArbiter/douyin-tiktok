// Code generated by goctl. DO NOT EDIT.
package types

type TokenReq struct {
	Token string `form:"token,optional"`
}

type FeedReq struct {
	LatestTime int64 `form:"latest_time"`
	TokenReq
}

type UserIdReq struct {
	TokenReq
	UserId int64 `form:"user_id,optional"`
}

type FavoriteActionReq struct {
	TokenReq
	VideoId    int64 `form:"video_id"`
	ActionType int32 `form:"action_type"`
}

type CommentActionReq struct {
	TokenReq
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text,optional"`
	CommentId   int64  `form:"comment_id,optional"`
}

type VideoIdReq struct {
	TokenReq
	VideoId int64 `form:"video_id"`
}

type PublishActionReq struct {
	TokenReq
	Title string `form:"title"`
}
