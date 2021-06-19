package global_status

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/global_status"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	GlobalReq struct {
		NodeId uint64 `json:"node_id"`
	}
	GlobalResp struct {
		CPU struct {
			Cpus  string   `json:"cpus"`  //核心数
			Load  []string `json:"load"`  //cpu 负载
			Model string   `json:"model"` //cpu type
			Used  string   `json:"used"`  //cpu使用率
		} `json:"cpu"`
		Config struct {
			LastChange     string `json:"last_change"`
			LastChangeFrmt string `json:"last_change_frmt"` //最近一次配置时间
		} `json:"config"`
		DateFrmt string `json:"date_frmt"` //当前时间
		DateTime string `json:"date_time"`
		Disk     struct {
			Devices []struct { //磁盘信息
				Available  string `json:"available"`
				Capacity   string `json:"capacity"`
				Device     string `json:"device"`
				Mountpoint string `json:"mountpoint"`
				Size       string `json:"size"`
				Type       string `json:"type"`
				Used       string `json:"used"`
			} `json:"devices"`
			Swap []struct { //swap 信息
				Device string `json:"device"`
				Total  string `json:"total"`
				Used   string `json:"used"`
			} `json:"swap"`
		} `json:"disk"`
		Kernel struct {
			Mbuf struct { //MBUF 信息
				Max   string `json:"max"`
				Total string `json:"total"`
			} `json:"mbuf"`
			Memory struct { //内存信息
				Total string `json:"total"`
				Used  string `json:"used"`
			} `json:"memory"`
			Pf struct { //状态表大小
				Maxstates string `json:"maxstates"`
				States    string `json:"states"`
			} `json:"pf"`
		} `json:"kernel"`
		Uptime   string   `json:"uptime"`   //运行时间 秒
		Versions []string `json:"versions"` //版本信息
	}
)

//全局状态
func GetGlobalStatus(req *GlobalReq) (res *GlobalResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}

	resp, err := global_status.GetGlobal(loginInfo)
	if err != nil {
		return res, err
	}
	res = &GlobalResp{}
	//cpu信息
	res.CPU.Model = resp.Data.System.CPU.Model
	res.CPU.Load = resp.Data.System.CPU.Load
	res.CPU.Cpus = resp.Data.System.CPU.Cpus
	res.CPU.Used = resp.Data.System.CPU.Used

	//最近一次配置保存时间
	res.Config = resp.Data.System.Config
	res.DateFrmt = resp.Data.System.DateFrmt
	res.DateTime = resp.Data.System.DateTime

	//磁盘信息
	res.Disk = resp.Data.System.Disk

	//Kernel
	res.Kernel = resp.Data.System.Kernel
	//运行时间
	res.Uptime = resp.Data.System.Uptime
	//版本信息
	res.Versions = resp.Data.System.Versions

	//
	return res, err
}
