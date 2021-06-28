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
