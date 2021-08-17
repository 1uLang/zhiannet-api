package audit_assets_relation

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	AuditAssetsRelation struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		AssetsId    uint64 `gorm:"column:assets_id" json:"assets_id" form:"assets_id"`             //资产id
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户id
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户id
		AssetsType  int    `gorm:"column:assets_type" json:"assets_type" form:"assets_type"`       //0 数据库 1主机 2应用 3设备
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	ListReq struct {
		UserId      uint64   `json:"user_id"`
		AdminUserId uint64   `json:"admin_user_id"`
		AssetsId    []uint64 `json:"assets_id"`
		AssetsType  int      `json:"assets_type"`
		PageNum     int      `json:"page_num"`
		PageSize    int      `json:"page_size"`
	}

	AddReq struct {
		UserId      []uint64 `json:"user_id"`
		AdminUserId []uint64 `json:"admin_user_id"`
		AssetsId    uint64   `json:"assets_id"`
		AssetsType  int      `json:"assets_type"`
	}
)

func GetList(req *ListReq) (list []*AuditAssetsRelation, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Debug().Table("audit_assets_relation").Where("assets_type=?", req.AssetsType)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id = ?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id = ?", req.AdminUserId)
		}

		if len(req.AssetsId) > 0 {
			model = model.Where("assets_id in(?)", req.AssetsId)
		}

	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	return
}

func GetListByUid(req []uint64) (resMap map[uint64]*AuditAssetsRelation, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("audit_assets_relation")
	if len(req) == 0 {
		return
	}
	model = model.Where("user_id in(?)", req)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	var list []*AuditAssetsRelation
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	resMap = make(map[uint64]*AuditAssetsRelation)
	for _, v := range list {
		resMap[v.Id] = v
	}
	return
}

func GetListByAdminUid(req []uint64) (resMap map[uint64]*AuditAssetsRelation, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("audit_assets_relation")
	if len(req) == 0 {
		return
	}
	model = model.Where("admin_user_id in(?)", req)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	var list []*AuditAssetsRelation
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	resMap = make(map[uint64]*AuditAssetsRelation)
	for _, v := range list {
		resMap[v.Id] = v
	}
	return
}

// 添加操作
func Add(req *AuditAssetsRelation) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := model.MysqlConn.Debug().Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

//修改操作
func Edit(req *AuditAssetsRelation, id uint64) (rows int64, err error) {
	var entity AuditAssetsRelation

	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.AssetsId = req.AssetsId
	entity.AdminUserId = req.AdminUserId
	entity.UserId = req.UserId
	res := model.MysqlConn.Model(&AuditAssetsRelation{}).Where("id=?", id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected
	return
}

//删除菜单
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Where("id in (?)", ids).Delete(&AuditAssetsRelation{})
	return res.Error
}

//重置
func Reset(req *AddReq) (err error) {
	if req == nil {
		return
	}
	if len(req.UserId) > 0 {
		model.MysqlConn.Where("user_id >? and assets_id=?", 0, req.AssetsId).Delete(&AuditAssetsRelation{})
		for _, v := range req.UserId {
			Add(&AuditAssetsRelation{
				AdminUserId: 0,
				UserId:      v,
				AssetsType:  req.AssetsType,
				CreateTime:  int(time.Now().Unix()),
				AssetsId:    req.AssetsId,
			})
		}
	}
	if len(req.AdminUserId) > 0 {
		model.MysqlConn.Debug().Where("admin_user_id > ? and assets_id=?", 0, req.AssetsId).Delete(&AuditAssetsRelation{})
		for _, v := range req.AdminUserId {
			Add(&AuditAssetsRelation{
				AdminUserId: v,
				UserId:      0,
				AssetsType:  req.AssetsType,
				CreateTime:  int(time.Now().Unix()),
				AssetsId:    req.AssetsId,
			})
		}
	}

	return nil
}

//通过用户名获取用户信息
//func GetInfoByUsername(name string) (info *AuditAssetsRelation, err error) {
//	err = model.MysqlConn.Table("audit_assets_relation").Where("username=?", name).First(&info).Error
//	return
//}
