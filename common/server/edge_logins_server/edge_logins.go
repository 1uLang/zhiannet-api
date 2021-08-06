package edge_logins_server

import "github.com/1uLang/zhiannet-api/common/model/edge_logins"

//通过UID获取
func GetListByUid(uid []uint64) (res map[uint64]*edge_logins.EdgeLogins, total int64, err error) {
	return edge_logins.GetListByUid(uid)
}

//获取详情
func GetInfoByUid(uid uint64) (res *edge_logins.EdgeLogins, err error) {
	return edge_logins.GetInfoByUid(uid)
}

func Save(req *edge_logins.EdgeLogins) (row int64, err error) {
	return edge_logins.SaveOpt(req)
}

func UpdateOpt(id uint64, isOn uint8) (row int64, err error) {
	return edge_logins.UpdateOpt(id, isOn)
}

//获取otp状态
func GetOtpByName(name string) (res bool, err error) {
	return edge_logins.GetOtpByName(name)
}
