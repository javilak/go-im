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
	RedisPw     string
	RedisAddr   string
	RedisDbName int

	MongoClient *mongo.Client
	MongoDb     string
	MongoPw     string
	MongoAddr   string
	MongoDbName string
	MongoPort   string
	MongoUsr    string

	Mailpw   string
	Mailac   string
	Mailaddr string
)

func init() {
	ini = viper.New()
	ini.SetConfigName("conf")
	ini.SetConfigType("yaml")
	ini.AddConfigPath("./conf/")
	ini.AddConfigPath("../conf/")
	err := ini.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("read init err: %s", err.Error()))
	}
	RedisIni()
	Mongoini()
	mailini()
}

func RedisIni() {
	RedisAddr = ini.GetString("Redis.Addr")
	RedisPw = ini.GetString("Redis.Pw")
	RedisDb = ini.GetString("Redis.Db")
	RedisDbName = ini.GetInt("Redis.Name")
}
func Mongoini() {
	MongoAddr = ini.GetString("Mongo.Addr")
	MongoPw = ini.GetString("Mongo.Pw")
	MongoUsr = ini.GetString("Mongo.Usr")
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

func mailini() {
	Mailaddr = ini.GetString("mail.addr")
	Mailac = ini.GetString("mail.accout")
	Mailpw = ini.GetString("mail.passwd")
}
