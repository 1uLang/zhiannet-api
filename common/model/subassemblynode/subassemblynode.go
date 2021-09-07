package subassemblynode

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	Subassemblynode struct {
		Id        uint64 `json:"id" gorm:"column:id"`                 // 节点id
		Name      string `json:"name" gorm:"column:name"`             // 节点名称
		Addr      string `json:"addr" gorm:"column:addr"`             // 节点地址
		Port      int64  `json:"port" gorm:"column:port"`             // 节点端口
		Type      int    `json:"type" gorm:"column:type"`             // 节点类型：ddos、下一代防火墙、云WAF
		Idc       string `json:"idc" gorm:"column:idc"`               // 数据中心
		State     int    `json:"state" gorm:"column:state"`           // 启用、禁用
		IsDelete  int    `json:"is_delete" gorm:"column:is_delete"`   // 删除
		IsSsl     int    `json:"is_ssl" gorm:"column:is_ssl"`         // 是否使用ssl协议 https访问
		Key       string `json:"key" gorm:"column:key"`               // api key
		Secret    string `json:"secret" gorm:"column:secret"`         // api secret
		ConnState int    `json:"conn_state" gorm:"column:conn_state"` // 组件连接状态
	}
	NodeReq struct {
		Type     int    `json:"idc" gorm:"column:idc"`     // 数据中心
		State    string `json:"state" gorm:"column:state"` // 启用、禁用fd
		PageNum  int    `json:"page_num" `                 //页数
		PageSize int    `json:"page_size" `                //每页条数
	}
)

func updateIdc() {
	//修改字段类型
	{

		err := model.MysqlConn.Exec("alter table subassemblynode modify column idc varchar(50);").Error
		if err != nil {
			fmt.Println("update table subassemblynode column error : ", err)
			return
		}
	}
	list := []Subassemblynode{}
	err := model.MysqlConn.Model(&Subassemblynode{}).Find(&list).Error
	if err != nil {
		fmt.Println("update field idc error : ", err)
	} else {
		for _, v := range list {
			switch v.Idc {
			case "1":
				v.Idc = "成都IDC"
			case "2":
				v.Idc = "杭州IDC"
			case "3":
				v.Idc = "济南IDC"
			default:
				continue
			}
			_ = model.MysqlConn.Model(&Subassemblynode{}).Where("id=?", v.Id).Update("idc", v.Idc).Error
		}
	}

}

//获取节点
func GetList(req *NodeReq) (list []*Subassemblynode, total int64, err error) {
	updateIdc()
	//从数据库获取
	model := model.MysqlConn.Model(&Subassemblynode{}).Where("is_delete=?", 0).Order("state DESC,type ASC")
	if req != nil {
		if req.State != "" {
			model = model.Where("state=?", req.State)
		}
		if req.Type > 0 {
			model = model.Where("type=?", req.Type)
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

//获取数量
//获取节点
func GetNum(req *NodeReq) (total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&Subassemblynode{}).Where("is_delete=?", 0)
	if req != nil {
		if req.State != "" {
			model = model.Where("state=?", req.State)
		}
		if req.Type > 0 {
			model = model.Where("type=?", req.Type)
		}
	}
	err = model.Count(&total).Error

	return
}

//检查name 是否存在
func CheckMenuNameUnique(name string, id uint64) bool {
	model := model.MysqlConn.Model(&Subassemblynode{}).Where("name=?", name).Where("is_delete=", 0)
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
	if !CheckMenuNameUnique(req.Name, 0) {
		err = fmt.Errorf("组件名称已存在")
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

	if !CheckMenuNameUnique(req.Name, id) {
		err = fmt.Errorf("组件名称已存在")
		return
	}

	err = model.MysqlConn.Where("id=?", id).Find(&entity).Error
	if err != nil {
		return
	}
	entity.Name = req.Name
	entity.IsDelete = req.IsDelete
	entity.Addr = req.Addr
	entity.Port = req.Port
	entity.Type = req.Type
	entity.Idc = req.Idc
	entity.State = req.State
	if req.Key != "" {
		entity.Key = req.Key
	}
	if req.Secret != "" {
		entity.Secret = req.Secret
	}
	entity.IsSsl = req.IsSsl
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
	res := model.MysqlConn.Model(&Subassemblynode{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

//获取节点详细信息
func GetNodeInfoById(id uint64) (info *Subassemblynode, err error) {
	err = model.MysqlConn.Where("id=?", id).First(&info).Error
	return
}

//更新组件可用状态
func UpdateConnState(id uint64, conn int) (row int64, err error) {
	tx := model.MysqlConn.Model(&Subassemblynode{}).Where("id=?", id).Update("conn_state", conn)
	return tx.RowsAffected, tx.Error
}
