package ddos_host_ip

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	DdosHostIp struct {
		Id         uint64 `json:"id" gorm:"column:id"`                   // id
		Addr       string `json:"addr" gorm:"column:addr"`               // ip地址
		NodeId     uint64 `json:"node_id" gorm:"column:node_id"`         // 节点ID
		CreateTime int64  `json:"create_time" gorm:"column:create_time"` // 创建时间
	}
	HostReq struct {
		Addr     string `json:"addr" `
		NodeId   uint64 `json:"node_id" `
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
	AddHost struct {
		Addr   string `json:"addr" gorm:"column:addr"`       // ip地址
		NodeId uint64 `json:"node_id" gorm:"column:node_id"` // 节点ID
	}
)

//获取节点
func GetList(req *HostReq) (list []*DdosHostIp, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&DdosHostIp{})
	if req != nil {
		if req.Addr != "" {
			model = model.Where("addr=?", req.Addr)
		}
		if req.NodeId > 0 {
			model = model.Where("node_id=?", req.NodeId)
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
	if err != nil {
		return
	}
	return
}

//检查name 是否存在
//func CheckMenuNameUnique(name string, id int) bool {
//	model := model.MysqlConn.Model(&DdosHostIp{}).Where("name=?", name)
//	if id != 0 {
//		model = model.Where("id!=?", id)
//	}
//	var num int64
//	model.Count(&num)
//	return num == 0
//}

// 添加ip操作
func Add(req *AddHost) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	host := DdosHostIp{
		Addr:       req.Addr,
		NodeId:     req.NodeId,
		CreateTime: time.Now().Unix(),
	}
	res := model.MysqlConn.Create(&host)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = host.Id
	return
}

//修改菜单操作
func Edit(req *AddHost, id uint64) (rows int64, err error) {
	var entity DdosHostIp
	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.NodeId = req.NodeId
	entity.Addr = req.Addr
	res := model.MysqlConn.Model(&DdosHostIp{}).Where("id=?", id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected
	return
}

//删除
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Where("id in (?)", ids).Delete(&DdosHostIp{})
	return res.Error
}
