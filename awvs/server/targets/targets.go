package targets

import (
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
)

//AWVS 目标 api接口

//List 目标列表
func List(req *targets.ListReq) (info map[string]interface{}, err error) {
	return targets.List(req)
}

//条数查询所有
func Search(req *targets.ListReq) (info []interface{}, err error) {
	return targets.Search(req)
}

//Add 新建目标
func Add(req *targets.AddReq) (targetId string, err error) {
	return targets.Add(req)
}
func GetConfig(id uint64) (*targets.GetConfigResp, error) {

	return targets.GetConfig(id)
}
func SetConfig(req *targets.SetConfigReq) error {

	return targets.SetConfig(req)
}

//Delete 删除目标
func Delete(target_id string) (err error) {
	return targets.Delete(target_id)
}

//Update 修改目标
func Update(target_id string, req *targets.UpdateReq) (err error) {
	return targets.Update(target_id, req)
}

//SetLogin 目标列表
func SetLogin(target_id string, req *targets.SetLoginReq) (err error) {
	return targets.SetLogin(target_id, req)
}

//获取扫描目标数量
func GetTargetsNum(req *targets.AddrListReq) (total int64, err error) {

	return targets.GetNum(req)
}
