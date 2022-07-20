package controller

import (
	"hrm_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) {
	var status bool

	//接收用户名
	username := c.PostForm("username")
	user, err := model.FindUserByUsername(username)

	if err != nil {
		status = false
	} else {
		status = true
	}
	//返回json结果
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": user})

}
