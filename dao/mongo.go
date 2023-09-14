package dao

import (
	"context"
	"go-im/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: conf.MongoUsr,
		Password: conf.MongoPw,
	}).ApplyURI(conf.MongoAddr))
	if err != nil {
		log.Println(err)
	}
	return client.Database(conf.MongoDb)
}
