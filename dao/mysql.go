package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	user     string   = "root"              //数据库用户名
	password string   = "xCwoKcy5sAPHZW)^b" //数据库密码
	address  string   = "172.17.16.3:3306"  //数据库ip+端口  127.0.0.1:3306
	dbname   string   = "hrm"               //数据库名称
	db       *gorm.DB                       //数据库实例
)

func InitMysql() (db *gorm.DB) {
	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbname + "?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}

	return DB

}
