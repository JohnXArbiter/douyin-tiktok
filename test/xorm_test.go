package main

import (
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	"fmt"
	"testing"
	"time"
	"xorm.io/xorm"
)

func getUserEngine() *xorm.Engine {
	return utils.InitXorm("mysql", utils.MysqlConf{Dsn: "root:123456@tcp(43.143.241.157:3306)/douyin_user?charset=utf8mb4&parseTime=True&loc=Local"})
}

func TestUserInfo(t *testing.T) {
	engine := utils.InitXorm("mysql", utils.MysqlConf{Dsn: "root:123456@tcp(43.143.241.157:3306)/douyin_user?charset=utf8mb4&parseTime=True&loc=Local"})

	s := engine.Table("user_info")

	//ui := &userModel.UserInfo{Username: "qjdlk", Password: "sadmsdnfjk"}
	//
	//insert, err := s.Insert(ui)
	//fmt.Println(insert, err)

	ui := &userModel.UserInfo{Id: 3}
	get, err := s.Cols("id, name, avatar").Get(ui)
	fmt.Printf("%#v %v %v\n", ui, err, get)
}

func TestTime(t *testing.T) {
	e := getUserEngine()
	s := e.Table("user_message")

	var timee int64 = 1691679993
	//var msgs1 []userModel.UserMessage
	//for i := 0; i < 4; i++ {
	//	now := time.Now().Local()
	//	msg := userModel.UserMessage{
	//		UserId:     1,
	//		ToUserId:   2,
	//		Content:    strconv.FormatInt(rand.Int63(), 10),
	//		CreateTime: now,
	//	}
	//	if i == 2 {
	//		timee = now.Unix()
	//	}
	//	msgs1 = append(msgs1, msg)
	//	//time.Sleep(time.Second * 1)
	//}
	//
	//insert, err := s.Insert(msgs1)
	//fmt.Println("插插插插插插插插插插插插插插插", timee, insert, err)
	local := time.Unix(timee, 0).Local()
	var msgs2 []userModel.UserMessage
	err := s.Where("`user_id` = ? AND `to_user_id` = ? AND `create_time` > ?",
		1, 2, local).Desc("`create_time`").Find(&msgs2)
	fmt.Println("查查查查查查查查查查查查查查查查查查查查", local, timee, err)
	fmt.Println(msgs2)
}
