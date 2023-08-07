package relation

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowedUserByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowedUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowedUserByUserIdLogic {
	return &ListFollowedUserByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowedUserByUserIdLogic) ListFollowedUserByUserId(req *types.UserIdReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		id    = req.UserId
		idStr = strconv.FormatInt(id, 10)
		key   = utils.UserFollow + idStr
		//ids       []int64
		userInfos []model.UserInfo
	)

	zs1, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err == redis.Nil || len(zs1) == 0 { //
		fmt.Println(zs1, err)

		var follows = l.loadFollowedIdFromMongo(id)
		if follows == nil {
			return nil, errors.New("出错啦")
		}

		zs2 := make([]redis.Z, 0)
		for _, follow := range follows {
			z := redis.Z{
				Score:  float64(follow.Time),
				Member: follow.UserId,
			}
			zs2 = append(zs2, z)
		}
		if err = l.svcCtx.Redis.ZAdd(l.ctx, key, zs2...).Err(); err != nil {
			logx.Errorf("[REDIS ERROR] ListFollowedUserByUserId 关注列表存储redis失败 %v\n", err)
			zs1 = zs2 // 无法用redis排序就直接返回mongo中的结果
		}
	} else if err != nil {
		logx.Errorf("[REDIS ERROR] ListFollowedUserByUserId sth wrong with redis %v\n", err)
	}

	for _, z := range zs1 {
		id, _ := z.Member.(int64)
		userInfo := model.UserInfo{Id: id}
		userInfos = append(userInfos, userInfo)
	}

	if err := l.svcCtx.UserInfo.Find(&userInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFollowedUserByUserId 批量查询userInfo失败 %v\n", err)
		return nil, errors.New("出错啦")
	}

	resp := utils.GenOkResp()
	resp["user_list"] = userInfos
	return resp, nil
}

func (l *ListFollowedUserByUserIdLogic) loadFollowedIdFromMongo(id int64) []model.RelatedUsers {
	var follow model.UserRelation

	filter := bson.M{"_id": id}
	if err := l.svcCtx.UserRelation.FindOne(l.ctx, filter).Decode(&follow); err != nil {
		logx.Errorf("[MONGO ERROR] ListFollowedUserByUserId->loadFollowedIdFromMongo 查询关注文档失败 %v\n", err)
		return nil
	}
	return follow.Follows
}
