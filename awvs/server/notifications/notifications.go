package notifications

import (
	"github.com/1uLang/zhiannet-api/awvs/model/notifications"
)

//notifications 接口层

//Notifications 通知消息
func Notifications() (info map[string]interface{}, err error) {
	return notifications.Notifications()
}
