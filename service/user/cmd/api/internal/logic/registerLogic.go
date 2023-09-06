package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/user/model"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.LoginReq) (map[string]interface{}, error) {
	var userInfo, err = l.validate(req)
	if err != nil {
		return nil, err
	}

	userInfo.Id = idgen.NextId()
	userInfo.Name = "user" + strconv.FormatInt(int64(rand.Int31()), 10)
	userInfo.Avatar = l.svcCtx.BgUrl + "/avatar" + strconv.Itoa(rand.Intn(3)) + ".jpg"
	userInfo.BackgroundImage = l.svcCtx.BgUrl + "/bg" + strconv.Itoa(rand.Intn(6)) + ".jpg"
	if _, err = l.svcCtx.UserInfo.Insert(userInfo); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.New("账号已经被抢走啦🫠")
		}
		return nil, errors.New("注册失败，请重试")
	}

	token, err := utils.GenToken(userInfo)
	if err != nil {
		return nil, errors.New("出错啦，请重试！")
	}

	key := utils.UserLogged + strconv.FormatInt(userInfo.Id, 10)
	if err = l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour).Err(); err != nil {
		logx.Errorf("[REDIS ERROR] Register 保存用户token失败，userid：%v %v\n", userInfo.Id, err)
		l.svcCtx.Redis.Set(l.ctx, key, token, 7*24*time.Hour) // 重试
	}

	resp := utils.GenOkResp()
	resp["user_id"] = userInfo.Id
	resp["token"] = token
	return resp, nil
}

// 校验账号密码
func (l *RegisterLogic) validate(req *types.LoginReq) (*model.UserInfo, error) {
	var (
		username = strings.TrimSpace(req.Username)
		password = []byte(strings.TrimSpace(req.Password))
		uPattern = "^[a-zA-Z0-9]{6,20}$"
		pPattern = "^[a-zA-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]{6,20}$"
	)

	uregex, _ := regexp.Compile(uPattern)
	if !uregex.MatchString(username) {
		return nil, errors.New("账号格式错误，只能包含数字和字母，5-20位")
	}

	pregex, _ := regexp.Compile(pPattern)
	if !pregex.MatchString(string(password)) {
		return nil, errors.New("密码格式错误，只能包含数字、字母和英文符号，6-20位")
	}

	password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userInfo := &model.UserInfo{
		Username: username,
		Password: string(password),
	}
	return userInfo, nil
}
