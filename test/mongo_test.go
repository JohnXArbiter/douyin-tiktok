package main

import (
	"context"
	"douyin-tiktok/common/utils"
	userModel "douyin-tiktok/service/user/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	mc := utils.InitMongo(utils.MongoConf{Url: "mongodb://admin:admin@43.143.241.157:27017"})

	collection := mc.Database("douyin_user").Collection("user_relation")

	ur := userModel.UserRelation{
		UserId: 12,
		Followers: []userModel.RelatedUser{{
			UserId: 123,
			Time:   time.Now().Unix(),
		}},
	}
	one, err := collection.InsertOne(context.Background(), &ur)

	var filter = bson.M{"_id": 15}
	followedUser := bson.M{"$addToSet": bson.M{
		"follow": bson.M{
			"followed_id": 1233,
			"time":        time.Now().Unix(),
		}},
	}
	optios := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(context.Background(), filter, followedUser, optios)
	fmt.Println(one, err)

	asd := &userModel.UserRelation{}
	cur, err := collection.Find(context.Background(), bson.M{"_id": 15})
	if cur.Next(context.Background()) {
		err = cur.Decode(asd)
		fmt.Println(one, asd)

	}
}
