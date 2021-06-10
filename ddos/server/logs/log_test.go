package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}

//攻击日志
func Test_attack_log_list(t *testing.T) {
	InitDB()
	list, err := GetAttackLogList(&AttackLogReq{
		NodeId: 1,
		Addr:   "182.150.0.37",
	})
	fmt.Println(list)
	fmt.Println(err)

}

//流量日志
func Test_traffic_log_list(t *testing.T) {

	InitDB()
	list, err := GetTrafficLogList(&TrafficLogReq{
		NodeId: 1,
		Addr:   "182.150.0.37",
		Level:  3,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//链接日志
func Test_link_log_list(t *testing.T) {

	InitDB()
	list, err := GetLinkLogList(&LinkLogReq{
		NodeId: 1,
		Addr:   "182.150.0.37",
		Level:  3,
	})
	fmt.Println(list)
	fmt.Println(err)
}
