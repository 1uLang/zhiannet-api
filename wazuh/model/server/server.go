package server

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/request"
)

type InfoResp struct {
	Title       string `json:"title"`
	ApiVersion  string `json:"api_version"`
	Revision    int    `json:"revision"`
	LicenseName string `json:"license_name"`
	LicenseUrl  string `json:"license_url"`
	Hostname    string `json:"hostname"`
	Timestamp   string `json:"timestamp"`
}

func Info(req *request.Request) (*InfoResp, error) {

	req.Method = "get"
	req.Params = nil

	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf(resp.Message)
	}
	info := InfoResp{}
	bytes, _ := json.Marshal(resp.Data)
	fmt.Println(string(bytes))
	_ = json.Unmarshal(bytes, &info)
	return &info, err
}
