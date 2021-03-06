package host_relation

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"gorm.io/gorm"
)

type (
	HostRelation struct {
		ID                uint64 `gorm:"column:id" json:"id" form:"id"`
		UUID              string `gorm:"column:uuid" json:"uuid" form:"uuid"`
		AdminId           uint64 `gorm:"column:admin_id" json:"admin_id" form:"admin_id"`
		UserId            uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`
		ProhibitMigrating int    `gorm:"column:prohibit_migrating" json:"prohibit_migrating" form:"prohibit_migrating"`
		CreateTime        uint64 `gorm:"column:create_time" json:"create_time" form:"create_time"` //时间
	}
	ListReq struct {
		UUID    string `json:"uuid"`
		AdminId uint64 `json:"admin_id"`
		UserId  uint64 `json:"user_id"`
	}
)

func InitTable() {
	err := model.MysqlConn.AutoMigrate(&HostRelation{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}
func GetList(req *ListReq) (list []*HostRelation, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("host_relation")
	if req != nil {
		if req.UUID != "" {
			model = model.Where("uuid = ?", req.UUID)
		}
		if req.AdminId > 0 {
			model = model.Where("admin_id = ?", req.AdminId)
		}
		if req.UserId > 0 {
			model = model.Where("user_id = ?", req.UserId)
		}

	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}

	err = model.Limit(999).Order("id desc").Find(&list).Error
	return
}
func Add(req *HostRelation) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	var info *HostRelation
	err = model.MysqlConn.First(&info, "uuid=?", req.UUID).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, err
		}
		info = &HostRelation{}
	}
	info.UUID = req.UUID
	info.AdminId = req.AdminId
	info.CreateTime = req.CreateTime
	if info.ID == 0 {
		res := model.MysqlConn.Create(&info)
		err = res.Error
		insertId = info.ID
		return
	}
	res := model.MysqlConn.Save(&info)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = info.ID
	return
}
func UpdateMigrating(uuid string, Migrating int) (row int64, err error) {
	tx := model.MysqlConn.Table("host_relation").Where("uuid=?", uuid).Updates(map[string]interface{}{"prohibit_migrating": Migrating})
	row = tx.RowsAffected
	err = tx.Error
	return
}
