package logic

import (
	"context"
	"douyin-tiktok/service/user/model"

	"douyin-tiktok/service/user/cmd/rpc/internal/svc"
	"douyin-tiktok/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoListLogic {
	return &GetInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoListLogic) GetInfoList(in *__.GetInfoListReq) (*__.GetInfoListResp, error) {
	var users, respUsers = make([]model.UserInfo, 0), make([]*__.User, 0)
	err := l.svcCtx.UserInfo.In("`id`", in.UserIds).
		Cols("`id`, `name`, `avatar`").Find(users)
	if err != nil {
		logx.Errorf("[DB ERROR] GetInfoList 批量查询用户信息失败 %v\n", err)
		return &__.GetInfoListResp{Code: -1}, err
	}

	for _, user := range users {
		respUser := &__.User{
			Id:     user.Id,
			Name:   user.Name,
			Avatar: &user.Avatar,
		}
		respUsers = append(respUsers, respUser)
	}
	resp := &__.GetInfoListResp{Users: respUsers}
	return resp, nil
}
