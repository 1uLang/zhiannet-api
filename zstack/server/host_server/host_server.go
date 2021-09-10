package host_server

import (
	"github.com/1uLang/zhiannet-api/zstack/model/host_relation"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"time"
)

//all主机列表
func AllHostList(req *host.HostListReq) (resp *host.HostListResp, err error) {
	resp = &host.HostListResp{}
	var resList *host.HostListResp
	resList, err = host.HostList(req)
	if err != nil || resp == nil {
		return
	}
	if len(resList.Inventories) == 0 {
		return
	}

	hosts, err := HostList(&host.HostListReq{})
	if err != nil || hosts == nil {
		return
	}
	hostMap := map[string]string{}
	for _, v := range hosts.Inventories {
		hostMap[v.UUID] = v.ManagementIp
	}

	for _, v := range resList.Inventories {
		if v.State == "Destroyed" {
			continue
		}
		v.ManagementIp = "127.0.0.1"
		if hostIp, ok := hostMap[v.HostUUID]; ok {
			v.ManagementIp = hostIp
		}
		resp.Inventories = append(resp.Inventories, v)
	}
	return
}

//主机列表
func HostList(req *host.HostListReq) (resp *host.HostListResp, err error) {
	resp = &host.HostListResp{}
	var resList *host.HostListResp
	resList, err = host.HostList(req)
	if err != nil || resp == nil {
		return
	}
	if len(resList.Inventories) == 0 {
		return
	}
	//获取用户的
	list, _, err := host_relation.GetList(&host_relation.ListReq{
		//AdminId: req.Uid,
	})
	if err != nil {
		return
	}
	if len(list) == 0 {
		return
	}
	//物理主机信息
	hosts, err := GetHosts(&host.HostsReq{})
	if err != nil || hosts == nil {
		return
	}
	hostMap := map[string]string{}
	for _, v := range hosts.Inventories {
		hostMap[v.UUID] = v.ManagementIp
	}

	uuidMap := map[string]bool{}
	for _, v := range list {
		uuidMap[v.UUID] = v.ProhibitMigrating == 1
	}
	for _, v := range resList.Inventories {
		if v.State == "Destroyed" {
			continue
		}
		if prohibitMigrating, ok := uuidMap[v.UUID]; ok {
			v.ProhibitMigrating = prohibitMigrating
			v.ManagementIp = "127.0.0.1"
			if hostIp, ok := hostMap[v.HostUUID]; ok {
				v.ManagementIp = hostIp
			}
			resp.Inventories = append(resp.Inventories, v)
		}
	}

	return
}

//物理机列表
func GetHosts(req *host.HostsReq) (res *host.HostsResp, err error) {
	return host.Hosts(req)
}

//新增主机
func CreateHost(req *host.CreateHostReq, uid uint64) (res *host.CreateHostResp, err error) {
	res, err = host.CreateHost(req)
	if err != nil || res == nil {
		return
	}
	if res.Inventory.UUID != "" {
		//主机关联用户
		host_relation.Add(&host_relation.HostRelation{
			UUID: res.Inventory.UUID,
			//AdminId: uid,
			CreateTime: uint64(time.Now().Unix()),
		})
	}
	return
}

//启动主机
func StartHost(req *host.ActionReq) (res *host.ActionResp, err error) {
	return host.StartHost(req)
}

//停止主机
func StopHost(req *host.ActionReq) (res *host.ActionResp, err error) {
	return host.StopHost(req)
}

//暂停主机
func Suspend(req *host.SuspendReq) (res *host.SuspendResp, err error) {
	return host.Suspend(req)
}

//恢复主机
func UnSuspend(req *host.SuspendReq) (res *host.SuspendResp, err error) {
	return host.UnSuspend(req)
}

//重启
func RestartHost(req *host.ActionReq) (res *host.ActionResp, err error) {
	return host.RestartHost(req)
}

//删除
func DeleteHost(req *host.ActionReq) (res *host.DeleteResp, err error) {
	return host.DelHost(req)
}

//可迁移的物理主机
func MigrationCandidateHost(req *host.ActionReq) (res *host.HostListResp, err error) {
	return host.MigrationCandidateHost(req)
}

//迁移
func MigrationHost(req *host.ActionReq) (res *host.MigrationResp, err error) {
	return host.MigrationHost(req)
}

//云盘列表
func DiskList(req *host.DiskListReq) (res *host.DiskListResp, err error) {
	return host.DiskList(req)
}

//规格
func SpecList(req *host.SpecListReq) (res *host.SpecListResp, err error) {
	return host.SpecList(req)
}

//镜像
func ImageList(req *host.ImageListReq) (res *host.ImageListResp, err error) {
	return host.ImageList(req)
}

//3层网络
func NetworkList(req *host.NetworkListReq) (res *host.NetworkListResp, err error) {
	return host.NetworkList(req)
}
