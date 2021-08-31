package edge_server

import "github.com/1uLang/zhiannet-api/common/model"

type (
	EdgeServer struct {
		ID          uint64 `gorm:"column:id" json:"id" form:"id"`
		Ison        int64  `gorm:"column:isOn" json:"ison" form:"ison"`                      //是否启用
		State       uint8  `gorm:"column:state" json:"state" form:"state"`                   //是否删除
		ServerNames []byte `gorm:"column:serverNames" json:"serverNames" form:"serverNames"` //域名集合
		HttpsJSON   []byte `gorm:"column:https" json:"https" form:"https"`                   //https配置
	}
)

func GetList() (list []EdgeServer, err error) {
	//从数据库获取
	err = model.MysqlConn.Table("edgeServers").Where("isOn=1").Where("state=1").Where("type='httpProxy'").Order("id desc").Find(&list).Error
	return
}
