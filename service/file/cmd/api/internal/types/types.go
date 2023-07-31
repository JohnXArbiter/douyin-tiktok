// Code generated by goctl. DO NOT EDIT.
package types

type TokenReq struct {
	Token string `json:"token" form:"token"`
}

type PublishActionReq struct {
	TokenReq
	Title string `json:"title"`
}

type UserIdReq struct {
	TokenReq
	UserId int64 `json:"user_id" path:"user_id" form:"user_id"`
}
