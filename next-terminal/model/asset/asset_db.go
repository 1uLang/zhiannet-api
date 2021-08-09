package asset

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	nextTerminalAssets struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		Auth        uint8  `gorm:"column:is_auth" json:"is_auth" form:"is_auth"`                   //is_auth
		AssetsId    string `gorm:"column:asset_id" json:"asset_id" form:"asset_id"`                //资产ID
		Name        string `gorm:"column:name" json:"name" form:"name"`                            //资产名称
		Proto       string `gorm:"column:proto" json:"proto" form:"proto"`                         //资产协议
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    int    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
)

func addAsset(req *nextTerminalAssets)(err error  ){

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	return  db_model.MysqlConn.Create(&req).Error
}
func updateAsset(req *nextTerminalAssets)(err error  ){
	if req == nil || req.AssetsId == "" {
		return fmt.Errorf("参数错误")
	}
	ent := nextTerminalAssets{}
	md := db_model.MysqlConn.Model(req).Where("asset_id=?", req.AssetsId)
	if req.UserId != 0 {
		md = md.Where("user_id=?",req.UserId)
	}else if req.AdminUserId != 0 {
		md = md.Where("admin_user_id=?",req.AdminUserId)
	}
	err = md.Find(&ent).Error
	if err != nil {
		return err
	}

	ent.Name = req.Name
	ent.Proto = req.Proto

	return db_model.MysqlConn.Model(ent).Where("id=?", ent.Id).Save(&ent).Error
}
func deleteAsset(assetId string)error {
	res := db_model.MysqlConn.Model(&nextTerminalAssets{}).Where("asset_id = ?", assetId).Update("is_delete", 1)
	return res.Error
}
func resetAuthorize(id string) error {
	return db_model.MysqlConn.Model(&nextTerminalAssets{}).Where("is_delete=?", 0).
		Where("asset_id=?", id).Where("is_auth=?", 1).Update("is_delete", 1).Error
}
func getInfo(assetId string) (ent *nextTerminalAssets,err error) {
	ent = &nextTerminalAssets{}
	md := db_model.MysqlConn.Model(ent).Where("asset_id=?", assetId).Where("is_auth=0")

	err = md.Find(&ent).Error
	return
}

func GetList(req *ListReq) (list []*nextTerminalAssets, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&nextTerminalAssets{}).Where("is_delete=?", 0)
	if req != nil {
		fmt.Println("req ...",req)
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