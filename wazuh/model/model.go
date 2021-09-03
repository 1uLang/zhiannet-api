package model

import "encoding/json"

func ToMap(obj interface{}) map[string]interface{} {

	ret := map[string]interface{}{}

	buf, _ := json.Marshal(obj)
	_ = json.Unmarshal(buf, &ret)
	return ret
}
