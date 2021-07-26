package black_white_list

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"net/http"
	"strings"
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

type (
	BWListReq struct { //黑白名单请求参数
		Addr string
		Page int
	}

	StatusBwlist struct { //黑白名单列表响应参数
		Address  string `xml:"address"`
		Page     string `xml:"page"`
		Total    string `xml:"total"`
		SortType string `xml:"sort_type"`
		Bwlist   []struct {
			Address string `xml:"address"`
			Flags   string `xml:"flags"`
			Hits    string `xml:"hits"`
			Comment string `xml:"comment"`
		} `xml:"bwlist" json:"bwlist,omitempty"`
	}

	EditBWReq struct {
		Addr  []string
		White bool
	}

	//新增修改 返回的操作
	Success struct {
		Delay  string `xml:"delay"`
		Info   string `xml:"info"`
		URL    string `xml:"url"`
		Params string `xml:"params"`
	}
	Failure struct {
		Info   string `xml:"info"`
		URL    string `xml:"url"`
		Params string `xml:"params"`
	}
)

//黑白名单列表
func BWList(req *BWListReq, loginReq *request.LoginReq, retry bool) (res *StatusBwlist, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_BWLIST_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "select",                    //查询
		"param_page":        fmt.Sprintf("%v", req.Page), //分页
		"param_address":     req.Addr,                    //单个IP查询
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_BWLIST_URL)

	fmt.Println(string(resp.Body()), err)

	err = xml.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	//if len(apiRes.Host) == 0 { //可能登陆信息过期
	//	failure := &audit_db.Failure{}
	//	xml.Unmarshal(resp.Body(), &failure)
	//	if retry && failure.Info == _const.FAILURE_INFO {
	//		return HostList(req, loginReq, false)
	//	}
	//}
	return res, err
}

//添加黑白名单
func AddBW(req *EditBWReq, loginReq *request.LoginReq, retry bool) (res *Success, err error) {
	addr := ""
	if len(req.Addr) > 0 {
		if req.White { //白名单
			addr = "+" + req.Addr[0]
		} else { //黑名单
			addr = req.Addr[0]
		}
	}
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_BWLIST_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "submit", //添加
		"param_address":     addr,     //单个IP
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_BWLIST_URL)

	fmt.Println(string(resp.Body()), err)

	err = xml.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	if res.Delay == "" { //说明错误
		return res, fmt.Errorf("地址错误，请输入正确的黑白名单地址")
	}
	return res, err
}

//删除黑白名单
func DeleteBW(req *EditBWReq, loginReq *request.LoginReq, retry bool) (res *Success, err error) {
	addr := strings.Join(req.Addr, ",")
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_BWLIST_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "delete", //删除
		"param_address":     addr,     //单个IP
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_BWLIST_URL)

	fmt.Println(string(resp.Body()), err)

	err = xml.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
