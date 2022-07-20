package model

import (
	"hrm_go/dao"
)

type User struct {
	Uid          int    `json:"uid"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Status       int    `json:"status"`
	Is_admin     int    `json:"is_admin"`
	Company_id   int    `json:"company_id"`
	Company_name string `json:"company_name"`
	Real_name    string `json:"real_name"`
	Node_ids     string `json:"node_ids"`
	Update_time  int    `json:"update_time"`
}

//通过登录名查找用户
func FindUserByUsername(usr string) (u *User, err error) {
	var user User
	db := dao.InitMysql()

	errs := db.Raw("SELECT * FROM hr_user u LEFT JOIN hr_company c ON "+
		"u.company_id = c.company_id WHERE u.username =?", usr).Scan(&user).Error
	if errs != nil {
		return nil, errs
	} else {
		return &user, nil
	}
}
