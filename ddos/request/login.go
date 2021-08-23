package request

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/utils"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var Cookie string = "%2D%88%D9%3A%19%87%9C%28%8A%AF%EA%F8v%9CB%95%87%9E%8A3%98%EF%ED%F0%8C%075KAPfR%8B%82%85%D0%CC%AC4%00%96%DB%AE%88f%3E%7C%D6v1%C5%2D%A1%BE%7BQ%E4%B4u%D9%F9%E5%EBo%F9Qp%133l%C6%93%97%99%F2%DDdI%5Bx%AC%89%F6%05ha%ADw%16%10%F873%ABq%E4%F3lT%7D%E7%7F9%11IxU%BA%21%B7%2C%DC"

type (
	LoginReq struct {
		Name     string
		Password string
		Addr     string
		Port     string
		IsSsl    bool
	}
	LoginRes struct {
		Redirect `xml:"redirect"`
		Failure  `xml:"failure"`
	}
	Redirect struct {
		Page string `xml:"page"`
	}
	Failure struct {
		Info   string `xml:"info"`
		Url    string `xml:"url"`
		Params string `xml:"params"`
	}
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取cookie
func Login(req *LoginReq) (string, error) {
	var err error
	// https://182.131.30.171:28443/cgi-bin/login.cgi
	client := GetHttpClient(req)
	url := req.Addr + _const.DDOS_LOGIN_URL
	resp, err := client.R().
		SetHeader("Content-Type", "multipart/form-data; boundary=<calculated when audit_db is sent>").
		SetFormData(map[string]string{
			"param_type":     "login",
			"param_username": req.Name,
			"param_password": req.Password,
		}).Post(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error(err)
		return Cookie, err
	}

	//logrus.Info(string(resp.Body()))
	var data = LoginRes{}
	err = xml.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Error(err)
		return Cookie, err
	}
	//logrus.Info(data)
	if data.Failure.Info != "" {
		logrus.Debug(data.Failure)
		err = fmt.Errorf(data.Failure.Info)
		return Cookie, err
	}
	if len(resp.Cookies()) > 0 {
		cook := resp.Cookies()[0]
		Cookie = cook.Value
	}
	fmt.Println(Cookie)

	return Cookie, err
	//logrus.Info( err)
}

func GetCookie(req *LoginReq) (cookie string) {
	var err error
	key := fmt.Sprintf("ddos-cookie-%v:%v", req.Addr, req.Name)
	//cache.CheckCache(key, Login(req), 3600, true)
	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(req)
	}, 600, true)
	if err != nil {
		return
	}
	cookie = fmt.Sprintf("%v", resp)
	return
}

//获取请求客户端
func GetHttpClient(req *LoginReq) *resty.Client {
	if req.IsSsl {
		Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return Client
}

//检测是否可用
func (this *LoginReq) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ddos-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     1,
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取ddos节点信息失败")
		return
	}
	for _, v := range nodes {
		logReq := &LoginReq{
			Name:     v.Key,
			Password: v.Secret,
			Addr:     v.Addr,
			IsSsl:    v.IsSsl == 1,
		}
		logReq.Addr = utils.CheckHttpUrl(logReq.Addr, v.IsSsl == 1)
		token, err := Login(logReq)
		var conn int = 1
		if err != nil || token == "" {
			//登录失败 不可用
			conn = 0
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "DDoS防护状态不可用",
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
			if conn == 1 {
				edge_messages.Add(&edge_messages.Edgemessages{
					Level:     "success",
					Subject:   "组件状态恢复正常",
					Body:      "DDoS防护恢复可用状态",
					Type:      "AdminAssembly",
					Params:    "{}",
					Createdat: uint64(time.Now().Unix()),
					Day:       time.Now().Format("20060102"),
					Hash:      "",
					Role:      "admin",
				})
			}
		}
	}

}
