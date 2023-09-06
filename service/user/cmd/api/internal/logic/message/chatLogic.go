package message

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"douyin-tiktok/service/user/model"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"sort"
	"strconv"
	"time"
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
		latestTime = req.PreMsgTime
		key        = utils.UserMessageFlag + strconv.FormatInt(userId, 10) + ":" + strconv.FormatInt(toUserId, 10)
	)

	session := l.svcCtx.Xorm.Table("user_message")
	if latestTime == 0 {
		session.Where("(`user_id` = ? AND `to_user_id` = ?) OR (`user_id` = ? AND `to_user_id` = ?)",
			userId, toUserId, toUserId, userId).Desc("`create_time`")
	} else {
		val, err := l.svcCtx.Redis.Get(l.ctx, key).Result()
		if err != nil && err != redis.Nil {
			logx.Errorf("[DB ERROR] Chat redis获取聊天记录时间戳失败 %v\n", err)
			return nil, errors.New("出错了")
		} else if err == redis.Nil {
			return nil, nil
		}
		timeStamp, _ := strconv.ParseInt(val, 10, 64)
		session.Where("(`user_id` = ? AND `to_user_id` = ? AND `create_time` < ?) OR (`user_id` = ? AND `to_user_id` = ? AND `create_time` < ?)",
			userId, toUserId, timeStamp, toUserId, userId, timeStamp)
	}

	if err := session.Limit(20).Find(&msgs); err != nil {
		logx.Errorf("[DB ERROR] Chat 查询聊天记录失败 %v\n", err)
		return nil, errors.New("出错啦")
	}

	messages := model.UserMessages(msgs)
	sort.Sort(messages)

	//if len(msgs) > 0 {
	err := l.svcCtx.Redis.Set(l.ctx, key, msgs[0].CreateTime, 5*time.Second).Err()
	if err != nil {
		logx.Errorf("[DB ERROR] Chat redis记录聊天记录时间戳失败 %v\n", err)
		return nil, errors.New("出错啦")
	}
	//}

	if len(msgs) == 0 {
		msgs = append(msgs, model.UserMessage{CreateTime: 0})
	}

	resp := utils.GenOkResp()
	resp["message_list"] = msgs
	return resp, nil
}
