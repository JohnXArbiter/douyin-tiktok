package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"github.com/redis/go-redis/v9"
	"strconv"

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
	var (
		users     = make([]model.UserInfo, 0)
		respUsers = make([]*__.User, 0)
		userId    = in.UserId
		key       = utils.UserFollow + strconv.FormatInt(userId, 10)
		zs        []redis.Z
	)

	// 1.先查出用户信息
	if err := l.svcCtx.UserInfo.In("`id`", in.TargetUserIds).
		Cols("`id`, `name`, `avatar`").Find(users); err != nil {
		logx.Errorf("[DB ERROR] GetInfoList 批量查询用户信息失败 %v\n", err)
		return &__.GetInfoListResp{Code: -1}, err
	}

	// 2.再查登录用户的关注列表
	if exists, err := l.svcCtx.Redis.Exists(l.ctx, key).Result(); err != nil {
		logx.Errorf("[REDIS ERROR] GetInfoList sth wrong with redis %v\n", err)
	} else if exists != 1 {
		relationCommonLogic := NewRelationCommonLogic(l.ctx, l.svcCtx)
		zs = relationCommonLogic.LoadIdsAndStore(userId, 1, key)
	} else {
		zs, err = l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
		if err != nil {
			logx.Errorf("[REDIS ERROR] GetInfoList sth wrong with redis %v\n", err)
		}
	}

	// 3.将关注的用户存到 map
	idsMap := make(map[int64]bool)
	for _, z := range zs {
		id, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		idsMap[id] = true
	}

	// 4.map 为 true 的就是关注过的
	for _, user := range users {
		respUser := &__.User{
			Id:       user.Id,
			Name:     user.Name,
			Avatar:   &user.Avatar,
			IsFollow: idsMap[user.Id],
		}
		respUsers = append(respUsers, respUser)
	}
	resp := &__.GetInfoListResp{Users: respUsers}
	return resp, nil
}
