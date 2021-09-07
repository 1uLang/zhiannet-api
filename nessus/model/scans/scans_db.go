package scans

import (
	"encoding/json"
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	NessusScans struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		ScansId     uint64 `gorm:"column:scans_id" json:"scans_id" form:"scans_id"`                //资产ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
		Description string `gorm:"column:description" json:"description" form:"description"`       //备注
		Addr        string `gorm:"column:addr" json:"addr" form:"addr"`                            //备注
		Config      []byte `gorm:"column:config" json:"config" form:"config"`                      //登录设置 配置
	}
	ScansListReq struct {
		UserId      uint64 `json:"user_id" gorm:"column:user_id"`                                  // 用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		PageNum     int    `json:"page_num" `                                                      //页数
		PageSize    int    `json:"page_size" `                                                     //每页条数
		ScansId     uint64 `gorm:"column:scans_id" json:"scans_id" form:"scans_id"`                //资产ID
	}
	NessusScanReport struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		ScansId     uint64 `gorm:"column:scans_id" json:"scans_id" form:"scans_id"`                //资产ID
		HistoryId   uint64 `gorm:"column:history_id" json:"history_id" form:"history_id"`          //资产ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
		Addr        string `gorm:"column:addr" json:"addr" form:"addr"`                            //备注
	}
	SetConfigResp struct {
		ID string `json:"id"`
		GetConfigResp
	}
	GetConfigResp struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Os       int    `json:"os"`
		Port     int    `json:"port"`
	}
)

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&NessusScans{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
	err = db_model.MysqlConn.AutoMigrate(&NessusScanReport{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

func GetInfo(id string) (info NessusScans, err error) {
	var entity NessusScans
	err = db_model.MysqlConn.Model(&NessusScans{}).Where("scans_id=?", id).Where("is_delete=0").Find(&entity).Error
	return entity, err
}
func SetConfig(conf *SetConfigResp) (err error) {
	bytes, err := json.Marshal(conf.GetConfigResp)
	if err != nil {
		return err
	}
	return db_model.MysqlConn.Model(&NessusScans{}).Where("scans_id=?", conf.ID).Where("is_delete=0").Update("config", string(bytes)).Error
}
func GetConfig(id string) (info *GetConfigResp, err error) {
	var entity NessusScans
	err = db_model.MysqlConn.Model(&NessusScans{}).Where("scans_id=?", id).Where("is_delete=0").Find(&entity).Error
	if err != nil {
		return nil, err
	}
	config := &GetConfigResp{}
	err = json.Unmarshal(entity.Config, &config)
	return config, err
}

//获取节点
func GetList(req *ScansListReq) (list []*NessusScans, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&NessusScans{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.ScansId != 0 {
			model = model.Where("scans_id=?", req.ScansId)
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
func GetNum(req *ScansListReq) (total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&NessusScans{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.ScansId != 0 {
			model = model.Where("scans_id=?", req.ScansId)
		}
	}
	err = model.Count(&total).Error

	return
}

//添加数据
func AddScans(req *NessusScans) (insertId uint64, err error) {
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

//用scans ID 删除
func DeleteByScansIds(ids []string) (err error) {
	res := db_model.MysqlConn.Model(&NessusScans{}).Where("scans_id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func DeleteByIds(ids []uint64) (err error) {
	res := db_model.MysqlConn.Model(&NessusScans{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func AddScansReport(req *NessusScanReport) (err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := db_model.MysqlConn.Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func DeleteScansReportById(ids []string) (err error) {
	res := db_model.MysqlConn.Model(&NessusScanReport{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func DeleteScansReportByHistoryId(id, historyId string) (err error) {
	res := db_model.MysqlConn.Model(&NessusScanReport{}).Where("scans_id = ?", id).Where("history_id = ?", historyId).Update("is_delete", 1)
	return res.Error
}

//获取节点
func GetListReport(req *ScansListReq) (list []*NessusScanReport, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&NessusScanReport{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.ScansId != 0 {
			model = model.Where("scans_id=?", req.ScansId)
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
