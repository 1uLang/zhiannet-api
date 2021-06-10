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

func ParseResp(resp []byte,retObj ...interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(resp) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(resp, &ret)
	fmt.Println(string(resp))
	if err != nil {
		return nil, err
	}
	if ret["reqCode"] == 400 {
		return nil, fmt.Errorf("%v",ret["msg"])
	}
	if len(retObj) > 0 {
		buf,_ := json.Marshal(ret["data"])
		err = json.Unmarshal(buf,retObj[0])
	}
	return ret, nil
}

func ParseResp2(resp []byte) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(resp) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(resp, ret)
	fmt.Println(string(resp))
	if err != nil {
		return nil, err
	}
	if ret["returnCode"] != "1" {
		return nil, fmt.Errorf("%v",ret["returnMsg"])
	}
	return ret["data"].(map[string]interface{}), nil
}