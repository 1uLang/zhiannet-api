package server

import (
	"fmt"

	"github.com/1uLang/zhiannet-api/common/model"
	param "github.com/1uLang/zhiannet-api/resmon/const"
	rsm "github.com/1uLang/zhiannet-api/resmon/model"
)

type Subassemblynode struct {
	ID        int64  `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Addr      string `gorm:"column:addr"`
	State     uint8  `gorm:"column:state"`     // 1启用、0禁用
	IsDelete  uint8  `gorm:"column:is_delete"` // 1删除
	IsSSL     uint8  `gorm:"is_ssl"`           // 1是 0不是
	Key       string `gorm:"column:key"`
	ConnState int    `gorm:"column:conn_state"`
}

// GetNodeInfo 获取节点信息，并更新缓存
func GetNodeInfo() (sn Subassemblynode, err error) {
	sn = Subassemblynode{}
	err = model.MysqlConn.Model(&Subassemblynode{}).Where("type = 9 AND state = 1 AND is_delete = 0").First(&sn).Error
	if sn.ID > 0 {
		param.TEA_KEY = sn.Key
		// param.BASE_URL = sn.Addr
		if sn.IsSSL == 1 {
			param.BASE_URL = fmt.Sprintf(`https://%s`, sn.Addr)
		} else {
			param.BASE_URL = fmt.Sprintf(`http://%s`, sn.Addr)
		}
	}
	return sn, err
}

func GetNodeURL(id string) *rsm.DownInfo {
	// 由于其操作肯定在列表之后，所以这里可以直接获取缓存
	di := rsm.DownInfo{
		Host:    param.BASE_URL,
		DownUrl: fmt.Sprintf("%s/%s", param.BASE_URL, rsm.GetCPUType(GetResmon(id))),
	}

	return &di
}

// AddResmon 增加Resmon记录
func AddResmon(id string, agentID uint8) error {
	rm := rsm.ResMonModel{
		ID:     id,
		OSType: agentID,
	}

	err := model.MysqlConn.Save(&rm).Error
	if err != nil {
		return fmt.Errorf("添加或修改记录失败：%w", err)
	}

	return nil
}

// DeleteResmon 删除Resmon记录
func DeleteResmon(id string) error {
	rm := rsm.ResMonModel{}
	model.MysqlConn.Model(rsm.ResMonModel{}).Where("id = ?", id).First(&rm)
	if rm.OSType == 0 {
		return nil
	}

	err := model.MysqlConn.Delete(&rsm.ResMonModel{}, id).Error
	if err != nil {
		return fmt.Errorf("删除记录失败：%w", err)
	}

	return nil
}

func GetResmon(id string) uint8 {
	rm := rsm.ResMonModel{}
	model.MysqlConn.Model(rsm.ResMonModel{}).Where("id = ?", id).First(&rm)
	if rm.OSType == 0 {
		return 1
	}
	return rm.OSType
}
