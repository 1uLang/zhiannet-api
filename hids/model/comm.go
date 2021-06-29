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

func ParseResp(resp []byte, retObj ...interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(resp) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(resp, &ret)

	if err != nil {
		return nil, err
	}

	reqCode, isExist := ret["reqCode"]

	if isExist {
		if reqCode == 400 {
			return nil, fmt.Errorf("%s", ret["msg"])
		}
		if ret["returnCode"] != "1" {
			return nil, fmt.Errorf("%v", ret["returnMsg"])
		}
	}
	data, isExist := ret["data"]
	if isExist && len(retObj) > 0 {
		buf, _ := json.Marshal(data)
		err = json.Unmarshal(buf, retObj[0])
	}
	return ret, nil
}

type (
	HidsResp struct {
		Addr   string `json:"addr"`
		AppId  string `json:"app_id"`
		Secret string `json:"secret"`
	}
)

//获取漏扫节点配置信息
func GetHidsInfo() (resp *HidsResp, err error) {
	var list []*subassemblynode.Subassemblynode
	list, _, err = subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     5, //主机防护系统
		State:    "1",
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(list) == 0 {
		return resp, fmt.Errorf("获取漏扫节点错误")
	}
	info := list[0]
	addr := utils.CheckHttpUrl(info.Addr, info.IsSsl == 1)
	resp = &HidsResp{
		Addr:   addr,
		AppId:  info.Key,
		Secret: info.Secret,
	}
	return resp, err
}
