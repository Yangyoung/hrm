package controller

import (
	"hrm_go/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//招聘统计
func GetRecruitNum(c *gin.Context) {
	var status bool
	companyId := c.PostForm("companyId")
	recruit, err := model.GetRecruitNum(companyId)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "data": recruit, "msg": err})
}

//招聘进展
func GetRecuitList(c *gin.Context) {
	var status bool
	companyId, _ := strconv.Atoi(c.PostForm("compoanyId"))
	platformName := c.PostForm("platformName")
	departName := c.PostForm("departName")
	uName := c.PostForm("uName")
	position := c.PostForm("position")
	dateStart := c.PostForm("dateStart")
	dateEnd := c.PostForm("dateEnd")
	interviewProcessList, err := model.GetInterviewProcessList(companyId, platformName, departName, uName, position, dateStart, dateEnd)
	if err != nil {
		status = false
	} else {
		status = true
	}

	c.JSON(http.StatusOK, gin.H{"status": status, "data": interviewProcessList, "msg": err})
}

//招聘人员列表
func InterviewList(c *gin.Context) {
	var status bool
	companyId, _ := strconv.Atoi(c.PostForm("companyId"))
	platformName := c.PostForm("platformName")
	interviewName := c.PostForm("interviewName")
	departName := c.PostForm("departName")
	positionName := c.PostForm("positionName")
	start := c.PostForm("timeStart")
	end := c.PostForm("timeEnd")

	interviews, err := model.GetAllInterview(companyId, platformName, departName, interviewName, positionName, start, end)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "data": interviews, "msg": err})
}
