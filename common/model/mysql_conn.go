package model

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MysqlConn *gorm.DB

func InitMysqlLink() {
	var err error
	dsn := "root:root@tcp(192.168.168.17:3307)/deges?charset=utf8mb4&parseTime=True&loc=Local"
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
