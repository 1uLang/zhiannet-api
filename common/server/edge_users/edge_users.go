package edge_users_server

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"time"
)

//判断用户密码是否过期 有效90天
func CheckPwdInvalid(name string) (res bool, err error) {
	info, err := edge_users.GetInfoByUsername(name)
	if err != nil || info == nil {
		return
	}
	if int64(info.PwdAt) < time.Now().Add(-time.Second*60*60*24*90).Unix() {
		res = true
	}
	return
}

//更新密码
func UpdatePwd(id uint64, pwd string) (res int64, err error) {
	return edge_users.UpdatePwd(id, pwd)

}

//更新密码时间
func UpdatePwdAt(id uint64) (res int64, err error) {
	return edge_users.UpdatePwdAt(id)

}
