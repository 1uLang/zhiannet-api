package assets

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	Asset struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		AssetsId    string `gorm:"column:asset_id" json:"asset_id" form:"asset_id"`                //资产ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    int    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	AssetsListReq struct {
		UserId      uint64 `json:"user_id" gorm:"column:user_id"`                                  // 用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		PageNum     int    `json:"page_num" `                                                      //页数
		PageSize    int    `json:"page_size" `                                                     //每页条数
		AssetId     string `gorm:"column:asset_id" json:"asset_id" form:"asset_id"`                //资产ID
	}
)

//获取节点
func GetList(req *AssetsListReq) (list []*Asset, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&Asset{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.AssetId != "" {
			model = model.Where("asset_id=?", req.AssetId)
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
	err = model.Debug().Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//获取数量
func GetNum(req *AssetsListReq) (total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&Asset{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.AssetId != "" {
			model = model.Where("asset_id=?", req.AssetId)
		}
	}
	err = model.Count(&total).Error

	return
}

//添加数据
func AddAsset(req *Asset) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := db_model.MysqlConn.Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

//用asset ID 删除
func DeleteByAssetIds(ids []string) (err error) {
	res := db_model.MysqlConn.Model(&Asset{}).Where("asset_id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func DeleteByIds(ids []uint64) (err error) {
	res := db_model.MysqlConn.Model(&Asset{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}
