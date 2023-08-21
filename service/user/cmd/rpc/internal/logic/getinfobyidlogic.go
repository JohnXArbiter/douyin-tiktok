package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"fmt"
	"strconv"

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
		isFollow        bool
		userId          = in.UserId
		targetUserId    = in.TargetUserId
		targetUserIdStr = strconv.FormatInt(targetUserId, 10)
		key             = utils.UserFollow + strconv.FormatInt(userId, 10)
	)
	fmt.Println("dasadsfhjkdfhjk")

	var userInfo = &model.UserInfo{Id: in.UserId}
	has, err := l.svcCtx.UserInfo.Cols("`id`, `name`, `avatar`").Get(userInfo)
	if err != nil || !has {
		logx.Errorf("[DB ERROR] GetInfoById rpc根据id查询用户信息失败 %v\n", err)
		return &__.GetInfoByIdResp{Code: -1}, err
	}

	if userId != targetUserId {
		score, err := l.svcCtx.Redis.ZScore(l.ctx, key, targetUserIdStr).Result()
		if err != nil {
			logx.Errorf("[REDIS ERROR] ListFollowedUsersOrFans sth wrong with redis %v\n", err)
		}
		if score != 0 {
			isFollow = true
		} else {
			relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
			userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 1)
			if userRelation != nil && err == nil {
				_, _ = relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Followers, key)
				isFollow = true
			}
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

	resp := &__.GetInfoByIdResp{User: user}
	fmt.Println("23456789")
	return resp, nil
}
