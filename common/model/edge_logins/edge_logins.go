package edge_logins

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"gorm.io/gorm"
)

type EdgeLogins struct {
	Id      uint64 `gorm:"column:id" json:"id" form:"id"`                //ID
	AdminId uint64 `gorm:"column:adminId" json:"adminId" form:"adminId"` //管理员ID
	UserId  uint64 `gorm:"column:userId" json:"userId" form:"userId"`    //用户ID
	IsOn    uint8  `gorm:"column:isOn" json:"isOn" form:"isOn"`          //是否启用
	Type    string `gorm:"column:type" json:"type" form:"type"`          //认证方式
	Params  string `gorm:"column:params" json:"params" form:"params"`    //参数
	State   uint8  `gorm:"column:state" json:"state" form:"state"`       //状态
}

func GetListByUid(req []uint64) (resMap map[uint64]*EdgeLogins, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("edgeLogins")
	if len(req) == 0 {
		return
	}
	model = model.Where("userId in(?)", req)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	var list []*EdgeLogins
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	resMap = make(map[uint64]*EdgeLogins)
	for _, v := range list {
		resMap[v.UserId] = v
	}
	return
}

//通过id获取用户信息
func GetInfoById(id uint64) (info *EdgeLogins, err error) {
	err = model.MysqlConn.Table("edgeLogins").Where("id=?", id).First(&info).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

//通过用户id获取信息
func GetInfoByUid(uid uint64) (info *EdgeLogins, err error) {
	err = model.MysqlConn.Table("edgeLogins").Where("userId=?", uid).First(&info).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

//创建
func SaveOpt(req *EdgeLogins) (row int64, err error) {
	var tx *gorm.DB
	if req.Id > 0 {
		//更新
		tx = model.MysqlConn.Table("edgeLogins").Save(&req)
		row = tx.RowsAffected
	} else {
		tx = model.MysqlConn.Table("edgeLogins").Create(&req)
		row = int64(req.Id)
	}

	err = tx.Error
	return
}

//更新状态
func UpdateOpt(id uint64, isOn uint8) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeLogins").Where("id=?", id).Updates(map[string]interface{}{"isOn": isOn})
	row = tx.RowsAffected
	err = tx.Error
	return
}

//先获取用户信息 在用用户信息获取otp状态
func GetOtpByName(name string) (res bool, err error) {
	info, err := edge_users.GetInfoByUsername(name)
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if err != nil || info == nil {
		return
	}
	var otpInfo *EdgeLogins
	err = model.MysqlConn.Debug().Table("edgeLogins").Where("userId=?", info.ID).First(&otpInfo).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if err != nil || otpInfo == nil {
		return
	}
	if otpInfo.IsOn == 1 && otpInfo.Type == "otp" {
		res = true
	}
	return
}
