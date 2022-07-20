package model

import "hrm_go/dao"

type Interview struct {
	Interview_id   int    `json:"interview_id"`
	Person_name    string `json:"person_name"`
	Interview_time string `json:"interview_time"`
	Phone          string `json:"phone"`
	Company_name   string `json:"company_name"`
	Platform_name  string `json:"platform_name"`
	Depart_name    string `json:"depart_name"`
	Position_name  string `json:"position_name"`
	Age            string `json:"age"`
	Grade_num      string `json:"grade_num"`
	Is_offer       string `json:"is_offer"`
	Is_join        string `json:"is_join"`
}

//招聘统计
func GetRecruitNum(conpanyId string) (inerview []Interview, err error) {
	db := dao.InitMysql()
	err = db.Table("hr_interview").Where("company_id=? and status=?", conpanyId, 1).Find(&inerview).Error
	if err != nil {
		return nil, err
	}
	return
}

//面试进展列表
func GetInterviewProcessList(companyId int, platformName, departName, uName, positionName, start, end string) (interviewList []Interview, err error) {
	db := dao.InitMysql()
	condition := db.Where("ag.company_id=? and ag.status=1", companyId)
	//平台
	if platformName != "" {
		condition.Where("p.platform_name like ?", "%"+platformName+"%")
	}
	err = condition.Table("hr_interview ag").Joins("left join hr_platform p on p.platform_id=ag.platform_id").Joins("left join hr_depart d on " +
		"d.depart_id=ag.department_id").Joins("left join hr_position hp on hp.position_id=ag.post_id").Joins("left join hr_company hc on hc.company_id=ag.company_id").Order("ag.interview_id desc").Select("ag.*,p.platform_name,d.depart_name,hp.position_name,hc.company_name").Find(&interviewList).Error
	if err != nil {
		return nil, err
	}
	return
}

//面试列表
func GetAllInterview(companyId int, platformName, departName, interviewName, positionName, start, end string) (interviews []Interview, err error) {
	db := dao.InitMysql()
	condition := db.Where("ag.company_id=? and ag.status=1", companyId)
	//平台
	if platformName != "" {
		condition.Where("p.platform_name like ?", platformName+"%")
	}
	err = condition.Table("hr_interview ag").Joins("left join hr_platform p on p.platform_id=ag.platform_id").Joins("left join hr_depart d on " +
		"d.depart_id=ag.department_id").Joins("left join hr_position hp on hp.position_id=ag.post_id").Joins("left join hr_company hc on hc.company_id=ag.company_id").Select("ag.*,p.platform_name,d.depart_name,hp.position_name,hc.company_name").Find(&interviews).Error
	if err != nil {
		return nil, err
	}
	return
}
