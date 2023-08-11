package main

import (
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	"fmt"
	jsoniter "github.com/json-iterator/go"
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

func TestAsd(t *testing.T) {
	c := ContractJson{}
	cover := Cover{}
	second := Second{}
	sixth := Sixth{}
	ninth := Ninth{}
	appendix1 := Appendix1{}
	appendix2 := Appendix2{}
	arbitrate := &Arbitrate{}
	lawsuit := &Lawsuit{}
	c.Arbitrate = arbitrate
	c.Lawsuit = lawsuit
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	marshal, err := json.Marshal(c)
	covers, err := json.Marshal(cover)
	seconds, err := json.Marshal(second)
	sixths, err := json.Marshal(sixth)
	ninths, err := json.Marshal(ninth)
	a1, err := json.Marshal(appendix1)
	a2, err := json.Marshal(appendix2)

	fmt.Printf("%f\n", float64(2)/float64(5)*100)

	fmt.Println(string(marshal), err)
	fmt.Println(covers)
	fmt.Println(seconds)
	fmt.Println(sixths)
	fmt.Println(ninths)
	fmt.Println(a1)
	fmt.Println(a2)

}

type ContractJson struct {
	Cover     `json:"cover"`
	Second    `json:"second"`
	Sixth     `json:"sixth"`
	Ninth     `json:"ninth"`
	Appendix1 `json:"appendix1"`
	Appendix2 `json:"appendix2"`
}

type Cover struct {
	InfoProcessor string `json:"info_processor"`
	PAddress      string `json:"p_address"`
	PContactWay   string `json:"p_contact_way"`
	PContacts     string `json:"p_contacts"`
	PJob          string `json:"p_job"`
	InfoReceiver  string `json:"info_receiver"`
	RAddress      string `json:"r_address"`
	RContactWay   string `json:"r_contact_way"`
	RContacts     string `json:"r_contacts"`
	RJob          string `json:"r_job"`
	HasOrAppoint  string `json:"has_or_appoint"`
	CoverYear1    int    `json:"cover_year1"`
	CoverMonth1   int    `json:"cover_month1"`
	CoverDay1     int    `json:"cover_day1"`
	Contract      string `json:"contract"`
}

type Second struct {
	Obligation string `json:"obligation"`
}

type Sixth struct {
	ContactInfo string `json:"contact_info"`
}

type Ninth struct {
	NotifyAddress  string     `json:"notify_address"`
	NinthDay1      int        `json:"ninth_day1"`
	NinthDay2      int        `json:"ninth_day2"`
	ProcessType    int        `json:"process_type"`
	Arbitrate      *Arbitrate `json:"arbitrate,omitempty"`
	Lawsuit        *Lawsuit   `json:"lawsuit,omitempty"`
	InfoProcessor2 string     `json:"info_processor2"`
	NinthYear3     int        `json:"ninth_year3"`
	NinthMonth3    int        `json:"ninth_month3"`
	NinthDay3      int        `json:"ninth_day3"`
	InfoReceiver2  string     `json:"info_receiver2"`
	NinthYear4     int        `json:"ninth_year4"`
	NinthMonth4    int        `json:"ninth_month4"`
	NinthDay4      int        `json:"ninth_day4"`
}

type Arbitrate struct {
	Check1             int8   `json:"check1"`
	Check2             int8   `json:"check2"`
	Check3             int8   `json:"check3"`
	Check4             int8   `json:"check4"`
	Check5             int8   `json:"check5"`
	ArbitrateInstitute string `json:"arbitrate_institute"`
	ArbitrateAddress   string `json:"arbitrate_address"`
}

type Lawsuit struct {
	Num1        int    `json:"num1"`
	Num2        int    `json:"num2"`
	SignAddress string `json:"sign_address"`
}

type Appendix1 struct {
	Purpose            string `json:"purpose"`
	ProcessWay         string `json:"process_way"`
	Scale              string `json:"scale"`
	InfoType           string `json:"info_type"`
	SensitiveLevelType string `json:"sensitive_level_type"` // 出境敏感个人信息种类
	ReceiverInfo       string `json:"receiver_info"`        // 境外接收方信息
	TransmitType       string `json:"transmit_type"`        // 传输方式
	Appendix1Year1     int    `json:"appendix1_year1"`      // 出境后保存期限
	Appendix1Month1    int    `json:"appendix1_month1"`
	Appendix1Day1      int    `json:"appendix1_day1"`
	Appendix1Year2     int    `json:"appendix1_year2"`
	Appendix1Month2    int    `json:"appendix1_month2"`
	Appendix1Day2      int    `json:"appendix1_day2"`
	PreserveAddress    string `json:"preserve_address"` // 出境后保存地点
	Other              string `json:"other,omitempty"`  // 其他事项
}

type Appendix2 struct {
	Content string `json:"content,omitempty"` // 双方约定的其他条款（如需要）
}

// 暂存
type Contract struct {
	Id                int    `json:"id"`
	BusinessId        int    `json:"business_id"`
	CompaniesId       int    `json:"companies_id"`
	CreateTime        int    `json:"create_time"`
	Status            int    `json:"status"`             // 1：确认第一条，以此类推，12：填写完成
	Cover             string `json:"cover"`              // 封面json
	CoverProgress     int    `json:"cover_progress"`     // 封面填写进度
	Second            string `json:"second"`             // 第二条json
	SecondProgress    int    `json:"second_progress"`    // 第二条填写进度
	Sixth             string `json:"sixth"`              // 第六条json
	SixthProgress     int    `json:"sixth_progress"`     // 第六条填写进度
	Ninth             string `json:"ninth"`              // 第九条json
	NinthProgress     int    `json:"ninth_progress"`     // 第九条填写进度
	Appendix1         string `json:"appendix1"`          // 附页1json
	Appendix1Progress int    `json:"appendix1_progress"` // 附页1填写进度
	Appendix2         string `json:"appendix2"`          // 附页2json
	Appendix2Progress int    `json:"appendix2_progress"` // 附页2填写进度
}

//CREATE TABLE IF NOT EXISTS `tableName`
//(
//`id`                 INT NOT NULL,
//`business_id`        INT,
//`create_time`        INT,
//`status`             INT COMMENT '1：确认第一条，以此类推，12：填写完成',
//`cover`              TEXT COMMENT '封面json',
//`cover_progress`     INT COMMENT '封面填写进度',
//`second`             TEXT COMMENT '第二条json',
//`second_progress`    INT COMMENT '第二条填写进度',
//`sixth`              TEXT COMMENT '第六条json',
//`sixth_progress`     INT COMMENT '第六条填写进度',
//`ninth`              TEXT COMMENT '第九条json',
//`ninth_progress`     INT COMMENT '第九条填写进度',
//`appendix1`          TEXT COMMENT '附录1json',
//`appendix1_progress` INT COMMENT '附录1填写进度',
//`appendix2`          TEXT COMMENT '附录2json',
//`appendix2_progress` INT COMMENT '附录2填写进度',
//PRIMARY KEY (`id`)
//) ENGINE = InnoDB;
