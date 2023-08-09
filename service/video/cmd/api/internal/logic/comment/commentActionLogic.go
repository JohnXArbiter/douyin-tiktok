package comment

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/model"
	"errors"
	"time"

	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		userId     = loggedUser.Id
		actionType = req.ActionType
	)

	if actionType == 1 {
		vc := &model.VideoComment{
			UserId:     userId,
			VideoId:    req.VideoId,
			Content:    req.CommentText,
			CreateAt:   time.Now().Unix(),
			CreateDate: time.Now().Format("01-02"),
		}
		if _, err := l.svcCtx.VideoComment.Insert(vc); err != nil {
			logx.Errorf("[DB ERROR] CommentAction 插入评论失败 %v\n", err)
			return nil, errors.New("评论失败")
		}
		resp := utils.GenOkResp()
		resp["comment"] = vc
		return resp, nil
	} else {
		id := req.CommentId
		if id == 0 || actionType != 2 {
			return nil, errors.New("没有这条评论")
		}
		vc := &model.VideoComment{Id: id, UserId: userId}
		if _, err := l.svcCtx.VideoComment.Delete(vc); err != nil {
			logx.Errorf("[DB ERROR] CommentAction 删除评论失败 %v\n", err)
			return nil, errors.New("删除失败")
		}
		return utils.GenOkResp(), nil
	}
}
