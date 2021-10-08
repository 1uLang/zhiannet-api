package channels

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	Channels struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                               //渠道id
		Name        string `gorm:"column:name" json:"name" form:"name"`                         //名称
		User        string `gorm:"column:user" json:"user" form:"user"`                         //联系人
		Mobile      string `gorm:"column:mobile" json:"mobile" form:"mobile"`                   //联系电话
		ProductName string `gorm:"column:product_name" json:"product_name" form:"product_name"` //产品名称
		Status      int    `gorm:"column:status" json:"status" form:"status"`                   //状态 1启用、0禁用
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`    //创建时间
		Domain      string `gorm:"column:domain" json:"domain" form:"domain"`                   //渠道域名
		Logo        string `gorm:"column:logo" json:"logo" form:"logo"`                         //产品logo
		Remake      string `gorm:"column:remake" json:"remake" form:"remake"`                   //备注
	}
	ChannelReq struct {
		Name     string   `json:"name"`       // name
		Status   string   `json:"status" `    // status
		PageNum  int      `json:"page_num" `  //页数
		PageSize int      `json:"page_size" ` //每页条数
		Ids      []uint64 `json:"ids"`
	}
)

//初始化建表
func InitTable() {
	err := model.MysqlConn.AutoMigrate(&Channels{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

//获取列表
func GetList(req *ChannelReq) (list []*Channels, total int64, err error) {
	InitTable()
	//从数据库获取
	model := model.MysqlConn.Model(&Channels{}).Order("status DESC,id ASC")
	if req != nil {
		if req.Status != "" {
			model = model.Where("status=?", req.Status)
		}
		if req.Name != "" {
			model = model.Where("like=?", "%"+req.Name+"%")
		}

		if len(req.Ids) > 0 {
			model = model.Where("id in(?)", req.Ids)
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
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&list).Error
	if err != nil {
		return
	}
	return
}

// 添加
func Add(req *Channels) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	req.CreateTime = int(time.Now().Unix())
	res := model.MysqlConn.Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

//修改
func Edit(req *Channels, id uint64) (rows uint64, err error) {
	var entity Channels

	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.Name = req.Name
	entity.User = req.User
	entity.Mobile = req.Mobile
	entity.ProductName = req.ProductName
	entity.Status = req.Status
	entity.Domain = req.Domain
	entity.Remake = req.Remake
	if req.Logo != "" {
		entity.Logo = req.Logo
	}
	res := model.MysqlConn.Model(&Channels{}).Where("id=?", id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = uint64(res.RowsAffected)
	return
}

//删除
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&Channels{}).Delete("id in (?)", ids)
	return res.Error
}

//获取详细信息
func GetChannelById(id uint64) (info *Channels, err error) {
	err = model.MysqlConn.Where("id=?", id).First(&info).Error
	return
}

//更新可用状态
func UpdateState(id uint64, status int) (row int64, err error) {
	tx := model.MysqlConn.Model(&Channels{}).Where("id=?", id).Update("status", status)
	return tx.RowsAffected, tx.Error
}
