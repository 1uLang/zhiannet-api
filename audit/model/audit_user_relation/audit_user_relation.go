package audit_user_relation

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	AuditUserRelation struct {
		Id          uint64 `json:"id" gorm:"column:id"`                       // id
		AdminUserId uint64 `json:"admin_user_id" gorm:"column:admin_user_id"` // waf 管理端用户
		AuditUserid uint64 `json:"audit_user_id" gorm:"column:audit_user_id"` // 审计系统用户ID
		UserId      uint64 `json:"user_id" gorm:"column:user_id"`             // waf用户端 用户ID
		IsDelete    int    `json:"is_delete" gorm:"column:is_delete"`         // 1删除 0未删除
		CreateTime  int64  `json:"create_time" gorm:"column:create_time"`     // 创建时间
	}
	AuditReq struct {
		AdminUserId uint64 `json:"admin_user_id" ` //waf 管理端用户ID
		UserId      uint64 `json:"user_id" `       //waf用户端 用户ID
		AuditUserId uint64 `json:"audit_user_id" ` //审计系统 用户ID
	}
)

//获取关联信息
func GetInfo(req *AuditReq) (info *AuditUserRelation, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&AuditUserRelation{}).Where("is_delete=?", 0)
	if req != nil {
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.AuditUserId > 0 {
			model = model.Where("audit_user_id=?", req.AuditUserId)
		}
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
	}
	err = model.First(&info).Error

	return info, err
}

// 添加
func Add(req *AuditReq) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	host := AuditUserRelation{
		AdminUserId: req.AdminUserId,
		AuditUserid: req.AuditUserId,
		UserId:      req.UserId,
		CreateTime:  time.Now().Unix(),
	}
	res := model.MysqlConn.Create(&host)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = host.Id
	return
}

//修改操作
func Edit(req *AuditReq, id uint64) (rows int64, err error) {
	var entity AuditUserRelation
	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.AuditUserid = req.AuditUserId
	entity.AdminUserId = req.AdminUserId
	entity.UserId = req.UserId
	res := model.MysqlConn.Model(&AuditUserRelation{}).Where("id=?", id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected
	return
}

//删除
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&AuditUserRelation{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}
