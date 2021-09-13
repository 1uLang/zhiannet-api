package attack_check_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	black_white_list_server "github.com/1uLang/zhiannet-api/ddos/server/black_white_list"
	"strings"
	"time"
)
import host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
import logs_server "github.com/1uLang/zhiannet-api/ddos/server/logs"

type ddos struct{}

//添加到ddos全局黑名单

func (ddos) AddBlackIP(ips []string) error {

	//ddos 节点
	ds, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		return err
	}

	for _, v := range ds {
		for _, ip := range ips {
			needMessage := true
			list, _ := black_white_list_server.GetBWList(&black_white_list_server.BWReq{Addr: ip, NodeId: v.Id})
			for _, item := range list.Bwlist {

				if item.Address == ip && item.Flags == "blacklist" { //黑名单 - 不告警
					needMessage = false
				}
			}
			if needMessage {

				_, err = black_white_list_server.AddBW(&black_white_list_server.EditBWReq{NodeId: v.Id, Addr: []string{ip}})
				if err != nil {
					return err
				}
				edge_messages.Add(&edge_messages.Edgemessages{
					Level:     "error",
					Subject:   "入侵检测",
					Body:      fmt.Sprintf("主机防护入侵检测[%s],已自动添加至DDoS服务器IP黑名单", ip),
					Type:      "IntrusionDetection",
					Params:    "{}",
					Createdat: uint64(time.Now().Unix()),
					Day:       time.Now().Format("20060102"),
					Hash:      "",
					Role:      "admin",
				})
			}
		}

	}
	return nil
}

func (ddos) AttackCheck() error {

	//ddos 节点
	ds, _, err := host_status_server.GetDdosNodeList()
	if err != nil {
		return err
	}
	start, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	for _, v := range ds {
		attacks, err := logs_server.GetAttackLogList(&logs_server.AttackLogReq{NodeId: v.Id,
			StartTime: start, EndTime: start.AddDate(0, 0, 1)})
		if err != nil {
			return err
		}
		for _, attack := range attacks.Report {
			//去掉无效ip
			attack.FromAddress = strings.ReplaceAll(attack.FromAddress, "0.0.0.0/0", "")

			if attack.FromAddress == "" {
				continue
			}
			ips := strings.Split(attack.FromAddress, " ")
			if len(ips) > 0 {
				//将ip加入改节点黑名单
				for _, ip := range ips {
					_, err = black_white_list_server.AddBW(&black_white_list_server.EditBWReq{NodeId: v.Id, Addr: []string{ip}})
				}
				//将ip加到hids 黑名单中
				_ = addBlackList(ips)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
