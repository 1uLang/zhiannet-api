package edge_message_media_instances

import (
	"encoding/json"
	"github.com/1uLang/zhiannet-api/common/model"
)

type MessageMediaInstance struct {
	Id          uint32 `field:"id"`          // ID
	Name        string `field:"name"`        // 名称
	IsOn        uint8  `field:"isOn"`        // 是否启用
	MediaType   string `field:"mediaType"`   // 媒介类型
	Params      string `field:"params"`      // 媒介参数
	Description string `field:"description"` // 备注
	State       uint8  `field:"state"`       // 状态
}
type EmailInfo struct {
	Id       uint32 `json:"id"`
	Smtp     string `json:"smtp"`
	Username string `json:"username"`
	Password string `json:"password"`
	From     string `json:"from"`
}

func (MessageMediaInstance) TableName() string {
	return "edgeMessageMediaInstances"
}
func (this MessageMediaInstance) GetEmail() (EmailInfo, error) {
	msg := MessageMediaInstance{}
	err := model.MysqlConn.Table(this.TableName()).Where("state=1").Where("isOn=1").Where("mediaType=?", "email").Find(&msg).Error
	email := EmailInfo{}
	email.Id = msg.Id
	if msg.Params != "" {
		err = json.Unmarshal([]byte(msg.Params), &email)
	}
	return email, err
}
