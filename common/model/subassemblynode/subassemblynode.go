package subassemblynode

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	Subassemblynode struct {
		Id     uint64 `json:"id" gorm:"column:id"`         // 节点id
		Name   string `json:"name" gorm:"column:name"`     // 节点名称
		Addr   string `json:"addr" gorm:"column:addr"`     // 节点地址
		Port   int64  `json:"port" gorm:"column:port"`     // 节点端口
		Type   int    `json:"type" gorm:"column:type"`     // 节点类型：ddos、下一代防火墙、云WAF
		Idc    int    `json:"idc" gorm:"column:idc"`       // 数据中心
		State  int    `json:"state" gorm:"column:state"`   // 启用、禁用
		Status int    `json:"status" gorm:"column:status"` // 删除
		Key    string `json:"key" gorm:"column:key"`       // api key
		Secret string `json:"secret" gorm:"column:secret"` // api secret
	}
	NodeReq struct {
		Type  int    `json:"idc" gorm:"column:idc"`     // 数据中心
		State string `json:"state" gorm:"column:state"` // 启用、禁用

	}
)

//获取节点
func GetList(req *NodeReq) (list []*Subassemblynode, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&Subassemblynode{})
	if req != nil {
		if req.State != "" {
			model = model.Where("state=?", req.State)
		}
		if req.Type > 0 {
			model = model.Where("type=?", req.Type)
		}
	}
	err = model.Debug().Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//检查name 是否存在
func CheckMenuNameUnique(name string, id uint64) bool {
	model := model.MysqlConn.Model(&Subassemblynode{}).Where("name=?", name)
	if id != 0 {
		model = model.Where("id!=?", id)
	}
	var num int64
	model.Count(&num)
	return num == 0
}

// 添加菜单操作
func Add(req *Subassemblynode) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := model.MysqlConn.Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

//修改菜单操作
func Edit(req *Subassemblynode, id uint64) (rows int64, err error) {
	var entity Subassemblynode
	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.Name = req.Name
	entity.Status = req.Status
	entity.Addr = req.Addr
	entity.Port = req.Port
	entity.Type = req.Type
	entity.Idc = req.Idc
	entity.State = req.State
	entity.Key = req.Key
	entity.Secret = req.Secret
	res := model.MysqlConn.Model(&Subassemblynode{}).Where("id=?", id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected
	return
}

//删除菜单
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Where("id in (?)", ids).Delete(&Subassemblynode{})
	return res.Error
}

//获取节点详细信息
func GetNodeInfoById(id uint64) (info Subassemblynode, err error) {
	err = model.MysqlConn.Where("id=?", id).First(&info).Error
	return
}
