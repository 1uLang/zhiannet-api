package global_status

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/global_status"
)

//全局状态
func GetGlobalStatus(req *request.ApiKey) {
	global_status.GetStatusGlobal(req, true)
}

//NAT
func GetNATList(req *request.ApiKey) {
	global_status.GetNATList(req, true)
}
