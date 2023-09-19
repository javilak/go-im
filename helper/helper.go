package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"go-im/conf"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity primitive.ObjectID `json:"identity"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(identity, email string) (string, error) {
	Objectid, err2 := primitive.ObjectIDFromHex(identity)
	if err2 != nil {
		return "", err2
	}
	UserClaim := &UserClaims{
		Identity:         Objectid,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// GetCode
// 生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	u := uuid.New()
	return fmt.Sprintf("%x", u)
}

// SendCode 验证码邮件
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <" + conf.Mailac + ">"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS(conf.Mailaddr,
		smtp.PlainAuth("", conf.Mailac, conf.Mailpw, conf.Mailhost),
		&tls.Config{InsecureSkipVerify: true, ServerName: conf.Mailhost})
}
