package edge_message_recipients

import (
	"github.com/1uLang/zhiannet-api/common/model"
)

// 消息媒介接收人
type MessageRecipient struct {
	Id          uint32 `field:"id"`          // ID
	AdminId     uint32 `field:"adminId"`     // 管理员ID
	IsOn        uint8  `field:"isOn"`        // 是否启用
	InstanceId  uint32 `field:"instanceId"`  // 实例ID
	User        string `field:"user"`        // 接收人信息
	GroupIds    string `field:"groupIds"`    // 分组ID
	State       uint8  `field:"state"`       // 状态
	Description string `field:"description"` // 备注
}

func (MessageRecipient) TableName() string {
	return "edgeMessageRecipients"
}
func (this MessageRecipient) GetEmail() (string, error) {
	msg := MessageRecipient{}
	err := model.MysqlConn.Table(this.TableName()).Where("state=1").Where("isOn=1").
		Where("instanceId=?", this.InstanceId).
		Where("adminId=?", this.AdminId).Find(&msg).Error
	return this.User, err
}
