package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-im/helper"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClams, err := helper.AnalyseToken(token)
		fmt.Println(userClams)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户验证失败",
			})
			return
		}
		c.Set("user_claim", userClams)
		c.Next()
	}
}
