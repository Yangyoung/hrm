package controller

import (
	"hrm_go/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询总人数
func GetAllMemberNum(c *gin.Context) {
	var status bool
	companyId := c.PostForm("companyId")
	members, err := model.GetAllMemberNum(companyId)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": members})
}

//查询入职人数
func GetMmeber(c *gin.Context) {
	var status bool
	companyId := c.PostForm("companyId")
	members, err := model.GetMember(companyId)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": members})
}

//获取员工列表
func GetAllMemberList(c *gin.Context) {
	var status bool
	companyId, _ := strconv.Atoi(c.PostForm("companyId"))
	platformName := c.PostForm("platformName")
	departName := c.PostForm("departName")
	memberName := c.PostForm("memberName")
	memberNumber := c.PostForm("memberNumber")
	memberId := c.PostForm("memberId")

	memberList, err := model.GetAllMemberList(companyId, platformName, departName, memberName, memberNumber, memberId)

	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": memberList})
}

//删除员工
func DeleteMember(c *gin.Context) {
	var status bool = true
	memberId, _ := strconv.Atoi(c.PostForm("memberId"))
	companyId, _ := strconv.Atoi(c.PostForm("companyId"))

	err := model.DeleteMember(memberId, companyId)
	if err != nil {
		status = false
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": nil})
}
