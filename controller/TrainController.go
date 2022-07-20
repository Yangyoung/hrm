package controller

import (
	"hrm_go/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTrainNum(c *gin.Context) {
	var status bool
	companyId := c.PostForm("companyId")
	trainData, err := model.GetTrainNum(companyId)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": trainData})
}

func TrainingList(c *gin.Context) {
	var status bool
	page, _ := strconv.Atoi(c.PostForm("page"))
	platformName := c.PostForm("platformName")
	departName := c.PostForm("departName")
	employee := c.PostForm("employeeName")
	employeeNo := c.PostForm("employeeNo")
	companyId := c.PostForm("companyId")

	trainingListData, err := model.GetTrainList(companyId, platformName, departName, employee, employeeNo, page)
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": trainingListData})

}
