package model

import (
	"hrm_go/dao"
)

type Train struct {
	Train_log_id  int `json:"train_log_id"`
	Platform_id   int `json:"platform_id"`
	Department_id int `json:"department_id"`
	Create_time   int `json:"create_time"`
	Company_id    int `json:"company_id"`
	Status        int `json:"status"`
	Train_score   int `json:"train_score"`
	Member_id     int `json:"member_id"`
}

//查询培训次数
func GetTrainNum(companyId string) (trains []Train, err error) {
	db := dao.InitMysql()
	err = db.Table("hr_train_log").Where("company_id=? and status=? and date_format(train_time,'%Y')=DATE_FORMAT(now(),'%Y')", companyId, 1).Find(&trains).Error
	if err != nil {
		return nil, err
	}
	return
}

//查询培训列表
func GetTrainList(companyId, platformName, departName, employee, employeeNo string, page int) (traininigList []Train, err error) {
	db := dao.InitMysql()
	condition := db.Table("hr_train_log").Where("company_id=? ", companyId)
	err = condition.Limit(10).Offset(page * 10).Find(&traininigList).Error
	if err != nil {
		return nil, err
	}
	return
}
