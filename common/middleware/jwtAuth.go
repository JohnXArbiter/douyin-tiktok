package middleware

import (
	"context"
	"douyin-tiktok/common/utils"
	"errors"
	"net/http"
	"strconv"
)

// JwtAuthenticate jwt校验中间件
func JwtAuthenticate(r *http.Request, token string) error {
	if token == "" {
		return errors.New("请先登录")
	}

	claim, err := utils.ParseToken(token)
	if err != nil {
		return errors.New("身份认证错误或过期，请重新登录")
	}

	id := claim.Id
	key := utils.UserLogged + strconv.FormatInt(id, 10)

	redisToken, err := utils.UserServiceRedis.Get(r.Context(), key).Result()
	if redisToken != token {
		return errors.New("身份认证过期，请重新登录")
	}

	jwtUser := utils.JwtUser{
		Id:   id,
		Name: claim.Name,
	}
	ctx := context.WithValue(r.Context(), "user", jwtUser)
	r = r.WithContext(ctx)
	return nil
}
