package main

import (
	"hrm_go/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//登录
	r.POST("/loginCheck", controller.LoginCheck)

	//首页统计
	home := r.Group("/home")
	{
		//选择公司
		home.POST("/chooseCompany", controller.ChooseCompany)
		//查询公司人数
		home.POST("/getAllMemberNum", controller.GetAllMemberNum)
		//查询入职人数
		home.POST("/getMember", controller.GetMmeber)
		//查询当年培训次数
		home.POST("/getTrainNum", controller.GetTrainNum)
		//查询培训次数
		home.POST("/getRecruitNum", controller.GetRecruitNum)
	}

	//组织规划模块
	group := r.Group("/group")
	{
		//组织结构列表
		group.POST("/groupList", controller.GroupList)
		//员工列表
		group.POST("/memberList", controller.MemberList)
		//添加组织
		group.POST("/groupPlanAdd", controller.GroupPlanSave)
		//添加组织
		group.POST("/groupEdit", controller.GroupEdit)
	}

	//招聘模块
	recruit := r.Group("/recruit")
	{
		recruit.POST("/recruitList", controller.GetRecuitList)
		recruit.POST("/interviewList", controller.InterviewList)
	}

	//入职模块
	join := r.Group("/join")
	{
		join.POST("/joinList", controller.GetAllMemberList)
		join.POST("/deleteMember", controller.DeleteMember)
	}

	//培训模块
	training := r.Group("/training")
	{
		training.POST("/traininglist", controller.TrainingList)
	}
	r.Run(":1050")
}
