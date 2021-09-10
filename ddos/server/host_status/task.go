package host_status

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/zstack/model/host_relation"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/tidwall/gjson"
)

type CheckFlow struct{}
type HostData struct {
	Ip        string
	Suspend   bool
	Migration bool
}

func (c *CheckFlow) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ddos-CheckFlow-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:    1,
		PageNum: 1, PageSize: 99,
	})

	if err != nil {
		return
	}
	host := make([]HostData, 0)
	for _, v := range nodes {
		var page = 1
		for {
			list, _, err := GetHostList(&ddos_host_ip.HostReq{
				NodeId:   v.Id,
				PageSize: 99,
				PageNum:  page,
			})
			page++
			if err != nil {
				continue
			}

			if len(list) == 0 {
				break
			}
			for _, vv := range list {
				var suspend, migration bool
				var ip = vv.Addr
				if vv.BandwidthIn > 100 { //in方向流量大于100，修改全局并发迁移数为0
					migration = true
				}
				if vv.BandwidthOut > 100 { //out方向流量大于100，暂停主机电源
					suspend = true
				}
				host = append(host, HostData{
					Ip:        ip,
					Suspend:   suspend,
					Migration: migration,
				})

			}

		}

	}

	c.HostHandler(host)
}

func (c *CheckFlow) HostHandler(req []HostData) {
	if len(req) > 0 {
		list, err := host.HostList(&host.HostListReq{})
		if err != nil {
			return
		}
		if len(list.Inventories) > 0 {
			be, _ := json.Marshal(list.Inventories)
			dom := gjson.ParseBytes(be)
			for _, hostInfo := range dom.Array() {
				for _, v := range req {
					if hostInfo.Get("vmNics.0.ip").String() == v.Ip {
						if v.Suspend { //暂停电源
							host.Suspend(&host.SuspendReq{
								Uuid: hostInfo.Get("uuid").String(),
							})
						}

						if v.Migration { //设置禁止迁移
							//host.UpdateGlobalValue(&host.GlobalParamsReq{
							//	Category: "kvm",
							//	Value:    "0",
							//	Name:     "vm.migrationQuantity",
							//})
							host_relation.UpdateMigrating(hostInfo.Get("uuid").String(), 0)
						}
					}
				}

			}
		}

	}
}
