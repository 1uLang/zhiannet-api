package server

import (
	awvs_server "github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	hids_server "github.com/1uLang/zhiannet-api/hids/server"
	nessus_server "github.com/1uLang/zhiannet-api/nessus/server"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
)

func SetApiDbPath(path string) {
	model.ApiDbPath = path
}
func GetApiDbPath() string {
	return model.ApiDbPath
}

func InitMysqlLink() {
	model.InitMysqlLink()

	//初始化建表
	initTable()
}
func initTable() {

	next_terminal_server.InitTable()

	awvs_server.InitTable()

	hids_server.InitTable()

	nessus_server.InitTable()

	platform_backup_server.InitTable() //平台数据备份记录表
}
