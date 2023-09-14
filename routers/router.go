package routers

import (
	"github.com/gin-gonic/gin"
	"go-im/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)
	return r
}
