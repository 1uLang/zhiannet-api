package access_gateway

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	nextTerminalAccessGateway struct {
		Id          uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id" form:"id"`         //id
		Auth        uint8  `gorm:"column:is_auth" json:"is_auth" form:"is_auth"`                   //is_auth
		GatewayId   string `gorm:"column:gateway_id" json:"gateway_id" form:"gateway_id"`          //资产ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	nextTerminalAssetGateway struct {
		Id        uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id" form:"id"` //id
		AssetId   string `gorm:"column:asset_id" json:"asset_id" form:"asset_id"`        //资产ID
		GatewayId string `gorm:"column:gateway_id" json:"gateway_id" form:"gateway_id"`  //资产ID
		IsDelete  uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`     //1删除
	}
)

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&nextTerminalAccessGateway{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
	err = db_model.MysqlConn.AutoMigrate(&nextTerminalAssetGateway{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

func (nextTerminalAccessGateway) TableName() string {
	return "next_terminal_access_gateway"
}
func (nextTerminalAssetGateway) TableName() string {
	return "next_terminal_asset_gateway"
}
func addAccessGateway(req *nextTerminalAccessGateway) (err error) {

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	req.IsDelete = 0
	req.CreateTime = time.Now().Unix()
	return db_model.MysqlConn.Create(&req).Error
}
func deleteAccessGateway(id string) error {
	return db_model.MysqlConn.Model(&nextTerminalAccessGateway{}).Where("is_delete = 0").Where("gateway_id = ?", id).Update("is_delete", 1).Error
}
func getList(req *ListReq) (list []*nextTerminalAccessGateway, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&nextTerminalAccessGateway{}).Where("is_delete=?", 0)
	if req != nil {
		fmt.Println("req ...", req)
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
func authAccessGateway(req *AuthorizeReq) (err error) {
	tx := db_model.MysqlConn.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Model(&nextTerminalAccessGateway{}).Where("is_delete=?", 0).
		Where("is_auth = 1").Where("gateway_id = ?", req.Id).Update("is_delete", 1).Error
	if err != nil {
		return err
	}
	for _, v := range req.AdminUserIds {
		obj := nextTerminalAccessGateway{
			GatewayId:   req.Id,
			Auth:        1,
			AdminUserId: v,
			IsDelete:    0,
			CreateTime:  time.Now().Unix(),
		}
		err = tx.Create(obj).Error
		err = tx.Create(obj).Error
		if err != nil {
			return err
		}
	}
	for _, v := range req.UserIds {
		obj := nextTerminalAccessGateway{
			GatewayId:  req.Id,
			Auth:       1,
			UserId:     v,
			IsDelete:   0,
			CreateTime: time.Now().Unix(),
		}
		err = tx.Create(&obj).Error
		if err != nil {
			return err
		}
	}
	return
}
func authUserList(id string) ([]uint64, error) {

	var o []nextTerminalAccessGateway
	var ids []uint64
	err := db_model.MysqlConn.Model(&nextTerminalAccessGateway{}).Where("is_delete=?", 0).Where("gateway_id = ?", id).
		Where("is_auth = 1").Find(&o).Error
	for _, v := range o {
		if v.UserId > 0 {
			ids = append(ids, v.UserId)
		} else if v.AdminUserId > 0 {
			ids = append(ids, v.AdminUserId)
		}
	}
	return ids, err
}

func getUserNum(gateway string) (int64, error) {
	var count int64
	err := db_model.MysqlConn.Model(&nextTerminalAccessGateway{}).Where("is_delete=?", 0).Where("gateway_id = ?", gateway).
		Where("is_auth = 1").Count(&count).Error
	return count, err
}
func SyncAssetGateway(asset, gateway string) error {
	return db_model.MysqlConn.Create(&nextTerminalAssetGateway{AssetId: asset, GatewayId: gateway}).Error
}
func checkAssetGateway(gateway string) (bool, error) {
	var count int64
	err := db_model.MysqlConn.Model(&nextTerminalAssetGateway{}).Where("is_delete = 0").Where("gateway_id = ?", gateway).Count(&count).Error
	return err == nil && count > 0, err
}
func DeleteAsset(asset string) error {
	return db_model.MysqlConn.Model(&nextTerminalAssetGateway{}).Where("is_delete = 0").Where("asset_id = ?", asset).Update("is_delete", 1).Error
}
