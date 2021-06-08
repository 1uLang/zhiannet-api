package dashboard

import "github.com/1uLang/zhiannet-api/awvs/model/dashboard"

//AWVS 仪表板接口层

//MeState 仪表板数据
func MeState() (info map[string]interface{}, err error) {
	return dashboard.MeStats()
}
