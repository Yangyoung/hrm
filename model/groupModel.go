package model

import (
	"hrm_go/dao"
	"time"
)

type Group struct {
	Apply_id         int    `json:"apply_id"`
	Post_id          int    `json:"post_id"`
	Platform_id      int    `json:"platform_id"`
	Department_id    int    `json:"department_id"`
	Create_time      string `json:"create_time"`
	Department_name  string `json:"department_name"`
	Platform_name    string `json:"platform_name"`
	User_id          int    `json:"user_id"`
	Man_tag1         string `json:"man_tag1"`
	Post_num         int    `json:"post_num"`
	User_name        string `json:"user_name"`
	Company_id       int    `json:"company_id"`
	Post_name        string `json:"post_name"`
	Work_address     string `json:"work_address"`
	Work_time        string `json:"work_time"`
	Status           int    `json:"status"`
	Recruit_progress string `json:"recruit_progress"`
	Note             string `json:"note"`
	Check_type       int    `json:"check_type"`
	Real_name        string `json:"real_name"`
	Man_tag          string `json:"man_tag"`
	Position_name    string `json:"position_name"`
	Create_id        int    `json:"create_id"`
}

func GroupList(companyId, isAdmin, uid int, platformName, departName, applyName, checkType string) (groups []Group, err error) {
	db := dao.InitMysql()
	condition := db.Where("ag.company_id=? and ag.status=?", companyId, 1)
	//平台查询条件
	if platformName != "" {
		condition.Where("p.platform_name like ?", "%"+platformName+"%")
	}

	if isAdmin == 1 { //非管理员
		condition.Where("ag.create_id = ?", uid)
	}

	//部门查询条件
	if departName != "" {
		condition.Where("d.depart_name like ?", "%"+departName+"%")
	}

	//申请人条件
	if applyName != "" {
		condition.Where("u.real_name like ?", "%"+applyName+"%")
	}

	//审核状态条件
	if checkType != "" {
		condition.Where("ag.check_type like ?", "%"+checkType+"%")
	}

	err = condition.Table("hr_apply_group ag").Joins("left join hr_platform p on " +
		"p.platform_id=ag.platform_id").Joins("left join hr_depart d on " +
		"d.depart_id=ag.department_id").Joins("left join hr_user u on " +
		"u.uid=ag.user_id").Joins("left join hr_position hp on hp.position_id=ag.post_id").Select("ag.*,p.platform_name,d.depart_name,u.real_name,hp.position_name").Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return

}

//保存组织结构
func GroupSave(applyId, companyId, PostId, platformId, departmentId, userId, postNum int, departmentName, platformName, man_tag1, user_name, post_name, work_address, work_time, man_tag string) error {
	db := dao.InitMysql()
	var group Group
	if applyId > 0 {
		group = Group{
			Platform_id:     platformId,
			Post_id:         PostId,
			Department_id:   departmentId,
			Department_name: departmentName,
			Platform_name:   platformName,
			Man_tag1:        man_tag1,
			Post_num:        postNum,
			User_name:       user_name,
			Post_name:       post_name,
			Work_address:    work_address,
			Work_time:       work_time,
			Man_tag:         man_tag,
		}
		return db.Table("hr_apply_group").Where("apply_id=?", applyId).Updates(&group).Error
	} else {
		group = Group{
			Company_id:      companyId,
			Platform_id:     platformId,
			Post_id:         PostId,
			Department_id:   departmentId,
			Create_time:     time.Now().Format("2006-01-02 15:04:05"),
			Department_name: departmentName,
			Platform_name:   platformName,
			User_id:         userId,
			Man_tag1:        man_tag1,
			Post_num:        postNum,
			User_name:       user_name,
			Post_name:       post_name,
			Work_address:    work_address,
			Work_time:       work_time,
			Status:          1,
			Check_type:      1,
			Man_tag:         man_tag,
		}
		return db.Table("hr_apply_group").Omit("real_name,position_name,create_id").Create(&group).Error
	}

}

func GroupEdit(groupId int) (group Group, err error) {
	db := dao.InitMysql()
	err = db.Table("hr_apply_group").Where("apply_id =?", groupId).Take(&group).Error
	if err != nil {
		return Group{}, err
	}
	return
}
