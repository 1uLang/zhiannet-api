package host

//规格信息
import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
)

type (
	SpecListReq struct {
		UUid string `json:"uuid"`
	}

	SpecListResp struct {
		Inventories []SpecInventories `json:"inventories"`
	}
	SpecInventories struct {
		AllocatorStrategy string `json:"allocatorStrategy"`
		CPUNum            int64  `json:"cpuNum"`
		CPUSpeed          int64  `json:"cpuSpeed"`
		CreateDate        string `json:"createDate"`
		Description       string `json:"description"`
		LastOpDate        string `json:"lastOpDate"`
		MemorySize        int64  `json:"memorySize"`
		Name              string `json:"name"`
		SortKey           int64  `json:"sortKey"`
		State             string `json:"state"`
		Type              string `json:"type"`
		UUID              string `json:"uuid"`
	}
)

//获取云计算规格  创建云主机使用
func SpecList(req *SpecListReq) (resp *SpecListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_SPEC)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
