package message

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"
	"time"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatReq, loggedUser *utils.JwtUser) (map[string]interface{}, error) {
	var (
		userId     = loggedUser.Id
		toUserId   = req.ToUserId
		msgs       = make([]model.UserMessage, 0)
		latestTime = time.Unix(req.PreMsgTime, 0).Local()
	)

	if err := l.svcCtx.UserMessage.Where("`user_id` = ? AND `to_user_id` = ? AND `create_time` > ?",
		userId, toUserId, latestTime).Desc("`create_time`").Find(&msgs); err != nil {
		logx.Errorf("[DB ERROR] Chat 查询聊天记录失败 %v\n", err)
		return nil, errors.New("出错啦")
	}

	resp := utils.GenOkResp()
	resp["message_list"] = msgs
	return resp, nil
}
