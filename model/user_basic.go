package model

import (
	"context"
	"go-im/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity string `bson:"_id"`
	Account  string `bson:"account"`
	Passwd   string `bson:"password"`
	Nickname string `bson:"nickname"`
	Sex      int    `bson:"sex"`
	Email    string `bson:"email"`
	Avatar   string `bson:"avatar"`
	CreateAt int64  `bson:"created_at"`
	UpdateAt int64  `bson:"updated_at"`
}

func (UserBasic) CollectName() string {
	return "user_basic"
}

func GetUBbyAccountPw(account, passwd string) (*UserBasic, error) {
	ub := &UserBasic{}
	err := dao.Mongo.Collection(UserBasic{}.CollectName()).FindOne(context.Background(), bson.D{{Key: "account", Value: account},
		{Key: "password", Value: passwd}}).Decode(ub)
	return ub, err
}

func GetUBbyId(id primitive.ObjectID) (*UserBasic, error) {
	ub := &UserBasic{}
	err := dao.Mongo.Collection(UserBasic{}.CollectName()).FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(ub)
	return ub, err
}

func GetUBbyemail(email string) (int64, error) {
	return dao.Mongo.Collection(UserBasic{}.CollectName()).
		CountDocuments(context.Background(), bson.D{{Key: "email", Value: email}})

}
