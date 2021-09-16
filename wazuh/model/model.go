package model

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/utils"
)

func ToMap(obj interface{}) map[string]interface{} {

	ret := map[string]interface{}{}

	buf, _ := json.Marshal(obj)
	_ = json.Unmarshal(buf, &ret)
	return ret
}

type WazuhInfoResp struct {
	Addr      string
	Username  string
	Password  string
	Id        uint64
	ConnState int
}

//获取漏扫节点配置信息
func GetWazuhInfo() (resp *WazuhInfoResp, err error) {
	var list []*subassemblynode.Subassemblynode
	list, _, err = subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     12, //主机防护系统
		State:    "1",
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(list) == 0 {
		return resp, fmt.Errorf("获取漏扫节点错误")
	}
	info := list[0]
	addr := utils.CheckHttpUrl(info.Addr, info.IsSsl == 1)
	resp = &WazuhInfoResp{
		Addr:      addr,
		Username:  info.Key,
		Password:  info.Secret,
		Id:        info.Id,
		ConnState: info.ConnState,
	}
	return resp, err
}
