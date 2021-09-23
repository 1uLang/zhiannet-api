package groups

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"strings"
)

type ListResp struct {
	AffectedItems []struct {
		ConfigSum string `json:"configSum"`
		Count     int64  `json:"count"`
		MergedSum string `json:"mergedSum"`
		Name      string `json:"name"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

const groups_api_url = "/groups"

func Create(req *request.Request, name string) error {

	req.Method = "post"
	req.Path = groups_api_url
	req.Params = map[string]interface{}{
		"group_id": name,
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Error != 0 {
		return fmt.Errorf("添加失败：%s", resp.Message)
	}
	return nil
}
func List(req *request.Request, name ...string) (*ListResp, error) {
	req.Method = "get"
	req.Path = groups_api_url
	req.Params = nil

	if len(name) > 0 {
		req.Params = map[string]interface{}{
			"groups_list": name[0],
		}
	}

	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("查询失败：%s", resp.Message)
	}
	list := &ListResp{}

	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func Delete(req *request.Request, name []string) error {
	req.Method = "delete"
	req.Path = groups_api_url
	req.Params = map[string]interface{}{
		"groups_list": strings.Join(name, ","),
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Error != 0 {
		return fmt.Errorf("删除失败：%s", resp.Message)
	}
	return nil
}
