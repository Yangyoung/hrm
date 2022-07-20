package model

import (
	"hrm_go/dao"
)

type Company struct {
	Company_id   int    `json:"company_id"`
	Company_name string `json:"company_name"`
}

//首页选择公司
func ChooseCompany() (companys []Company, err error) {
	db := dao.InitMysql()

	if err = db.Table("hr_company").Select("company_id,company_name").Find(&companys).Error; err != nil {
		return nil, err
	}
	return
}
