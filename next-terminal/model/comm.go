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

type (
	NextTerminalResp struct {
		Id        uint64 `json:"id"`
		Addr      string `json:"addr"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		ConnState int    `json:"conn_state"`
	}
)

//获取漏扫节点配置信息
func GetNextTerminalInfo() (resp *NextTerminalResp, err error) {
	var list []*subassemblynode.Subassemblynode
	list, _, err = subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     7, //堡垒机
		State:    "1",
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(list) == 0 {
		return resp, fmt.Errorf("该节点暂未添加，请添加后重试")
	}
	info := list[0]
	addr := utils.CheckHttpUrl(info.Addr, info.IsSsl == 1)
	resp = &NextTerminalResp{
		Addr:      addr,
		Username:  info.Key,
		Password:  info.Secret,
		Id:        info.Id,
		ConnState: info.ConnState,
	}
	return resp, err
}
