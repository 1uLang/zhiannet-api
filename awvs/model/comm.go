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

func ParseResp(resp []byte) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(resp) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	if _, isexist := ret["error"]; isexist {
		switch ret["error"] {
		case "Duplicate username":
			return nil, fmt.Errorf("账户已注册")
		case "Current password is invalid":
			return nil, fmt.Errorf("当前密码错误")
		case "The requested file was not found":
			return nil, fmt.Errorf("无效的用户id")
		}
		return nil, fmt.Errorf(ret["error"].(string))
	}
	return ret, nil
}

type (
	WebScanResp struct {
		Addr string `json:"addr"`
		Key  string `json:"key"`
		Id   uint64 `json:"id"`
	}
)

//获取漏扫节点配置信息
func GetWebScanInfo() (resp *WebScanResp, err error) {
	var list []*subassemblynode.Subassemblynode
	list, _, err = subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     4, //漏扫节点
		State:    "1",
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(list) == 0 {
		return resp, fmt.Errorf("获取漏扫节点错误")
	}
	info := list[0]
	addr := utils.CheckHttpUrl(info.Addr, info.IsSsl == 1)
	resp = &WebScanResp{
		Addr: addr,
		Key:  info.Key,
		Id:   info.Id,
	}
	return resp, err
}
