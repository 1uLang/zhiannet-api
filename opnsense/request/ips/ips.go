package ips

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

type (
	IpsReq struct {
		Current      string //页数
		RowCount     string //每页条数
		SearchPhrase string //关键词
	}
	IpsListResp struct {
		Current    int `json:"current"`
		Parameters struct {
			FilterTxt string `json:"filter_txt"`
			Limit     string `json:"limit"`
			Offset    string `json:"offset"`
			SortBy    string `json:"sort_by"`
		} `json:"parameters"`
		RowCount int `json:"rowCount"`
		Rows     []struct {
			Action         string      `json:"action"`
			ActionDefault  string      `json:"action_default"`
			Classtype      string      `json:"classtype"`
			CreatedAt      string      `json:"created_at"`
			Enabled        int         `json:"enabled"`
			EnabledDefault int         `json:"enabled_default"`
			Gid            interface{} `json:"gid"`
			MatchedPolicy  string      `json:"matched_policy"`
			Msg            string      `json:"msg"`
			Reference      string      `json:"reference"`
			Rev            int         `json:"rev"`
			Sid            int         `json:"sid"`
			Source         string      `json:"source"`
			Status         string      `json:"status"`
			UpdatedAt      string      `json:"updated_at"`
		} `json:"rows"`
		Total int `json:"total"`
	}
	EditIpsReq struct {
		Sid int64 `json:"sid"`
	}
	DelIpsReq struct {
		Sid []string `json:"sid"`
	}
	EditResp struct {
		Result string `json:"result"`
	}
)

var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(time.Second * 2)

//获取ips规则列表
func GetIpsList(req *IpsReq, apiKey *request.ApiKey) (list *IpsListResp, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		SetFormData(map[string]string{
			"current":      req.Current, //
			"rowCount":     req.RowCount,
			"searchPhrase": req.SearchPhrase,
		}).
		Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_IPS_LIST_URL))
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//编辑 启用｜停用 规则
func EditIps(req *EditIpsReq, apiKey *request.ApiKey) (res bool, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		Post(fmt.Sprintf("https://%v:%v%v/%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_IPS_EDIT_URL, req.Sid))
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	return editRes.Result == "saved", err
}

//删除 规则
func DelIps(req *DelIpsReq, apiKey *request.ApiKey) (res bool, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port,
			fmt.Sprintf(_const.OPNSENSE_IPS_DEL_URL, strings.Join(req.Sid, ","))))
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	return editRes.Result == "saved", err
}
