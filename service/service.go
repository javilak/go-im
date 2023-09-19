package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-im/helper"
	"go-im/model"
	"log"
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
	ub, err := model.GetUBbyAccountPw(account, pwd)
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

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claim")
	fmt.Println(u)
	Uc := u.(*helper.UserClaims)
	ub, err := model.GetUBbyId(Uc.Identity)
	if err != nil {
		log.Printf("[DB ERR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": ub,
	})

}

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	count, err := model.GetUBbyemail(email)
	if err != nil {
		fmt.Printf("[DB ERR: %v\n]", err)
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱已被使用",
		})
		return
	}
	err = helper.SendCode(email, "666")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "邮件已发送",
	})
}
