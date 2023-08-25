package publish

import (
	"context"
	"douyin-tiktok/common/utils"
	__user "douyin-tiktok/service/user/cmd/rpc/types"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"douyin-tiktok/service/video/model"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPublishedVideosByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPublishedVideosByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPublishedVideosByUserIdLogic {
	return &ListPublishedVideosByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPublishedVideosByUserIdLogic) ListPublishedVideosByUserId(req *types.UserIdReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		userId     = loggedUser.Id
		rpcChan    = make(chan *__user.User)
		videoInfos []model.VideoInfo
	)

	// 请求用户 rpc 服务更新作品
	userRpcReq := &__user.GetInfoByIdReq{
		UserId:       userId,
		TargetUserId: req.UserId,
	}
	go l.fetchUserInfo(rpcChan, userRpcReq)

	// 查询 video_info
	err := l.svcCtx.VideoInfo.Where("`user_id` = ?", userId).Desc("`publish_at`").Find(&videoInfos)
	if err != nil {
		logx.Errorf("[DB ERROR] ListPublishedVideosByUserId 查询用户id为：%v的视频列表失败 %v\n", userId, err)
		return nil, errors.New("数据获取失败！")
	}
	fmt.Println(videoInfos)
	user := <-rpcChan

	for _, vi := range videoInfos {
		vi.Author = user
	}

	resp := utils.GenOkResp()
	resp["video_list"] = videoInfos
	return resp, nil
}

func (l *ListPublishedVideosByUserIdLogic) fetchUserInfo(rpcChan chan *__user.User, req *__user.GetInfoByIdReq) {
	var res *__user.User
	go func() {
		rpcChan <- res
	}()

	resp, err := l.svcCtx.UserRpc.GetInfoById(l.ctx, req)
	if err != nil || resp.Code != 0 {
		logx.Errorf("[RPC ERROR] fetchUserInfo rpc 获取用户信息失败 %v\n", err)
	} else {
		res = resp.User
	}
}
