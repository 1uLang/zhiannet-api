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
	"github.com/1uLang/zhiannet-api/utils"
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

		ReqType     string      //请求方式
		QueryParams interface{} //请求参数
	}
	LoginRes struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Path  string `json:"path"`
			Token string `json:"token"`
		} `json:"data"`
	}

	UserReq struct { //节点
		UserId      uint64 `json:"user_id"`
		AdminUserId uint64 `json:"admin_user_id"`
	}
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取token
func Login(req *LoginReq) (token string, err error) {
	client := GetHttpClient(req)
	url := utils.CheckHttpUrl(req.Addr+_const.AUDIT_LOGIN_URL, req.IsSsl)
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

func GetLoginInfo(audit *UserReq) (logReq *LoginReq, err error) {
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		State:    "1",
		Type:     6,
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取审计系统节点信息失败")
		return logReq, err
	}
	node := nodes[0]

	if audit.AdminUserId == 1 { //等保云 超级管理员
		logReq = &LoginReq{
			Name:     node.Key,
			Password: node.Secret,
			Addr:     node.Addr,
			IsSsl:    node.IsSsl == 1,
		}
	} else {
		//等保平台 其他用户
		logReq = &LoginReq{
			//Name: node.Key,
			//Password: node.Secret,
			Addr:  node.Addr,
			IsSsl: node.IsSsl == 1,
		}
		//获取关联的审计平台用户
		auditInfo, err := audit_user_relation.GetInfo(&audit_user_relation.AuditReq{
			AdminUserId: audit.AdminUserId,
			UserId:      audit.UserId,
		})
		if err != nil {
			err = fmt.Errorf("账号错误")
		}
		info, err := audit_user.GetInfo(&audit_user.AuditReq{UserId: auditInfo.AuditUserid})
		if err != nil {
			err = fmt.Errorf("获取用户信息错误")
			return logReq, err
		}
		logReq.Name = info.UserName
		logReq.Password = info.Pwd
	}

	key := fmt.Sprintf("audit-get-token-%v:%v", logReq.Addr, logReq.Name)

	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(logReq)
	}, 600, true)
	if err != nil {
		return
	}
	logReq.Token = fmt.Sprintf("%v", resp)
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
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     6,
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取审计系统节点信息失败")
		return
	}
	for _, v := range nodes {
		logReq := &LoginReq{
			Name:     v.Key,
			Password: v.Secret,
			Addr:     v.Addr,
			IsSsl:    v.IsSsl == 1,
		}
		token, err := Login(logReq)
		var conn int = 1
		if err != nil || token == "" {
			//登录失败 不可用
			conn = 0
		}
		if conn != v.ConnState {
			subassemblynode.UpdateConnState(v.Id, conn)
		}
	}

}
