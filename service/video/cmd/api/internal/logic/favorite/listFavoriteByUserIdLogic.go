package favorite

import (
	"context"
	"douyin-tiktok/common/utils"
	__ "douyin-tiktok/service/user/cmd/rpc/types"
	"douyin-tiktok/service/video/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFavoriteByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFavoriteByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFavoriteByUserIdLogic {
	return &ListFavoriteByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFavoriteByUserIdLogic) ListFavoriteByUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId    = req.UserId
		userIdStr = strconv.FormatInt(userId, 10)
		key       = utils.VideoFavorite + userIdStr
		resp      = utils.GenOkResp()
	)

	favoriteCommonLogic := NewFavoriteCommonLogic(l.ctx, l.svcCtx)

	zs, err := l.svcCtx.Redis.ZRevRangeWithScores(l.ctx, key, 0, -1).Result()
	if err != nil && err != redis.Nil {
		logx.Errorf("[REDIS ERROR] ListFavoriteVideos sth wrong with redis %v\n", err)
	} else if err == redis.Nil || len(zs) == 0 {
		zs, err = favoriteCommonLogic.LoadIdsAndStore(key, userId)
		if err != nil {
			return nil, err
		}
	}

	videoIds := make([]int64, 0)
	for _, z := range zs {
		id, _ := strconv.ParseInt(z.Member.(string), 10, 64)
		videoIds = append(videoIds, id)
	}

	videoInfos := make([]model.VideoInfo, 0)
	if err = l.svcCtx.VideoInfo.In("`id`", videoIds).Find(&videoInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFavoriteByUserId 批量查询videoInfo失败 %v\n", err)
		return nil, errors.New("出错啦")
	}

	userIds, visMap := make([]int64, 0), make(map[int64]*model.VideoInfo)
	for _, videoInfo := range videoInfos {
		userIds = append(userIds, videoInfo.UserId)
		visMap[videoInfo.Id] = &videoInfo
	}

	getInfoListReq := &__.GetInfoListReq{UserIds: userIds}
	getInfoListResp, err := l.svcCtx.UserRpc.GetInfoList(l.ctx, getInfoListReq)
	if err != nil || getInfoListResp.Code != 0 {
		logx.Errorf("[RPC ERROR] ListFavoriteByUserId 获取用户信息失败 %v\n", err)
	}

	uisMap := make(map[int64]*__.User)
	for _, userInfo := range getInfoListResp.Users {
		uisMap[userInfo.Id] = userInfo
	}

	videoList := make([]*model.VideoInfo, 0)
	for _, id := range videoIds {
		videoInfo := visMap[id]
		videoInfo.Author = uisMap[videoInfo.UserId]
		videoList = append(videoList, videoInfo)
	}

	resp["video_list"] = videoList
	return resp, nil
}
