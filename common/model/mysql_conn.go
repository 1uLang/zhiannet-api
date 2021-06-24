package model

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MysqlConn *gorm.DB

func init() {
	InitMysqlLink()
}
func InitMysqlLink() {
	var err error
	dsn := "root:mysql8@tcp(45.195.61.132:3306)/zhian-edges?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlConn, err = gorm.Open(gmysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表前缀
			SingularTable: true, //表名复数形式
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("mysql link err ")
	}
}
