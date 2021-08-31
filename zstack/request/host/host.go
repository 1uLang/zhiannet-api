package host

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
	"github.com/iwind/TeaGo/logs"
)

type (
	HostListReq struct {
	}

	//主机列表响应参数
	HostListResp struct {
		Inventories []Inventories `json:"inventories"`
		Error       Error         `json:"error"`
	}

	Inventories struct {
		UUID                 string       `json:"uuid"`
		Name                 string       `json:"name"`
		Description          string       `json:"description"`
		ZoneUUID             string       `json:"zoneUuid"`
		ClusterUUID          string       `json:"clusterUuid"`
		ImageUUID            string       `json:"imageUuid"`
		HostUUID             string       `json:"hostUuid"`
		LastHostUUID         string       `json:"lastHostUuid"`
		InstanceOfferingUUID string       `json:"instanceOfferingUuid"`
		RootVolumeUUID       string       `json:"rootVolumeUuid"`
		Platform             string       `json:"platform"`
		DefaultL3NetworkUUID string       `json:"defaultL3NetworkUuid"`
		Type                 string       `json:"type"`
		HypervisorType       string       `json:"hypervisorType"`
		MemorySize           int64        `json:"memorySize"`
		CPUNum               int          `json:"cpuNum"`
		CPUSpeed             int          `json:"cpuSpeed"`
		AllocatorStrategy    string       `json:"allocatorStrategy"`
		CreateDate           string       `json:"createDate"`
		LastOpDate           string       `json:"lastOpDate"`
		State                string       `json:"state"`
		VMNics               []VMNics     `json:"vmNics"`
		AllVolumes           []AllVolumes `json:"allVolumes"`
		VMCdRoms             []VMCdRoms   `json:"vmCdRoms"`
	}
	VMNics struct {
		UUID           string    `json:"uuid"`
		VMInstanceUUID string    `json:"vmInstanceUuid"`
		UsedIPUUID     string    `json:"usedIpUuid"`
		L3NetworkUUID  string    `json:"l3NetworkUuid"`
		IP             string    `json:"ip"`
		Mac            string    `json:"mac"`
		HypervisorType string    `json:"hypervisorType"`
		Netmask        string    `json:"netmask"`
		Gateway        string    `json:"gateway"`
		DriverType     string    `json:"driverType"`
		UsedIps        []UsedIps `json:"usedIps"`
		InternalName   string    `json:"internalName"`
		DeviceID       int       `json:"deviceId"`
		Type           string    `json:"type"`
		CreateDate     string    `json:"createDate"`
		LastOpDate     string    `json:"lastOpDate"`
	}

	UsedIps struct {
		UUID          string `json:"uuid"`
		IPRangeUUID   string `json:"ipRangeUuid"`
		L3NetworkUUID string `json:"l3NetworkUuid"`
		IPVersion     int    `json:"ipVersion"`
		IP            string `json:"ip"`
		Netmask       string `json:"netmask"`
		Gateway       string `json:"gateway"`
		IPInLong      int64  `json:"ipInLong"`
		VMNicUUID     string `json:"vmNicUuid"`
		CreateDate    string `json:"createDate"`
		LastOpDate    string `json:"lastOpDate"`
	}
	AllVolumes struct {
		UUID               string `json:"uuid"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		PrimaryStorageUUID string `json:"primaryStorageUuid"`
		VMInstanceUUID     string `json:"vmInstanceUuid"`
		RootImageUUID      string `json:"rootImageUuid"`
		InstallPath        string `json:"installPath"`
		Type               string `json:"type"`
		Format             string `json:"format"`
		Size               int64  `json:"size"`
		ActualSize         int    `json:"actualSize"`
		DeviceID           int    `json:"deviceId"`
		State              string `json:"state"`
		Status             string `json:"status"`
		CreateDate         string `json:"createDate"`
		LastOpDate         string `json:"lastOpDate"`
		IsShareable        bool   `json:"isShareable"`
	}
	VMCdRoms struct {
		UUID           string `json:"uuid"`
		VMInstanceUUID string `json:"vmInstanceUuid"`
		DeviceID       int    `json:"deviceId"`
		IsoUUID        string `json:"isoUuid,omitempty"`
		IsoInstallPath string `json:"isoInstallPath,omitempty"`
		Name           string `json:"name"`
		CreateDate     string `json:"createDate"`
		LastOpDate     string `json:"lastOpDate"`
	}

	//错误信息
	Error struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		Details     string `json:"details"`
	}
	//暂停 恢复暂停 请求参数
	SuspendReq struct {
		Uuid string `json:"uuid"` //主机uuID
	}

	// 暂停 恢复暂停 响应参数
	SuspendResp struct {
		Location string `json:"location"`
	}

	//全局参数修改 请求参数
	GlobalParamsReq struct {
		Category string `json:"category"`
		Name     string `json:"name"`
		//ResourceUuid string `json:"ResourceUuid"`
		Value string `json:"value"`
	}
)

//主机列表
func HostList(req *HostListReq) (resp *HostListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_HOST_LIST)
	//logReq.QueryParams = map[string]string{
	//	//"uid":      fmt.Sprintf("%v", req.Uid),
	//	"name":     fmt.Sprintf("%v", req.Name),
	//	"ip":       fmt.Sprintf("%v", req.IP),
	//	"appType":  fmt.Sprintf("%v", req.AppType),
	//	"status":   fmt.Sprintf("%v", req.Status),
	//	"timelong": fmt.Sprintf("%v", req.TimeLong),
	//}
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//暂停主机
func Suspend(req *SuspendReq) (resp *SuspendResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logs.Println("session uuid ", logReq.UUID)
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "pauseVmInstance": {},
		  "systemTags": [],
		  "userTags": []
		}`
	logReq.ReqType = "put"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//恢复暂停主机
func UnSuspend(req *SuspendReq) (resp *SuspendResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logs.Println("session uuid ", logReq.UUID)
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "resumeVmInstance": {},
		  "systemTags": [],
		  "userTags": []
		}`
	logReq.ReqType = "put"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//全局设置
func UpdateGlobalValue(req *GlobalParamsReq) (resp *SuspendResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logs.Println("session uuid ", logReq.UUID)
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_GLOBAL, req.Category, req.Name))
	logReq.QueryParams = `{
		  "updateGlobalConfig": {"value":"` + req.Value + `"},
		  "systemTags": [],
		  "userTags": []
		}`
	//fmt.Println(logReq.Addr)
	logReq.ReqType = "put"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
