package request

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/utils"
	"net/http"
	"time"
)

type (
	ApiKey struct {
		Username   string
		Password   string
		Addr       string
		Port       string
		Cookie     string
		XCsrfToken string
		IsSsl      bool //是否使用ssl协议登陆
	}
	TestGlobalStatus struct {
		Plugins []string `json:"plugins"`
		System  string   `json:"system"`
	}
)

//检测是否可用
func (this *ApiKey) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("opnsense-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     2,
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取云防火墙节点信息失败")
		return
	}
	for _, v := range nodes {
		var conn int = 1
		logReq := &ApiKey{
			Username: v.Key,
			Password: v.Secret,
			Addr:     v.Addr,
			IsSsl:    v.IsSsl == 1,
		}
		logReq.Addr = utils.CheckHttpUrl(logReq.Addr, v.IsSsl == 1)
		tokenMap, err := Login(logReq)
		if tokenMap != nil {
			logReq.XCsrfToken = tokenMap["x-csrftoken"]
			logReq.Cookie = tokenMap["cookie"]
		}
		resp, err := GetGlobal(logReq)
		if err != nil || resp == nil {
			//调用接口失败 不可用
			conn = 0
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "云防火墙状态不可用",
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})
		}

		if conn != v.ConnState {
			subassemblynode.UpdateConnState(v.Id, conn)
		}
	}

}

//测试接口能否调用 全局接口
func GetGlobal(apiKey *ApiKey) (res *TestGlobalStatus, err error) {
	url := apiKey.Addr + _const.OPNSENSE_GLOBAL_STATUS_URL + fmt.Sprintf("%v", time.Now().Unix())
	client := GetHttpClient(apiKey)
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).Get(url)
	//Get("https://" + apiKey.Addr + ":" + apiKey.Port + _const.OPNSENSE_GLOBAL_STATUS_URL + fmt.Sprintf("%v", time.Now().Unix()))
	if err != nil {
		//fmt.Println(err)
		return res, err
	}
	if resp.StatusCode() == 200 {
		err = json.Unmarshal(resp.Body(), &res)
	}
	//fmt.Println(string(resp.Body()))
	return res, err
}
