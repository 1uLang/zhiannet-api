package edge_logs_server

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_logs"
)

//所有列表
func GetLogList(req *edge_logs.UserLogReq) (list []*edge_logs.UserLogResp, total int64, err error) {
	list, total, err = edge_logs.GetList(req)
	return
}

//计数
func GetLogNum(req *edge_logs.UserLogReq) (total int64, err error) {
	total, err = edge_logs.GetNum(req)
	return
}
