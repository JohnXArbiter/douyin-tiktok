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

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageAction, logged *utils.JwtUser) error {
	var actionType = req.ActionType
	if actionType == 1 {
		msg := &model.UserMessage{
			UserId:     logged.Id,
			ToUserId:   req.ToUserId,
			Content:    req.Content,
			CreateTime: time.Now().Unix(),
		}
		if _, err := l.svcCtx.UserMessage.Insert(msg); err != nil {
			logx.Errorf("[DB ERROR] MessageAction 插入聊天记录失败 %v\n", err)
			return errors.New("出错啦")
		}
	}
	return nil
}
