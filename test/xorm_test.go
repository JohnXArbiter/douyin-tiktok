package main

import (
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	videoModel "douyin-tiktok/service/video/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
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

func getVideoEngine() *xorm.Engine {
	return utils.InitXorm("mysql", utils.MysqlConf{Dsn: "root:123456789@tcp(43.143.241.157:3307)/douyin_video?charset=utf8mb4&parseTime=True&loc=Local"})
}

func TestAsd(t *testing.T) {
	engine := getVideoEngine()
	update, err := engine.SetExpr("comment_count", "comment_count + 1").ID("451113077114117").Update(&videoModel.VideoInfo{})
	fmt.Println(update, err)
}

// .Table("video_info")
func TestIncr(t *testing.T) {
	engine := getVideoEngine()
	if _, err := engine.Incr("`favorite_count`", 2).ID(451112560919813).
		Where("favorite_count >= ?", 0).Update(videoModel.VideoInfo{}); err != nil {
		logx.Errorf("[DB ERROR] FavoriteCheck 更新点赞数失败 %v\n", err)
	}
}

func TestFindVideoInfo(t *testing.T) {
	engine := getVideoEngine()
	ids := []int64{451112560919813, 451112674690309, 451113026004229, 451113077114117}
	videoInfos := make([]videoModel.VideoInfo, 0)
	if err := engine.In("`id`", ids).Find(&videoInfos); err != nil {
		logx.Errorf("[DB ERROR] ListFavoriteByUserId 批量查询videoInfo失败 %v\n", err)
	}
	for i := range videoInfos {
		log.Println(videoInfos[i])
	}
	fmt.Println(len(videoInfos))
}
