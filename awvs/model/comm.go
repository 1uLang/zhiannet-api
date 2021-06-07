package model

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(string(resp))
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
