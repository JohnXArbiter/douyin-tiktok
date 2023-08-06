package model

type UserInfo struct {
	Id              int64  `json:"id" xorm:"id"`                             // id
	Username        string `json:"username" xorm:"username"`                 // 用户名（账号）
	Password        string `json:"password" xorm:"password"`                 // 密码
	Name            string `json:"name" xorm:"name"`                         // 用户名称
	Avatar          string `json:"avatar" xorm:"avatar"`                     // 用户头像
	BackgroundImage string `json:"background_image" xorm:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count" xorm:"favorite_count"`     // 点赞数量
	FollowCount     int64  `json:"follow_count" xorm:"follow_count"`         // 关注总数
	FollowerCount   int64  `json:"follower_count" xorm:"follower_count"`     // 粉丝总数
	Signature       string `json:"signature" xorm:"signature"`               // 个人简介
	TotalFavorited  int64  `json:"total_favorited" xorm:"total_favorited"`   // 获赞数量
	WorkCount       int64  `json:"work_count" xorm:"work_count"`             // 作品数量
	IsFollow        bool   `json:"is_follow"`
}

//	type UserRelation struct {
//		Id            int64 `json:"id"`
//		UserId1       int64 `json:"user_id_1"`
//		UserId2       int64 `json:"user_id_2"`
//		Status        int8  `json:"to_user_id"` // 1：1关注2，2：2关注1，3：互关，0：双方都取关
//		User1UpdateAt int64 `json:"user_1_update_at"`
//		User2UpdateAt int64 `json:"user_2_update_at"`
//	}

type UserRelation struct {
	UserId int64 `bson:"_id" json:"user_id"`
	Follow struct {
		UserId int64 `json:"user_id"`
		Time   int64 `json:"time"`
	} `bson:"follow" json:"follow"`
}

type UserMessage struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`    // 发送者
	ToUserId int64  `json:"to_user_id"` // 对方
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}
