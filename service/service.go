package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-im/helper"
	"go-im/model"
	"net/http"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("password")
	fmt.Println(account, pwd)
	if account == "" || pwd == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号或密码不能为空",
			"name": account,
			"pwd":  pwd,
		})
		return
	}
	ub, err := model.GetUBbyAccountPw(account, helper.GetMd5(pwd))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}
