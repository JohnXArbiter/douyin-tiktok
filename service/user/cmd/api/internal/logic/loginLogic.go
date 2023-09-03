package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (map[string]interface{}, error) {
	var (
		username = strings.TrimSpace(req.Username)
		password = strings.TrimSpace(req.Password)
	)

	userInfo := &model.UserInfo{Username: username}
	has, err := l.svcCtx.UserInfo.Cols("`id`", "`username`", "`password`", "`name`").Get(userInfo)
	if err != nil || !has {
		logx.Errorf("[DB ERROR] ")
		return nil, errors.New("该账户不存在！")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password))
	if err != nil {
		return nil, errors.New("帐号或密码错误！")
	}

	token, err := utils.GenToken(userInfo)
	if err != nil {
		return nil, errors.New("出错啦，请重试！")
	}

	key := utils.UserLogged + strconv.FormatInt(userInfo.Id, 10)
	if err = l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] Login 保存用户token失败，userid：%v %v\n", userInfo.Id, err)
		l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}

	resp := utils.GenOkResp()
	resp["user_id"] = userInfo.Id
	resp["token"] = token
	return resp, nil
}
