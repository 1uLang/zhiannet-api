package acl

import (
	"bytes"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
)

type (
	AclListResp struct {
		ID         string `json:"id"`
		Direction  string `json:"direction"`  //方向
		Interface  string `json:"interface"`  //接口
		Ipprotocol string `json:"ipprotocol"` //tcp/IP版本
		Protocol   string `json:"protocol"`   //协议
		Src        string `json:"src"`        //源
		SrcPort    string `json:"src_port"`   //源 端口
		Dst        string `json:"dst"`        //目标
		DstPort    string `json:"dst_port"`   //目标 端口
		Descr      string `json:"descr"`      //描述
		Type       string `json:"type"`       //策略
		Disabled   bool   `json:"disabled"`   //状态 禁用
	}
	AclInfoReq struct {
		ID string `json:"id"`
	}
	AclInfoResp struct {
		ID         string           `json:"id"`
		Type       []SelectedParams `json:"type"`       //操作
		Disabled   bool             `json:"disabled"`   //状态 禁用
		Quick      bool             `json:"quick"`      //快速
		Interface  []SelectedParams `json:"interface"`  //接口
		Direction  []SelectedParams `json:"direction"`  //方向
		Ipprotocol []SelectedParams `json:"ipprotocol"` //tcp/IP版本
		Protocol   []SelectedParams `json:"protocol"`   //协议
		Srcnot     bool             `json:"srcnot"`     //源 反转
		Src        []SelectedParams `json:"src"`        //源
		Srcmask    []SelectedParams `json:"srcmask"`    //源 掩码

		Dstnot  bool             `json:"dstnot"`  //目标反转
		Dst     []SelectedParams `json:"dst"`     //目标
		Dstmask []SelectedParams `json:"dstmask"` //目标 掩码

		Log bool `json:"log"` //日志
		//Category []SelectedParams `json:"category"` //分类
		Descr string `json:"descr"` //描述

	}
	SelectedParams struct {
		Name      string `json:"name"`
		Value     string `json:"value"`
		Selected  bool   `json:"selected"`   //选中状态
		DataOther bool   `json:"data_other"` //为true并选中时表示 此下拉选择为输入类型，value为输入的参数
	}
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

//获取acl列表
func GetAclList(Interface string, apiKey *request.ApiKey) (list []*AclListResp, err error) {
	url := fmt.Sprintf("http://%v%v?if=%v", apiKey.Addr, _const.OPNSENSE_ACL_LIST_URL, Interface)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).Get(url)
	//Get(fmt.Sprintf("https://%v:%v%v?if=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_LIST_URL, Interface))
	//fmt.Println((resp.StatusCode()), err)
	if resp.StatusCode() == 200 {
		return ListMatch(Interface, bytes.NewReader(resp.Body()))
	}
	return list, err

}

//获取acl 详情
func GetAclInfo(req *AclInfoReq, apiKey *request.ApiKey) (info *AclInfoResp, err error) {
	url := fmt.Sprintf("http://%v%v?id=%v", apiKey.Addr, _const.OPNSENSE_ACL_INFO_URL, req.ID)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).Get(url)
	//Get(fmt.Sprintf("https://%v:%v%v?id=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_INFO_URL, req.ID))
	//fmt.Println((resp.StatusCode()), err)
	if resp.StatusCode() == 200 {
		return InfoMatch(bytes.NewReader(resp.Body()))
	}
	return info, err

}

//添加acl
func AddAcl(req map[string]string, apiKey *request.ApiKey) (res []string, err error) {
	url := fmt.Sprintf("http://%v%v", apiKey.Addr, _const.OPNSENSE_ACL_INFO_URL)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(req).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_INFO_URL))
	//fmt.Println((resp.StatusCode()), err)
	//fmt.Println(resp.Body())
	if resp.StatusCode() == 200 {
		//匹配出现的错误
		return MatchSaveErr(bytes.NewReader(resp.Body()))
	}

	return res, err
}

func EditAcl(req map[string]string, apiKey *request.ApiKey) (tips []string, err error) {
	url := fmt.Sprintf("http://%v%v?if=%v&id=%v", apiKey.Addr, _const.OPNSENSE_ACL_INFO_URL, req["interface"], req["id"])
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(req).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v?if=%v&id=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_INFO_URL, req["interface"], req["id"]))
	//fmt.Println((resp.StatusCode()), err)
	//fmt.Println(resp.Body())
	if resp.StatusCode() == 200 {
		//匹配出现的错误
		return MatchSaveErr(bytes.NewReader(resp.Body()))
	}

	return tips, err
}

//启动 停止
func StartUpAcl(id, Interface string, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v%v", apiKey.Addr, _const.OPNSENSE_ACL_LIST_URL)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"id":  id,
			"act": "toggle",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_LIST_URL))
	if resp.StatusCode() == 200 {
		res = true
		Apply(Interface, apiKey)
	}
	return res, err
}

//删除
func DelAcl(id, Interface string, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v%v?if=%v", apiKey.Addr, _const.OPNSENSE_ACL_LIST_URL, Interface)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"id":  id,
			"act": "del",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v?if=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_LIST_URL, Interface))
	if resp.StatusCode() == 200 {
		res = true
		Apply(Interface, apiKey)
	}
	return res, err
}

//应用修改
func Apply(Interface string, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v%v?if=%v", apiKey.Addr, _const.OPNSENSE_ACL_LIST_URL, Interface)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(map[string]string{
			"act": "apply",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v?if=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_ACL_LIST_URL, Interface))

	if resp.StatusCode() == 200 {
		res = true
	}
	return res, err
}
