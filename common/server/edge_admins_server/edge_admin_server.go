package edge_admins_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_admins"
	"gorm.io/gorm"
	"time"
)

func GetList() (list []*edge_admins.EdgeAdmins, total int64, err error) {
	return edge_admins.GetList()
}

//判断用户密码是否过期 有效90天
func CheckPwdInvalid(id uint64) (res bool, err error) {
	info, err := edge_admins.GetInfoById(id)
	if err != nil || info == nil {
		return
	}
	if int64(info.PwdAt) < time.Now().Add(-time.Second*60*60*24*90).Unix() {
		res = true
	}
	fmt.Println("----", info.PwdAt, time.Now().Add(-time.Second*60*60*24*90).Unix())
	return res, err
}

//更新密码
func UpdatePwd(id uint64, newPwd string) (res int64, err error) {
	var info *edge_admins.EdgeAdmins
	info, err = edge_admins.GetInfoById(id)
	if err != nil || info == nil {
		return 0, fmt.Errorf("更新密码失败,获取账号信息失败")
	}

	return edge_admins.UpdatePwd(info.Id, newPwd)

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
	ex, err = cache.SetNx(key, time.Minute*30)
	if err != nil {
		return false, err
	}
	if !ex {
		cache.Incr(key, time.Minute*30)
		value, err := cache.GetInt(key)
		if value > 4 && err == nil {
			//锁定30分钟
			key = "no_login_" + name
			cache.SetNx(key, time.Minute*30)
		}
	}
	return true, err
}

func GetUserInfoByName(name string) (info *edge_admins.EdgeAdmins, err error) {
	info, err = edge_admins.GetInfoByUsername(name)
	if err == gorm.ErrRecordNotFound { //可能找不到数据
		err = nil
	}
	return
}
