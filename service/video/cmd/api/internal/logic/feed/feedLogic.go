package feed

import (
	"context"
	"douyin-tiktok/common/utils"
	__user "douyin-tiktok/service/user/cmd/rpc/types"
	"douyin-tiktok/service/video/cmd/api/internal/logic/favorite"
	"douyin-tiktok/service/video/model"
	"errors"
	"fmt"
	"strconv"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		vis  []*model.VideoInfo
		resp = utils.GenOkResp()
	)
	if err := l.svcCtx.VideoInfo.Where("publish_at <= ?", req.LatestTime).
		Desc("publish_at").Limit(30).Find(&vis); err != nil {
		logx.Errorf("[DB ERROR] Feed 获取视频信息失败 %v\n", err)
		return nil, errors.New("出错啦😭")
	}

	if loggedUser == nil {
		resp["video_list"] = vis
		return resp, nil
	}

	var (
		userId    = loggedUser.Id
		userIdStr = strconv.FormatInt(userId, 10)
		uids      []int64
		rpcChan   = make(chan []*__user.User)
	)

	favoritedMap := make(map[int64]*model.VideoInfo)
	for _, vi := range vis {
		uids = append(uids, vi.UserId)
		favoritedMap[vi.Id] = vi
	}

	// 异步取用户信息
	go l.GetUserInfoListFromRpc(rpcChan, uids, userId)

	// 同时查点赞状态
	if userId != 0 {
		favoriteCommonLogic := favorite.NewFavoriteCommonLogic(l.ctx, l.svcCtx)
		videoIds, _ := favoriteCommonLogic.LoadIdsAndStore(utils.VideoFavorite+userIdStr, userId)
		for _, videoId := range videoIds {
			if vi, ok := favoritedMap[videoId]; ok {
				vi.IsFavorite = true
			}
		}
	}

	uis := <-rpcChan // 阻塞等待

	fmt.Println(uis)
	uisMap := make(map[int64]*__user.User)
	if uis != nil {
		for _, user := range uis {
			uisMap[user.Id] = user
		}
		for _, vi := range vis {
			vi.Author = uisMap[vi.UserId]
		}
	}

	resp["video_list"] = vis
	return resp, nil
}

func (l *FeedLogic) GetUserInfoListFromRpc(rpcChan chan []*__user.User, targetUserIds []int64, userId int64) {
	var res []*__user.User = nil
	defer func() {
		rpcChan <- res
	}()

	req := &__user.GetInfoListReq{
		TargetUserIds: targetUserIds,
		UserId:        userId,
	}
	resp, err := l.svcCtx.UserRpc.GetInfoList(l.ctx, req)
	if err == nil && resp.Code == 0 {
		res = resp.Users
	} else {
		logx.Errorf("[RPC ERROR] Feed->GetUserInfoListFromRpc 获取用户信息失败 %v\n", err)
	}
}
