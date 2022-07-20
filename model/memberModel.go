package model

import (
	"hrm_go/dao"
)

type Member struct {
	Member_id     string `json:"member_id"`
	Platform_id   string `json:"platform_id"`
	Department_id string `json:"department_id"`
	Member_name   string `json:"member_name"`
	Member_number string `json:"member_number"`
	Company_name  string `json:"company_name"`
	Depart_name   string `json:"depart_name"`
	Platform_name string `json:"platform_name"`
	Position_name string `json:"position_name"`
	Member_level  string `json:"member_level"`
	Join_time     string `json:"join_time"`
	Man_tag       string `json:"man_tag"`
	Man_tag1      string `json:"man_tag1"`
	Work_address  string `json:"work_address"`
	Leave_time    string `json:"leave_time"`
	Age           string `json:"age"`
	Is_dispute    string `json:"is_dispute"`
	Status        int    `json:"status"`
}

//查询总人数
func GetAllMemberNum(companyId string) (members []Member, err error) {
	db := dao.InitMysql()

	//db.table("hr_member").where().Find(&members)

	err = db.Table("hr_member").Select("member_id,platform_id,department_id,member_name,join_time").Where("company_id=?"+
		" and status=? and (leave_time is null or leave_time=?)", companyId, 1, "").Find(&members).Error
	if err != nil {
		return nil, err
	}
	return
}

//查询入职人数
func GetMember(companyId string) (members []Member, err error) {
	db := dao.InitMysql()
	err = db.Table("hr_member").Select("member_id").Find(&members).Error
	if err != nil {
		return nil, err
	}
	return
}

//人员列表
func GetAllMemberList(companyId int, platformName, departName, memberName, memberNumber, memberId string) (memers []Member, err error) {
	db := dao.InitMysql()
	condition := db.Where("ag.company_id=? and ag.status=?", companyId, 1)
	if platformName != "" {
		condition.Where("p.platform_name like ?", "%"+platformName+"%")
	}

	if departName != "" {
		condition.Where("d.depart_name like ?", "%"+departName+"%")
	}

	if memberName != "" {
		condition.Where("ag.member_name like ?", "%"+memberName+"%")
	}

	if memberNumber != "" {
		condition.Where("ag.member_number like ?", "%"+memberNumber+"%")
	}

	err = condition.Table("hr_member ag").Joins("left join hr_platform p on " +
		"p.platform_id=ag.platform_id").Joins("left join hr_depart d on " +
		"d.depart_id=ag.department_id").Joins("left join hr_position hp on " +
		"hp.position_id=ag.post_id").Joins("left join hr_company c on c.company_id=ag.company_id").Select("ag.*,p.platform_name,d.depart_name,hp.position_name,c.company_name").Find(&memers).Error
	if err != nil {
		return nil, err
	}
	return
}

// 删除员工
func DeleteMember(memberId, companyId int) error {
	db := dao.InitMysql()
	return db.Model(&Member{}).Table("hr_member").Where("member_id=? and company_id=?", memberId, companyId).Update("Status", 2).Error
}
