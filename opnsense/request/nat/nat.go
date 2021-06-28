package nat

import (
	"bytes"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

type (
	Nat1To1ListResp struct {
		ID        string `json:"id"`
		Interface string `json:"interface"` //接口
		Src       string `json:"src"`       //内部地址
		External  string `json:"external"`  //外部地址
		Dst       string `json:"dst"`       //目的地
		Descr     string `json:"descr"`     //描述
		Status    string `json:"status"`    //启用状态
	}
	Nat1To1InfoReq struct {
		Id string `json:"id"`
	}
	Nat1To1InfoResp struct {
		ID            string           `json:"id"`
		Disabled      bool             `json:"disabled"`      //启用
		Interface     []SelectedParams `json:"interface"`     //接口
		Type          []SelectedParams `json:"type"`          //类型
		External      string           `json:"external"`      //外部地址
		Srcnot        bool             `json:"srcnot"`        //源 反转
		Src           string           `json:"src"`           //内部地址
		Srcmask       []SelectedParams `json:"srcmask"`       //内部地址 掩码
		Dstnot        bool             `json:"dstnot"`        //目标反转
		Dst           []SelectedParams `json:"dst"`           //目的地
		Dstmask       []SelectedParams `json:"dstmask"`       //目的地掩码
		Category      []SelectedParams `json:"category"`      //分类 多选
		Descr         string           `json:"descr"`         //描述
		Natreflection []SelectedParams `json:"natreflection"` //NAT回流
	}
	SelectedParams struct {
		Name      string `json:"name"`
		Value     string `json:"value"`
		Selected  bool   `json:"selected"`   //选中状态
		DataOther bool   `json:"data_other"` //为true并选中时表示 此下拉选择为输入类型，value为输入的参数
	}
)

var client = resty.New() //.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(time.Second * 60)

//获取nat 1：1列表
func GetNat1To1List(apiKey *request.ApiKey) (list []*Nat1To1ListResp, err error) {
	url := fmt.Sprintf("http://%v%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_LIST_URL)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Get(url)
	//fmt.Println((resp.StatusCode()), err)
	if resp.StatusCode() == 200 {
		return ListMatch(bytes.NewReader(resp.Body()))
	}
	return list, err

}

//获取nat 1：1 详情
func GetNat1To1Info(req *Nat1To1InfoReq, apiKey *request.ApiKey) (info *Nat1To1InfoResp, err error) {
	url := fmt.Sprintf("http://%v%v?id=%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_INFO_URL, req.Id)
	client := request.GetHttpClient(apiKey)
	url = request.CheckHttpUrl(url, apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).Get(url)
	//Get(fmt.Sprintf("https://%v:%v%v?id=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_INFO_URL, req.Id))
	//fmt.Println((resp.StatusCode()), err)
	if resp.StatusCode() == 200 {
		return InfoMatch(bytes.NewReader(resp.Body()))
	}
	return info, err

}

//添加 nat 1：1
func AddNat1To1(req map[string]string, reqCateMap map[string][]string, apiKey *request.ApiKey) (res []string, err error) {
	url := fmt.Sprintf("http://%v%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_INFO_URL)
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
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_INFO_URL))
	//fmt.Println((resp.StatusCode()), err)
	//fmt.Println(resp.Body())
	if resp.StatusCode() == 200 {
		//匹配出现的错误
		return MatchSaveErr(bytes.NewReader(resp.Body()))
	}

	return res, err
}

//修改 nat 1：1
func EditNat1To1(req map[string]string, reqCateMap map[string][]string, apiKey *request.ApiKey) (res []string, err error) {
	urls := fmt.Sprintf("http://%v%v?id=%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_INFO_URL, req["id"])
	client := request.GetHttpClient(apiKey)
	urls = request.CheckHttpUrl(urls, apiKey)
	cates := url.Values{}
	if cate, ok := reqCateMap["category"]; ok {
		cates["category"] = cate
	}
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		SetFormData(req).
		SetFormDataFromValues(cates).Post(urls)
	//Post(fmt.Sprintf("https://%v:%v%v?id=%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_INFO_URL, req["id"]))
	//fmt.Println((resp.StatusCode()), err)
	//fmt.Println("cate",cates)
	//fmt.Println(string(resp.Body()))
	if resp.StatusCode() == 200 {
		//匹配出现的错误
		return MatchSaveErr(bytes.NewReader(resp.Body()))
	}

	return res, err
}

//启动 停止
func StartUpNat1To1(id string, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_STATUS_URL)
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
			"id":     id,
			"action": "toggle",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_STATUS_URL))
	if resp.StatusCode() == 200 {
		res = true
		Apply(apiKey)
	}
	return res, err
}

//删除nat 1：1
func DelNat1To1(id string, apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v:%v%v", request.UrlRemoveHttp(apiKey.Addr), _const.OPNSENSE_NAT_1TO1_STATUS_URL)
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
			"id":     id,
			"action": "del",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_STATUS_URL))
	if resp.StatusCode() == 200 {
		res = true
		Apply(apiKey)
	}
	return res, err
}

//应用修改
func Apply(apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("http://%v%v", apiKey.Addr, _const.OPNSENSE_NAT_1TO1_STATUS_URL)
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
			"apply": "Apply changes",
		}).Post(url)
	//Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_NAT_1TO1_STATUS_URL))

	if resp.StatusCode() == 200 {
		res = true
	}
	return res, err
}
