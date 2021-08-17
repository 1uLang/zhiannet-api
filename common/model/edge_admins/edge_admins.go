package edge_admins

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	EdgeAdmins struct {
		Id        uint64 `gorm:"column:id" json:"id" form:"id"`                      //ID
		Ison      uint8  `gorm:"column:isOn" json:"isOn" form:"isOn"`                //是否启用
		Username  string `gorm:"column:username" json:"username" form:"username"`    //用户名
		Password  string `gorm:"column:password" json:"password" form:"password"`    //密码
		Fullname  string `gorm:"column:fullname" json:"fullname" form:"fullname"`    //全名
		Issuper   uint8  `gorm:"column:isSuper" json:"isSuper" form:"isSuper"`       //是否为超级管理员
		CreatedAt uint64 `gorm:"column:createdAt" json:"createdAt" form:"createdAt"` //创建时间
		UpdatedAt uint64 `gorm:"column:updatedAt" json:"updatedAt" form:"updatedAt"` //修改时间
		State     uint8  `gorm:"column:state" json:"state" form:"state"`             //状态
		Modules   string `gorm:"column:modules" json:"modules" form:"modules"`       //允许的模块
		CanLogin  uint8  `gorm:"column:canLogin" json:"canLogin" form:"canLogin"`    //是否可以登录
		Theme     string `gorm:"column:theme" json:"theme" form:"theme"`             //模板设置
		PwdAt     uint64 `gorm:"column:pwdAt" json:"pwdAt" form:"pwdAt"`             //密码修改时间
	}
)

func GetList() (list []*EdgeAdmins, total int64, err error) {
	model := model.MysqlConn.Table("edgeAdmins").Where("isOn=?", 1)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}

	return
}

func GetListByUid(req []uint64) (resMap map[uint64]*EdgeAdmins, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("edgeAdmins").Where("isOn=?", 1)
	if len(req) == 0 {
		return
	}
	model = model.Where("id in(?)", req)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	var list []*EdgeAdmins
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	resMap = make(map[uint64]*EdgeAdmins)
	for _, v := range list {
		resMap[v.Id] = v
	}
	return
}

//通过id获取用户信息
func GetInfoById(id uint64) (info *EdgeAdmins, err error) {
	err = model.MysqlConn.Table("edgeAdmins").Where("id=?", id).First(&info).Error
	return
}

//通过用户名获取用户信息
func GetInfoByUsername(name string) (info *EdgeAdmins, err error) {
	err = model.MysqlConn.Table("edgeAdmins").Where("username=?", name).First(&info).Error
	return
}

//通过用户名获取用户信息
func GetInfoByPwd(name, pwd string) (info *EdgeAdmins, err error) {
	err = model.MysqlConn.Table("edgeAdmins").Where("username=? and password=?", name, pwd).First(&info).Error
	return
}

//更新账号密码修改时间
func UpdatePwdAt(id uint64) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeAdmins").Where("id=?", id).Update("pwdAt", time.Now().Unix())
	row = tx.RowsAffected
	err = tx.Error
	return
}

//更新账号密码
func UpdatePwd(id uint64, pwd string) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeAdmins").Where("id=?", id).Updates(map[string]interface{}{"pwdAt": time.Now().Unix(), "password": pwd})
	row = tx.RowsAffected
	err = tx.Error
	return
}
