package routers

import (
	"github.com/gin-gonic/gin"
	"go-im/middlewares"
	"go-im/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)
	//验证码
	r.POST("/send/code", service.SendCode)

	auth := r.Group("/u", middlewares.Auth())
	auth.GET("/user/detail", service.UserDetail)
	return r
}
