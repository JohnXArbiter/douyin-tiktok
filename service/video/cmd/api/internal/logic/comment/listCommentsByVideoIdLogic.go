package comment

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

type ListCommentsByVideoIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentsByVideoIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsByVideoIdLogic {
	return &ListCommentsByVideoIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentsByVideoIdLogic) ListCommentsByVideoId(req *types.VideoIdReq) (map[string]interface{}, error) {
	var videoComments []model.VideoComment
	if err := l.svcCtx.VideoComment.Where("`video_id` = ?", req.VideoId).
		Desc("`create_at`").Find(&videoComments); err != nil {
		logx.Errorf("[DB ERROR] ListCommentsByVideoId 查询视频评论失败 %v\n", err)
		return nil, errors.New("获取失败")
	}

	// 1.userId 去重
	ids, usersMap := make([]int64, 0), make(map[int64]interface{})
	for _, vc := range videoComments {
		if _, ok := usersMap[vc.UserId]; !ok {
			ids = append(ids, vc.VideoId)
		}
		usersMap[vc.UserId] = struct{}{}
	}

	// 2.rpc 获取用户信息
	resp := utils.GenOkResp()
	getInfoListReq := __user.GetInfoListReq{TargetUserIds: ids, UserId: 0}
	getInfoListResp, err := l.svcCtx.UserRpc.GetInfoList(l.ctx, &getInfoListReq)
	if err != nil || getInfoListResp.Code != 0 {
		logx.Errorf("[DB ERROR] ListCommentsByVideoId rpc 获取用户信息失败 %v\n", err)
		resp["comment_list"] = videoComments
		return resp, nil
	}

	// 3.将用户信息放进 map，方便取
	for _, user := range getInfoListResp.Users {
		userMap := map[string]interface{}{
			"id":        user.Id,
			"name":      user.Name,
			"avatar":    user.Avatar,
			"is_follow": user.IsFollow, // TODO ？
		}
		usersMap[user.Id] = userMap
	}

	// 4.将用户信息取出
	for _, comment := range videoComments {
		comment.User = usersMap[comment.UserId]
	}
	resp["comment_list"] = videoComments
	return resp, nil
}
