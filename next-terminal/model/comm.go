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
	JumpserverResp struct {
		Addr     string `json:"addr"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

//获取漏扫节点配置信息
func GetJumpserverInfo() (resp *JumpserverResp, err error) {
	var list []*subassemblynode.Subassemblynode
	list, _, err = subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     7, //堡垒机
		State:    "1",
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(list) == 0 {
		return resp, fmt.Errorf("获取堡垒机节点错误")
	}
	info := list[0]
	addr := utils.CheckHttpUrl(info.Addr, info.IsSsl == 1)
	resp = &JumpserverResp{
		Addr:     addr,
		Username: info.Key,
		Password: info.Secret,
	}
	return resp, err
}
