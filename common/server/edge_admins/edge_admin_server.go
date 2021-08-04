package edge_admins

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_admins"
	"time"
)

//判断用户密码是否过期 有效90天
func CheckPwdInvalid(name string) (res bool, err error) {
	info, err := edge_admins.GetInfoByUsername(name)
	if err != nil || info == nil {
		return
	}
	if int64(info.PwdAt) < time.Now().Add(-time.Second*60*60*24*90).Unix() {
		res = true
	}
	return res, err
}

//更新密码
func UpdatePwd(name, pwd, newPwd string) (res int64, err error) {
	var info *edge_admins.EdgeAdmins
	info, err = edge_admins.GetInfoByPwd(name, pwd)
	if err != nil || info == nil {
		return 0, fmt.Errorf("更新密码失败,账号或密码错误")
	}

	return edge_admins.UpdatePwd(info.Id, pwd)

}

//更新密码时间
func UpdatePwdAt(id uint64) (res int64, err error) {
	return edge_admins.UpdatePwdAt(id)

}

//登录限制检查
func LoginCheck(name string) (res bool, err error) {
	//先检查是否限制登录
	key := "no_login_" + name
	var value int
	value, err = cache.GetInt(key)
	if err != nil {
		return false, err
	}
	if value > 0 {
		return true, err
	}
	return res, err
}

//登录错误时  计数
func LoginErrIncr(name string) (res bool, err error) {
	key := name
	var ex bool
	ex, err = cache.SetNx(key, time.Minute)
	if err != nil {
		return false, err
	}
	if !ex {
		cache.Incr(key, time.Minute)
		value, err := cache.GetInt(key)
		if value > 4 && err == nil {
			//锁定30分钟
			key = "no_login_" + name
			cache.SetNx(key, time.Minute*30)
		}
	}
	return true, err
}
