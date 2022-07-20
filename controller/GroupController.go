package controller

import (
	"hrm_go/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查询组织架构列表
func GroupList(c *gin.Context) {
	var status bool
	//接收参数

	checkType := c.PostForm("checkType")
	companyId, _ := strconv.Atoi(c.PostForm("companyId"))
	isAdmin, _ := strconv.Atoi(c.PostForm("isAdmin"))
	uid, _ := strconv.Atoi(c.PostForm("uid"))
	platformName := c.PostForm("platformName")
	departName := c.PostForm("departName")
	applyName := c.PostForm("applyName")
	//查询
	groupList, err := model.GroupList(companyId, isAdmin, uid, platformName, departName, applyName, checkType)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": groupList})
}

//员工列表
func MemberList(c *gin.Context) {
	var status bool
	companyId, _ := strconv.Atoi(c.PostForm("companyId"))
	memberNumber := c.PostForm("memberNumber")
	platformName := c.PostForm("platformName")
	departName := c.PostForm("departName")
	memberName := c.PostForm("departName")
	members, err := model.GetAllMemberList(companyId, platformName, departName, memberName, memberNumber, "")
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": members})
}

//添加组织
func GroupPlanSave(c *gin.Context) {
	var status bool
	applyId, _ := strconv.Atoi(c.PostForm("applyId"))
	companyId, _ := strconv.Atoi(c.PostForm("company_id"))
	platformId, _ := strconv.Atoi(c.PostForm("platform_id"))
	platformName := c.PostForm("platform_name")
	departmentId, _ := strconv.Atoi(c.PostForm("department_id"))
	departmentName := c.PostForm("department_name")
	userId, _ := strconv.Atoi(c.PostForm("user_id"))
	userName := c.PostForm("user_name")
	postId, _ := strconv.Atoi(c.PostForm("post_id"))
	postName := c.PostForm("post_name")
	postNum, _ := strconv.Atoi(c.PostForm("post_num"))
	workAddress := c.PostForm("work_address")
	workTime := c.PostForm("work_time")
	manTag := c.PostForm("man_tag")
	manTag1 := c.PostForm("man_tag1")
	err := model.GroupSave(applyId, companyId, postId, platformId, departmentId, userId, postNum, departmentName, platformName, manTag1, userName, postName, workAddress, workTime, manTag)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": nil})
}

//编辑组织结构
func GroupEdit(c *gin.Context) {
	var status bool
	groupId, _ := strconv.Atoi(c.PostForm("groupId"))
	groupInfo, err := model.GroupEdit(groupId)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": groupInfo})
}
