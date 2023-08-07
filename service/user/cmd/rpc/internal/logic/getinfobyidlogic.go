package logic

import (
	"context"
	"douyin-tiktok/service/user/model"

	"douyin-tiktok/service/user/cmd/rpc/internal/svc"
	"douyin-tiktok/service/user/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoByIdLogic {
	return &GetInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoByIdLogic) GetInfoById(in *__.GetInfoByIdReq) (*__.GetInfoByIdResp, error) {
	var (
		isFollow     bool
		userId       = in.UserId
		targetUserId = in.TargetUserId
	)

	var userInfo = &model.UserInfo{Id: in.UserId}
	err := l.svcCtx.UserInfo.Find(userInfo)
	if err != nil {
		logx.Errorf("[DB ERROR] GetInfoById rpc根据id查询用户信息失败 %v\n", err)
		return &__.GetInfoByIdResp{Code: -1}, err
	}

	if userId != targetUserId {
		isFollow, err = l.svcCtx.UserRelation.Where("`user_id` = ? AND `to_user_id` = ?", userId, targetUserId).Exist()
		if err != nil {
			logx.Errorf("[DB ERROR] GetInfoById rpc查询关注记录失败 %v\n", err)
		}
	}

	user := &__.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     &userInfo.FollowCount,
		FollowerCount:   &userInfo.FollowerCount,
		Avatar:          &userInfo.Avatar,
		IsFollow:        isFollow,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &userInfo.TotalFavorited,
		WorkCount:       &userInfo.WorkCount,
		FavoriteCount:   &userInfo.FavoriteCount,
	}

	resp := &__.GetInfoByIdResp{
		Code: 0,
		User: user,
	}
	return resp, nil
}