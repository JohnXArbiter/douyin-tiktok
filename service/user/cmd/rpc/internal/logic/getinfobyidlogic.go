package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
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

	var userInfo = &model.UserInfo{Id: in.UserId}
	err := l.svcCtx.UserInfo.Cols("`id`, `name`, `avatar`").Find(userInfo)
	if err != nil {
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
		}
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		userRelation, err := relationCommonLogic.LoadIdsFromMongo(userId, 1)
		if userRelation == nil && err == nil {
			isFollow = true
		} else if err == nil {
			zs, err := relationCommonLogic.StoreRelatedUsers2Redis(userRelation.Follows, key)
			if err == nil {
				if err = l.svcCtx.Redis.ZAdd(l.ctx, key, zs...).Err(); err != nil {
					logx.Errorf("[REDIS ERROR] StoreRelatedUsers2Redis 关注列表存储redis失败 %v\n", err)
				} else {
					score = l.svcCtx.Redis.ZScore(l.ctx, key, targetUserIdStr).Val()
					if score != 0 {
						isFollow = true
					}
				}
			}
		} else {
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
