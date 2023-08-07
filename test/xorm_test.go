package main

import (
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	"fmt"
	"testing"
)

func TestUserInfo(t *testing.T) {
	engine := utils.InitXorm("mysql", utils.Postgresql{Dsn: "postgres:123456@localhost:5432/douyin_tiktok"})
	s := engine.Table("user_info")

	ui := &userModel.UserInfo{Username: "qjdlk", Password: "sadmsdnfjk"}

	insert, err := s.Insert(ui)
	fmt.Println(insert, err)

}
