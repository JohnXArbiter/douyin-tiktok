package publish

import (
	"context"
	"douyin-tiktok/common/utils"
	__user "douyin-tiktok/service/user/cmd/rpc/types"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"douyin-tiktok/service/video/model"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPublishedUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPublishedUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPublishedUserIdLogic {
	return &ListPublishedUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPublishedUserIdLogic) ListPublishedUserId(req *types.UserIdReq) (map[string]interface{}, error) {
	var (
		userId     = req.UserId
		respChan   = make(chan *__user.GetInfoByIdResp)
		videoInfos []model.VideoInfo
	)

	// 请求用户 rpc 服务更新作品
	userRpcReq := &__user.UserIdReq{UserId: userId}
	go l.fetchUserInfo(respChan, userRpcReq)

	// 查询 video_info
	err := l.svcCtx.VideoInfo.Where("`user_id` = ?", userId).Desc("publish_at").Find(&videoInfos)
	if err != nil {
		logx.Errorf("[DB ERROR] ListPublishedUserId 查询用户id为：%v的视频列表失败 %v\n", userId, err)
		return nil, errors.New("数据获取失败！")
	}

	userRpcResp := <-respChan
	if userRpcResp.Code == -1 {
		return nil, errors.New("数据获取失败！")
	}

	// 返回指定格式
	for _, vi := range videoInfos {
		vi.Author = userRpcResp.User
	}

	resp := utils.GenOkResp()
	resp["video_list"] = videoInfos
	return resp, nil
}

func (l *ListPublishedUserIdLogic) fetchUserInfo(dataChan chan *__user.GetInfoByIdResp, req *__user.UserIdReq) {
	var resp, err = l.svcCtx.UserRpc.GetInfoById(l.ctx, req)
	if err != nil {
		logx.Errorf("[RPC ERROR] ListPublishedUserId 获取用户信息失败 %v\n", err)
		resp.Code = -1
	}

	dataChan <- resp
}
