package host

//主机信息
import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
	"time"
)

type (
	HostListReq struct {
		Uid uint64 `json:"uid"`
	}

	//主机列表响应参数
	HostListResp struct {
		Inventories []Inventories `json:"inventories"`
		Error       Error         `json:"error"`
	}

	//创建主机请求参数
	CreateHostReq struct {
		Params ParamsHost `json:"params"`
	}
	ParamsHost struct {
		Name                 string   `json:"name"`
		InstanceOfferingUuid string   `json:"instanceOfferingUuid"` //计算规格UUID
		ImageUuid            string   `json:"imageUuid"`            //镜像UUID
		L3NetworkUuids       []string `json:"l3NetworkUuids"`       //三层网络UUID列表
		RootDiskOfferingUuid string   `json:"rootDiskOfferingUuid"` //根云盘规格UUID
	}
	CreateHostResp struct {
		Inventory Inventory `json:"inventory"`
		Error     Error     `json:"error"`
	}
	Inventory struct {
		AllVolumes           []AllVolumes `json:"allVolumes"`
		VMNics               []VMNics     `json:"vmNics"`
		AllocatorStrategy    string       `json:"allocatorStrategy"`
		ClusterUUID          string       `json:"clusterUuid"`
		CPUNum               int64        `json:"cpuNum"`
		CreateDate           string       `json:"createDate"`
		DefaultL3NetworkUUID string       `json:"defaultL3NetworkUuid"`
		Description          string       `json:"description"`
		HostUUID             string       `json:"hostUuid"`
		HypervisorType       string       `json:"hypervisorType"`
		ImageUUID            string       `json:"imageUuid"`
		InstanceOfferingUUID string       `json:"instanceOfferingUuid"`
		LastHostUUID         string       `json:"lastHostUuid"`
		LastOpDate           string       `json:"lastOpDate"`
		MemorySize           int64        `json:"memorySize"`
		Name                 string       `json:"name"`
		Platform             string       `json:"platform"`
		RootVolumeUUID       string       `json:"rootVolumeUuid"`
		State                string       `json:"state"`
		Type                 string       `json:"type"`
		UUID                 string       `json:"uuid"`
		ZoneUUID             string       `json:"zoneUuid"`
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
		ProhibitMigrating    bool         `json:"prohibitMigrating"` //禁止迁移
		ManagementIp         string       `json:"managementIp"`      //可用迁移物理机列表  物理机IP
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

	//操作主机请求参数
	ActionReq struct {
		Uuid     string `json:"uuid"`      //云主机ID
		HostUUid string `json:"host_uuid"` //物理机ID
	}
	ActionResp struct {
		Location string `json:"location"`
	}

	//迁移主机响应参数
	MigrationResp struct {
		Inventory Inventory `json:"inventory"`
		Error     Error     `json:"error"`
	}

	DeleteResp struct {
		Error Error `json:"error"`
	}

	//物理机请求参数
	HostsReq struct {
	}
	//物理机响应参数
	HostsResp struct {
		Inventories []Inventories `json:"inventories"`
		Error       Error         `json:"error"`
	}

	//修改云主机计算规格 请求参数
	UpdateSpecReq struct {
		SpecUUid string `json:"spec_uuid"` //规格ID
		HostUUid string `json:"host_uuid"` //云主机ID
	}
	//修改云主机计算规格 响应参数
	UpdateSpecResp struct {
		Inventory Inventory `json:"inventory"`
		Error     Error     `json:"error"`
	}
)

//云主机列表
func HostList(req *HostListReq) (resp *HostListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_HOST_LIST)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//物理机列表
func Hosts(req *HostsReq) (resp *HostsResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_HOSTS)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//创建主机
func CreateHost(req *CreateHostReq) (resp *CreateHostResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_HOST_LIST)
	logReq.QueryParams = req
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}
	actRes := ActionResp{}
	err = json.Unmarshal(res, &actRes)
	if err != nil || actRes.Location == "" {
		return
	}
	//轮询结果
	res, err = GetUrl(actRes.Location)
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

//启动主机
func StartHost(req *ActionReq) (resp *ActionResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "startVmInstance": {},
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

//停止云主机
func StopHost(req *ActionReq) (resp *ActionResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "stopVmInstance": {
			"type": "grace"
		  },
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

//重启主机
func RestartHost(req *ActionReq) (resp *ActionResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "rebootVmInstance": {
			"type": "grace"
		  },
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

//获取可热迁移的物理机列表
func MigrationCandidateHost(req *ActionReq) (resp *HostListResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_MIGRATION_CANDIDATE, req.Uuid))
	logReq.QueryParams = `{
		  "vmInstanceUuid": {},
		  "systemTags": [],
		  "userTags": []
		}`
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)

	return
}

//迁移
func MigrationHost(req *ActionReq) (resp *MigrationResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v",
		logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.Uuid))
	logReq.QueryParams = `{
		  "migrateVm": {
			"hostUuid": "` + req.HostUUid + `",
			"migrateFromDestination": false
		  },
		  "systemTags": [],
		  "userTags": []
		}`
	logReq.ReqType = "put"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	//err = json.Unmarshal(res, &resp)
	actRes := ActionResp{}
	err = json.Unmarshal(res, &actRes)
	if err != nil || actRes.Location == "" {
		return
	}
	//轮询结果
	res, err = GetUrl(actRes.Location)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

//删除主机
func DelHost(req *ActionReq) (resp *DeleteResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, fmt.Sprintf("%v/%v", _const.ZSTACK_HOST_LIST, req.Uuid))

	logReq.ReqType = "delete"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//更改云主机的计算规格
func UpdateSpec(req *UpdateSpecReq) (resp *UpdateSpecResp, err error) {
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, fmt.Sprintf(_const.ZSTACK_SUSPEND, req.HostUUid))
	logReq.QueryParams = fmt.Sprintf(`{
			"changeInstanceOffering": {
			"instanceOfferingUuid": "%v"
			  },
			"systemTags": [],
			"userTags": []
		}`, req.SpecUUid)
	logReq.ReqType = "put"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	//err = json.Unmarshal(res, &resp)
	actRes := ActionResp{}
	err = json.Unmarshal(res, &actRes)
	if err != nil || actRes.Location == "" {
		return
	}
	//轮询结果
	res, err = GetUrl(actRes.Location)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

func GetUrl(url string) (res []byte, err error) {
	resp, err := request.Client.SetDebug(false).R().Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode() == 202 {
		<-time.Tick(time.Millisecond * 200)
		return GetUrl(url)
	}

	return resp.Body(), err
}
