// Code generated by goctl. DO NOT EDIT.
package types

type TokenReq struct {
	Token string `form:"token"`
}

type UserIdReq struct {
	TokenReq
	UserId int64 `form:"user_id"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type FavoriteActionReq struct {
	TokenReq
	VideoId    int64 `form:"video_id"`
	ActionType int32 `form:"action_type"`
}

type RelationActionReq struct {
	TokenReq
	ToUserId   int64 `form:"to_user_id"`
	ActionType int32 `form:"action_type"`
}

type ChatReq struct {
	TokenReq
	ToUserId   int64 `form:"to_user_id"`
	PreMsgTime int64 `form:"pre_msg_time"`
}

type MessageAction struct {
	TokenReq
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
	Content    string `form:"content"`
}
