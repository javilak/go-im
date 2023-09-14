package test

import (
	"context"
	"fmt"
	"go-im/conf"
	"go-im/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestFindOne(t *testing.T) {
	conf.Mongoini()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: conf.MongoUsr,
		Password: conf.MongoPw,
	}).ApplyURI(conf.MongoAddr))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database(conf.MongoDb)
	ub := &model.UserBasic{}
	err = db.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ub>>", ub)
}
