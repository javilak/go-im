package conf

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ini      *viper.Viper
	AppMode  string
	HttpPort string
	Db       string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPw     string
	DbName   string

	RedisDb     string
	RedisPw     int
	RedisAddr   string
	RedisDbName int

	MongoClient *mongo.Client
	MongoDb     string
	MongoPw     int
	MongoAddr   string
	MongoDbName string
	MongoPort   string
)

func init() {
	ini = viper.New()
	ini.SetConfigName("conf")
	ini.SetConfigType("yaml")
	ini.AddConfigPath("./conf/")
	err := ini.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("read init err: %s", err.Error()))
	}
	RedisIni()
	mongoini()
}

func RedisIni() {
	RedisAddr = ini.GetString("Redis.Addr")
	RedisPw = ini.GetInt("Redis.Pw")
	RedisDb = ini.GetString("Redis.Db")
	RedisDbName = ini.GetInt("Redis.Name")
}
func mongoini() {
	MongoAddr = ini.GetString("Mongo.Addr")
	MongoPw = ini.GetInt("Mongo.Pw")
	MongoDb = ini.GetString("Mongo.Db")
	MongoDbName = ini.GetString("Mongo.Name")
	MongoPort = ini.GetString("Mongo.Port")

}

func Mongodb() {
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoAddr + ":" + MongoPort)
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	logrus.Info("mongodb connected success!")

}
