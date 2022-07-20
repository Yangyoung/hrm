package controller

import (
	"hrm_go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//选择公司
func ChooseCompany(c *gin.Context) {
	var status bool

	companyList, err := model.ChooseCompany()
	if err != nil {
		status = false
	} else {
		status = true
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "msg": err, "data": companyList})
}
