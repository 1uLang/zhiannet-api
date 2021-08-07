package cert

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	listReq struct {
		UserId      uint64
		AdminUserId uint64
	}
	nextTerminalCert struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		Auth        uint8  `gorm:"column:is_auth" json:"is_auth" form:"is_auth"`                   //is_auth
		CertId      string `gorm:"column:cert_id" json:"cert_id" form:"cert_id"`                   //资产ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    int    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
)

func getList(req *listReq) (list []*nextTerminalCert, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&nextTerminalCert{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
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
func addCert(req *nextTerminalCert) (err error) {

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	model := db_model.MysqlConn.Model(&nextTerminalCert{}).Where("is_delete=?", 0).
		Where("cert_id=?", req.CertId)

	if req.UserId > 0 {
		model = model.Where("user_id=?", req.UserId)
	}
	if req.AdminUserId > 0 {
		model = model.Where("admin_user_id=?", req.AdminUserId)
	}
	var total int64
	err = model.Count(&total).Error
	if err != nil || total > 0 { //存在则不需要添加
		return
	}
	return db_model.MysqlConn.Create(&req).Error
}
func countAuthCert(id string) (total int64,err error) {
	err = db_model.MysqlConn.Model(&nextTerminalCert{}).Where("is_delete=?", 0).
		Where("cert_id=?", id).Where("is_auth=?", 1).Count(&total).Error
	return
}
func deleteCert(id string) error {
	res := db_model.MysqlConn.Model(&nextTerminalCert{}).Where("cert_id = ?", id).Update("is_delete", 1)
	return res.Error
}
func resetAuthorize(id string) error {
	return db_model.MysqlConn.Model(&nextTerminalCert{}).Where("is_delete=?", 0).
		Where("cert_id=?", id).Where("is_auth=?", 1).Update("is_delete", 1).Error
}
func listAuthorize(id string) (list []*nextTerminalCert, err error) {

	err = db_model.MysqlConn.Model(&nextTerminalCert{}).Where("is_delete=?", 0).
		Where("cert_id=?", id).Where("is_auth=?", 1).Find(&list).Error
	return
}
