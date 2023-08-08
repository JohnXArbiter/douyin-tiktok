package middleware

import (
	"douyin-tiktok/common/utils"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
)

// JwtAuthenticate jwt校验中间件
func JwtAuthenticate(r *http.Request, token string) (*utils.JwtUser, error) {
	if token == "" { // 1.没有 token ×
		return nil, errors.New("请先登录")
	}

	claim, err := utils.ParseToken(token)
	if err != nil { // 2.如果 id 解析不出来，token 有问题 ×
		return nil, errors.New("身份认证错误或过期，请重新登录")
	}

	id := claim.Id
	key := utils.UserLogged + strconv.FormatInt(id, 10)

	redisToken, err := utils.UserServiceRedis.Get(r.Context(), key).Result()
	if err != nil { // 3.执行失败 ×
		logx.Errorf("[REDIS ERROR] JwtAuthenticate 获取 token string 失败 %v\n", err)
		return nil, errors.New("出错啦，请稍后")
	}
	if redisToken != token { // 4.没有 token，或者 id 错误 ×
		return nil, errors.New("身份认证过期，请重新登录")
	}

	utils.UserServiceRedis.Expire(r.Context(), key, 7*24*time.Hour)

	jwtUser := utils.JwtUser{
		Id:   id,
		Name: claim.Name,
	}
	return &jwtUser, nil
}
