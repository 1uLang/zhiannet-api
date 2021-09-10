package host_server

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"github.com/1uLang/zhiannet-api/zstack/model/host_relation"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/tidwall/gjson"
	"time"
)

type CheckHost struct{}

func (check *CheckHost) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("zstack-CheckHost-----------------------------------------------", err)
		}
	}()
	list, err := AllHostList(&host.HostListReq{})
	if err != nil {
		return
	}

	if len(list.Inventories) > 0 {
		for _, v := range list.Inventories {
			be, _ := json.Marshal(v)
			dom := gjson.ParseBytes(be)
			if ip := dom.Get("vmNics.0.ip").String(); ip != "" {
				AddHostIp(ip)

				//主机关联用户
				host_relation.Add(&host_relation.HostRelation{
					UUID: dom.Get("uuid").String(),
					//AdminId: uid,
					CreateTime: uint64(time.Now().Unix()),
				})
			}

		}
	}
}

func AddHostIp(ip string) (err error) {
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     1, //ddos
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		return
	}

	for _, v := range nodes {
		//ddos_host_ip.Add(&ddos_host_ip.AddHost{
		//	Addr:ip,
		//	NodeId:v.Id,
		//})
		host_status.AddAddr(&ddos_host_ip.AddHost{
			Addr:   ip,
			NodeId: v.Id,
		})
		//添加到数据表

	}
	return
}
