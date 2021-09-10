package host

//云盘信息
import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
)

type (
	DiskListReq struct {
		UUid string `json:"uuid"`
	}

	DiskListResp struct {
		Inventories []DiskInventories `json:"inventories"`
	}
	DiskInventories struct {
		AllocatorStrategy string `json:"allocatorStrategy"`
		CreateDate        string `json:"createDate"`
		Description       string `json:"description"`
		DiskSize          int64  `json:"diskSize"`
		LastOpDate        string `json:"lastOpDate"`
		Name              string `json:"name"`
		SortKey           int64  `json:"sortKey"`
		State             string `json:"state"`
		Type              string `json:"type"`
		UUID              string `json:"uuid"`
	}
)

//获取云计算规格  创建云主机使用
func DiskList(req *DiskListReq) (resp *DiskListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_DISK)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
