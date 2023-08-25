package middleware

import (
	"context"
	"douyin-tiktok/common/utils"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"time"
)

var NoAuthList = map[string]struct{}{
	"/feed": {},
}

// JwtAuthenticate jwt校验中间件
func JwtAuthenticate(r *http.Request, token string) (*utils.JwtUser, error) {

	path := r.URL.Path
	if _, ok := NoAuthList[path[7:]]; ok {
		return parseJwt(r.Context(), token)
	}

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
	if err != nil && err != redis.Nil { // 3.执行失败 ×
		logx.Errorf("[REDIS ERROR] JwtAuthenticate 获取 token string 失败 %v\n", err)
		return nil, errors.New("出错啦，请稍后")
	}
	if redisToken != token || err == redis.Nil { // 4.没有 token，或者 id 错误 ×
		return nil, errors.New("身份认证过期，请重新登录")
	}

	utils.UserServiceRedis.Expire(r.Context(), key, 7*24*time.Hour)

	jwtUser := utils.JwtUser{
		Id:   id,
		Name: claim.Name,
	}
	return &jwtUser, nil
}

func parseJwt(ctx context.Context, token string) (*utils.JwtUser, error) {
	claim, err := utils.ParseToken(token)
	if err != nil { // 2.如果 id 解析不出来，token 有问题 ×
		return nil, nil
	}

	id := claim.Id
	key := utils.UserLogged + strconv.FormatInt(id, 10)

	redisToken, err := utils.UserServiceRedis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil || redisToken != token {
		return nil, nil
	}

	utils.UserServiceRedis.Expire(ctx, key, 7*24*time.Hour)
	user := &utils.JwtUser{
		Id:   id,
		Name: claim.Name,
	}
	return user, nil
}
