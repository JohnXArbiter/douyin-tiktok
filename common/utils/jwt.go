package utils

import (
	"douyin-tiktok/service/user/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var secretKey = []byte("dytt")

type DouyinTiktokClaims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenToken(user *model.UserInfo) (string, error) {
	var now = time.Now().Local()
	claims := &DouyinTiktokClaims{
		Id:   user.Id,
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用指定的签名方法和声明创建一个新token
	tokenString, err := token.SignedString(secretKey)          // 创建并返回一个完整的token（jwt）。令牌使用令牌中指定的签名 方法进行签名 。
	return tokenString, err
}

func ParseToken(token string) (*DouyinTiktokClaims, error) {
	var claims = new(DouyinTiktokClaims)
	// ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	// 第一个值是token ，
	// 第二个值是我们之后需要把解析的数据放入的地方，
	// 第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	tokenClaim, err := jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

	if !tokenClaim.Valid || err != nil {
		return nil, err
	}
	return claims, nil
}
