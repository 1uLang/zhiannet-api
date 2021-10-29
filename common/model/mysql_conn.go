package model

import (
	yaml "gopkg.in/yaml.v2"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
)

var MysqlConn *gorm.DB
var AuditMysqlConn *gorm.DB

type (
	DBConfig struct {
		Dbs     Dbs `yaml:"dbs"`
		Auditdb Dbs `yaml:"auditdb"`
	}

	// Dbs
	Dbs struct {
		Prod Prod `yaml:"prod"`
	}

	// Prod
	Prod struct {
		Dsn    string `yaml:"dsn"`
		Prefix string `yaml:"prefix"`
		Driver string `yaml:"driver"`
	}
)

var ApiDbPath = "./build/configs/api_db.yaml"
var DSN string

//func init() {
//	InitMysqlLink()
//}
func InitMysqlLink() {
	var err error
	var yamlFile []byte
	conf := new(DBConfig)
	//yamlFile, err = ioutil.ReadFile(ApiDbPath)
	////yamlFile, err = ioutil.ReadFile("./build/configs/api_db.yaml")
	yamlFile, err = ioutil.ReadFile("/Users/dp/zhian/zhiannet-edge-line/EdgeAdmin/build/configs/api_db.yaml")
	if err != nil {
		//panic(fmt.Errorf("zhiannet package mysql link yamlFile.Get err #%v ", err))
		return
	}
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		//panic(fmt.Errorf("zhiannet package mysql link yaml.Unmarshal err %v", err))
		return
	}
	//dsn := "root:mysql8@tcp(45.195.61.132:3306)/zhian-edges?charset=utf8mb4&parseTime=True&loc=Local"
	DSN = conf.Dbs.Prod.Dsn
	MysqlConn, err = gorm.Open(gmysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表前缀
			SingularTable: true, //表名复数形式
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		//panic(fmt.Errorf("zhiannet-api package link mysql err %v", err))
	}

	//dsn = "root:mysql8@tcp(45.195.61.132:3306)/gfast_open_test?charset=utf8mb4&parseTime=True&loc=Local"
	//DSN = conf.Auditdb.Prod.Dsn
	//AuditMysqlConn, err = gorm.Open(gmysql.Open(DSN), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   "",   //表前缀
	//		SingularTable: true, //表名复数形式
	//	},
	//	Logger: logger.Default.LogMode(logger.Silent),
	//})
	//if err != nil {
	//	//panic("审计系统 mysql link err ")
	//}
}
