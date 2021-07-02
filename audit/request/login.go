package request

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/audit/const"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user_relation"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	LoginReq struct {
		Name     string
		Password string
		Addr     string
		Port     string
		IsSsl    bool
		Token    string
	}
	LoginRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Path  string `json:"path"`
			Token string `json:"token"`
		} `json:"data"`
	}
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取token
func Login(req *LoginReq) (token string, err error) {
	client := GetHttpClient(req)
	url := req.Addr + _const.AUDIT_LOGIN_URL
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetFormData(map[string]string{
			"username": req.Name,
			"password": req.Password,
		}).Post(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error("err1=", err)
		return token, err
	}

	//fmt.Println("respon ",string(resp.Body()))
	var data = LoginRes{}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Error("err2=", err)
		return token, err
	}
	if data.Code != 0 {
		return token, err
	}
	token = data.Data.Token
	req.Token = token
	return token, err
}

func GetToken(audit *audit_user_relation.AuditReq) (Token string, err error) {
	var req *LoginReq
	//audit := &audit_user_relation.AuditReq{
	//	AdminUserId: 1,
	//}
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		State:    "1",
		Type:     6,
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取审计系统节点信息失败")
		return Token, err
	}
	node := nodes[0]

	if audit.AdminUserId == 1 { //等保云 超级管理员
		req = &LoginReq{
			Name:     node.Key,
			Password: node.Secret,
			Addr:     node.Addr,
			IsSsl:    node.IsSsl == 1,
		}
	} else {
		//等保平台 其他用户
		req = &LoginReq{
			//Name: node.Key,
			//Password: node.Secret,
			Addr:  node.Addr,
			IsSsl: node.IsSsl == 1,
		}
		//获取关联的审计平台用户
		auditInfo, err := audit_user_relation.GetInfo(audit)
		if err != nil {
			err = fmt.Errorf("账号错误")
		}
		info, err := audit_user.GetInfo(&audit_user.AuditReq{UserId: auditInfo.AuditUserid})
		if err != nil {
			err = fmt.Errorf("获取用户信息错误")
			return Token, err
		}
		req.Name = info.UserName
		req.Password = info.Pwd
	}

	key := fmt.Sprintf("audit-token-%v:%v", req.Addr, req.Name)

	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(req)
	}, 60, true)
	if err != nil {
		return
	}
	Token = fmt.Sprintf("%v", resp)
	return
}

//获取请求客户端
func GetHttpClient(req *LoginReq) *resty.Client {
	if req.IsSsl {
		Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return Client
}
