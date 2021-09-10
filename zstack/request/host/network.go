package host

//3层网络信息
import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
)

type (
	NetworkListReq struct {
		UUid string `json:"uuid"`
	}

	NetworkListResp struct {
		Inventories []NetworkInventories `json:"inventories"`
	}
	NetworkInventories struct {
		Category        string           `json:"category"`
		CreateDate      string           `json:"createDate"`
		DNS             []string         `json:"dns"`
		HostRoute       []HostRoute      `json:"hostRoute"`
		IPRanges        []IPRanges       `json:"ipRanges"`
		IPVersion       int64            `json:"ipVersion"`
		L2NetworkUUID   string           `json:"l2NetworkUuid"`
		LastOpDate      string           `json:"lastOpDate"`
		Name            string           `json:"name"`
		NetworkServices []NetworkService `json:"networkServices"`
		State           string           `json:"state"`
		System          bool             `json:"system"`
		Type            string           `json:"type"`
		UUID            string           `json:"uuid"`
		ZoneUUID        string           `json:"zoneUuid"`
	}
	HostRoute struct {
		CreateDate    string `json:"createDate"`
		ID            int64  `json:"id"`
		L3NetworkUUID string `json:"l3NetworkUuid"`
		LastOpDate    string `json:"lastOpDate"`
		Nexthop       string `json:"nexthop"`
		Prefix        string `json:"prefix"`
	}

	IPRanges struct {
		CreateDate    string `json:"createDate"`
		EndIP         string `json:"endIp"`
		Gateway       string `json:"gateway"`
		IPVersion     int64  `json:"ipVersion"`
		L3NetworkUUID string `json:"l3NetworkUuid"`
		LastOpDate    string `json:"lastOpDate"`
		Name          string `json:"name"`
		Netmask       string `json:"netmask"`
		NetworkCidr   string `json:"networkCidr"`
		PrefixLen     int64  `json:"prefixLen"`
		StartIP       string `json:"startIp"`
		UUID          string `json:"uuid"`
	}
	NetworkService struct {
		L3NetworkUUID              string `json:"l3NetworkUuid"`
		NetworkServiceProviderUUID string `json:"networkServiceProviderUuid"`
		NetworkServiceType         string `json:"networkServiceType"`
	}
)

//获取云3层网络  创建云主机使用
func NetworkList(req *NetworkListReq) (resp *NetworkListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_NETWORK)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
