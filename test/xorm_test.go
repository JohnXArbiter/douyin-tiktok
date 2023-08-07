package main

import (
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	"fmt"
	"testing"
)

func TestUserInfo(t *testing.T) {
	engine := utils.InitXorm("mysql", utils.Mysql{Dsn: "root:123456@tcp(43.143.241.157:3306)/douyin_user?charset=utf8mb4&parseTime=True&loc=Local"})

	s := engine.Table("user_info")

	ui := &userModel.UserInfo{Username: "qjdlk", Password: "sadmsdnfjk"}

	insert, err := s.Insert(ui)
	fmt.Println(insert, err)

}
