package bwlist

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	HIDSBWList struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		IP          string `gorm:"column:ip" json:"ip" form:"ip"`                                  //agent ip
		Black       bool   `gorm:"column:black" json:"black" form:"black"`                         //是否为：黑名单
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	ListReq struct {
		IP 			string
		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
)

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&HIDSBWList{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

func checkBWIP(req *HIDSBWList) (bool, error) {

	model := db_model.MysqlConn.Model(&HIDSBWList{}).Where("is_delete=?", 0).Where("ip=?", req.IP).
		Where("black=?", req.Black)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
	}

	var total int64
	err := model.Count(&total).Error
	return total > 0, err
}

func AddBWList(req *HIDSBWList) (err error) {

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	ok, err := checkBWIP(req)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return db_model.MysqlConn.Create(&req).Error
}
func DeleteBWList(id uint64) error {
	res := db_model.MysqlConn.Model(&HIDSBWList{}).Where("id = ?", id).Update("is_delete", 1)
	return res.Error
}
func GetBWList(req *ListReq) (list []*HIDSBWList, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&HIDSBWList{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if  req.IP != "" {
			model = model.Where("ip=?", req.IP)
		}
	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = model.Debug().Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}
