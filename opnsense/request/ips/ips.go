package ips

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
	"strings"
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
	EditActionIpsReq struct {
		Sid    int64  `json:"sid"`
		Action string `json:"action"`
	}
	EditResp struct {
		Result string `json:"result"`
	}
	ApplyResp struct {
		Status string `json:"status"`
	}
	IpsAlarmReq struct {
		IpsReq
		FileId string `json:"fileid"`
	}
	IpsAlarmListResp struct {
		Rows []struct {
			Timestamp string `json:"timestamp"`
			FlowID    int64  `json:"flow_id"`
			InIface   string `json:"in_iface"`
			EventType string `json:"event_type"`
			SrcIP     string `json:"src_ip"`
			SrcPort   int    `json:"src_port"`
			DestIP    string `json:"dest_ip"`
			DestPort  int    `json:"dest_port"`
			Proto     string `json:"proto"`
			Alert     string `json:"alert"`
			AppProto  string `json:"app_proto"`
			Flow      struct {
				PktsToserver  int    `json:"pkts_toserver"`
				PktsToclient  int    `json:"pkts_toclient"`
				BytesToserver int    `json:"bytes_toserver"`
				BytesToclient int    `json:"bytes_toclient"`
				Start         string `json:"start"`
			} `json:"flow"`
			Filepos     int    `json:"filepos"`
			Fileid      string `json:"fileid"`
			AlertSid    int    `json:"alert_sid"`
			AlertAction string `json:"alert_action"`
		} `json:"rows"`
		TotalRows int    `json:"total_rows"`
		Origin    string `json:"origin"`
		RowCount  int    `json:"rowCount"`
		Total     int    `json:"total"`
		Current   int    `json:"current"`
	}
	IpsAlarmTimeResp struct {
		Size     int         `json:"size"`
		Modified string      `json:"modified"`
		Filename string      `json:"filename"`
		Sequence interface{} `json:"sequence"`
	}

	RuleListResp struct {
		Rows []struct {
			Description      string      `json:"description"`
			Filename         string      `json:"filename"`
			DocumentationURL string      `json:"documentation_url"`
			Documentation    string      `json:"documentation"`
			ModifiedLocal    interface{} `json:"modified_local"`
			Enabled          string      `json:"enabled"`
		} `json:"rows"`
		RowCount int `json:"rowCount"`
		Total    int `json:"total"`
		Current  int `json:"current"`
	}

	FirmwareInfo struct {
		ProductID      string `json:"product_id"`
		ProductVersion string `json:"product_version"`
		Package        []struct {
			Name       string `json:"name"`
			Version    string `json:"version"`
			Comment    string `json:"comment"`
			Flatsize   string `json:"flatsize"`
			Locked     string `json:"locked"`
			Automatic  string `json:"automatic"`
			License    string `json:"license"`
			Repository string `json:"repository"`
			Origin     string `json:"origin"`
			Provided   string `json:"provided"`
			Installed  string `json:"installed"`
			Path       string `json:"path"`
			Configured string `json:"configured"`
		} `json:"package"`
		Plugin []struct {
			Name       string `json:"name"`
			Version    string `json:"version"`
			Comment    string `json:"comment"`
			Flatsize   string `json:"flatsize"`
			Locked     string `json:"locked"`
			Automatic  string `json:"automatic"`
			License    string `json:"license"`
			Repository string `json:"repository"`
			Origin     string `json:"origin"`
			Provided   string `json:"provided"`
			Installed  string `json:"installed"`
			Path       string `json:"path"`
			Configured string `json:"configured"`
		} `json:"plugin"`
		Changelog []struct {
			Series  string `json:"series"`
			Version string `json:"version"`
			Date    string `json:"date"`
		} `json:"changelog"`
		Product struct {
			ProductAbi            string      `json:"product_abi"`
			ProductArch           string      `json:"product_arch"`
			ProductCheck          interface{} `json:"product_check"`
			ProductCopyrightOwner string      `json:"product_copyright_owner"`
			ProductCopyrightURL   string      `json:"product_copyright_url"`
			ProductCopyrightYears string      `json:"product_copyright_years"`
			ProductCrypto         string      `json:"product_crypto"`
			ProductEmail          string      `json:"product_email"`
			ProductFlavour        string      `json:"product_flavour"`
			ProductHash           string      `json:"product_hash"`
			ProductID             string      `json:"product_id"`
			ProductMirror         string      `json:"product_mirror"`
			ProductName           string      `json:"product_name"`
			ProductNickname       string      `json:"product_nickname"`
			ProductRepos          string      `json:"product_repos"`
			ProductSeries         string      `json:"product_series"`
			ProductTime           string      `json:"product_time"`
			ProductVersion        string      `json:"product_version"`
			ProductWebsite        string      `json:"product_website"`
		} `json:"product"`
	}
)

//var client = resty.New() //.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(time.Second * 2)

//获取ips规则列表
func GetIpsList(req *IpsReq, apiKey *request.ApiKey) (list *IpsListResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_LIST_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"current":      req.Current, //
			"rowCount":     req.RowCount,
			"searchPhrase": req.SearchPhrase,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//编辑 启用｜停用 规则
func EditIps(req *EditIpsReq, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("%v%v/%v", apiKey.Addr, _const.OPNSENSE_IPS_EDIT_URL, req.Sid)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	go ApplyIps(apiKey)
	return editRes.Result == "saved", err
}

//删除 规则
func DelIps(req *DelIpsReq, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr,
		fmt.Sprintf(_const.OPNSENSE_IPS_DEL_URL, strings.Join(req.Sid, ",")))
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	go ApplyIps(apiKey)
	return editRes.Result == "saved", err
}

//编辑 启用｜停用 规则
func ApplyIps(apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_APPLY_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Post(url)
	//fmt.Println(resp.StatusCode(), string(resp.Body()), err)
	editRes := ApplyResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	return editRes.Status != "", err
}

//修改操作方法
func EditActionIps(req *EditActionIpsReq, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("%v%v/%v", apiKey.Addr, _const.OPNSENSE_IPS_ACTIOB_URL, req.Sid)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"action": req.Action,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	go ApplyIps(apiKey)
	return editRes.Result == "saved", err
}

//获取ips-报警 列表
func GetIpsAlarmList(req *IpsAlarmReq, apiKey *request.ApiKey) (list *IpsAlarmListResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_ALARM_LIST_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"current":      req.Current, //
			"rowCount":     req.RowCount,
			"searchPhrase": req.SearchPhrase,
			"fileid":       req.FileId,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//报警下拉时间
func GetIpsAlarmTime(apiKey *request.ApiKey) (list []*IpsAlarmTimeResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_ALARM_TIME_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Get(url)
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//获取接口名
func GetIpsAlarmIface(apiKey *request.ApiKey) (list string, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_ALARM_IFACE_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Get(url)
	//fmt.Println(string(resp.Body()), err)
	//json 的key不固定，所以返回json字符串，用时按需取
	list = string(resp.Body())
	return list, err
}

//获取规则列表
func GetIpsRule(req *IpsReq, apiKey *request.ApiKey) (list *RuleListResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_IPS_RULE)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		//SetFormData(map[string]string{
		//	"current":      req.Current, //
		//	"rowCount":     req.RowCount,
		//	"searchPhrase": req.SearchPhrase,
		//}).
		Get(url)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//获取固件版本 suricata
func GetFirmwareInfo(apiKey *request.ApiKey) (list *FirmwareInfo, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_FIRMWARE)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Get(url)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}
