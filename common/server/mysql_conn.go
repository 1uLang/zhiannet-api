package server

import (
	agent_model "github.com/1uLang/zhiannet-api/agent/model"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user_relation"
	awvs_server "github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/common/server/platform_backup_server"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	hids_server "github.com/1uLang/zhiannet-api/hids/server"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_list"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_notice"
	nessus_server "github.com/1uLang/zhiannet-api/nessus/server"
	next_terminal_server "github.com/1uLang/zhiannet-api/next-terminal/server"
	nextcloud_model "github.com/1uLang/zhiannet-api/nextcloud/model"
	resmon_model "github.com/1uLang/zhiannet-api/resmon/model"
	wazuh_server "github.com/1uLang/zhiannet-api/wazuh/server"
	"github.com/1uLang/zhiannet-api/zstack/model/host_relation"
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

	wazuh_server.InitTable()

	//初始化渠道表
	channels.InitTable()

	//组件表
	subassemblynode.InitTable()

	//ddos ip表
	ddos_host_ip.InitTable()

	//监控
	monitor_list.InitTable()
	monitor_notice.InitTable()

	//数据备份
	nextcloud_model.InitTable()
	agent_model.InitTable()

	//审计系统
	audit_user_relation.InitTable()
	audit_assets_relation.InitTable()
	//资源监控
	resmon_model.InitTable()

	//zstack 云主机
	host_relation.InitTable()
}
