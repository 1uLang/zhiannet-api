package request

import (
	"crypto/sha512"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/utils"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type (
	LoginReq struct {
		Name     string
		Password string
		Addr     string
		Port     string
		IsSsl    bool
		UUID     string

		ReqType     string      //请求方式
		QueryParams interface{} //请求参数
	}
	LoginRes struct {
		Inventory struct {
			UUID        string `json:"uuid"`
			AccountUUID string `json:"accountUuid"`
			UserUUID    string `json:"userUuid"`
			ExpiredDate string `json:"expiredDate"`
			CreateDate  string `json:"createDate"`
		} `json:"inventory"`
	}

	UserReq struct { //节点
		UserId      uint64 `json:"user_id"`
		AdminUserId uint64 `json:"admin_user_id"`
	}

	//登录请求参数
	LogAccount struct {
		LogInByAccount LogInByAccount `json:"logInByAccount"`
	}
	LogInByAccount struct {
		Password    string `json:"password"`
		AccountName string `json:"accountName"`
	}
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取token
func Login(req *LoginReq) (uuid string, err error) {
	client := GetHttpClient(req)
	url := utils.CheckHttpUrl(req.Addr+_const.ZSTACK_LOGIN, req.IsSsl)

	login := LogAccount{
		LogInByAccount: LogInByAccount{
			Password:    Sha512(req.Password),
			AccountName: req.Name,
		},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(login).Put(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error("err1=", err)
		return uuid, err
	}

	//fmt.Println("respon ", string(resp.Body()), "code", resp.StatusCode())
	var data = LoginRes{}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Error("err2=", err)
		return uuid, err
	}
	if data.Inventory.UUID == "" {
		return uuid, err
	}
	uuid = data.Inventory.UUID
	req.UUID = uuid
	//fmt.Println("UUID+",uuid)
	return uuid, err
}

func GetLoginInfo(audit *UserReq) (logReq *LoginReq, err error) {
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		State:    "1",
		Type:     10,
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取云底座节点信息失败")
		return logReq, err
	}
	node := nodes[0]

	//等保云 超级管理员
	logReq = &LoginReq{
		Name:     node.Key,
		Password: node.Secret,
		Addr:     node.Addr,
		IsSsl:    node.IsSsl == 1,
	}

	key := fmt.Sprintf("zstack-get-token-%v:%v", logReq.Addr, logReq.Name)

	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(logReq)
	}, 600, true)
	if err != nil {
		return
	}
	logReq.UUID = fmt.Sprintf("%v", resp)
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
			fmt.Println("audit-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     10,
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取云底座节点信息失败")
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
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "云底座状态不可用",
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
					Body:      "云底座恢复可用状态",
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

func Sha512(pass string) (str string) {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha512.New()
	//输入数据
	hash.Write([]byte(pass))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	str = hex.EncodeToString(bytes)
	//返回哈希值
	return str
}

func login0() {
	req := &LoginReq{
		Addr:  "https://hids.zhiannet.com/sso/login",
		IsSsl: true,
	}
	client := GetHttpClient(req)
	url := utils.CheckHttpUrl(req.Addr, req.IsSsl)
	fmt.Println("url", url)
	reqMap := map[string]string{
		"callback": "jQuery3000044090075032042986_1629875256022",
		"ajax":     "yes",
		"sid":      "YTAyMGVlYTMtOWMwZi00YmI0LTgwOTYtOThmMmRkNGYwMjQz",
		"service":  "https://hids.zhiannet.com/passport/yunleiIndex",
		//"execution": "e1s1",
		//"_eventId":  "submit",
		//"access":    "1",
		"_": fmt.Sprintf("%v", time.Now().Unix()),
	}
	resp, err := client.SetDebug(true).R().
		//SetHeader("Content-Type", "application/json").
		SetQueryString(params2UrlEncode(reqMap)).
		Get(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error("err1=", err)
		return
	}

	fmt.Println("respon ", string(resp.Body()), "code", resp.StatusCode())
	//login1()
	//fmt.Println("UUID+",uuid)
	return
}

func login1(str string) {
	req := &LoginReq{
		Addr:  "https://hids.zhiannet.com/sso/login",
		IsSsl: true,
	}
	client := GetHttpClient(req)
	url := utils.CheckHttpUrl(req.Addr, req.IsSsl)
	fmt.Println("url", url)
	reqMap := map[string]string{
		"callback":  "jQuery3000044090075032042986_1629875256022",
		"userName":  "dengbao",
		"password":  "NmQ5ZGQwYzk1OGQxNDQxYmRlMzlkMzg3MjU5NWViYzE=",
		"vcode":     "1234",
		"sid":       "YTAyMGVlYTMtOWMwZi00YmI0LTgwOTYtOThmMmRkNGYwMjQz",
		"service":   "https://hids.zhiannet.com/passport/yunleiIndex",
		"execution": "e1s1",
		"_eventId":  "submit",
		"lt":        str,
		"access":    "1",
		"ajax":      "yes",
		"_":         fmt.Sprintf("%v", time.Now().Unix()),
	}
	resp, err := client.SetDebug(true).R().
		//SetHeader("Content-Type", "application/json").
		SetQueryString(params2UrlEncode(reqMap)).
		Get(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error("err1=", err)
		return
	}

	fmt.Println("respon ", string(resp.Body()), "code", resp.StatusCode())

	//fmt.Println("UUID+",uuid)
	return
}

func params2UrlEncode(params map[string]string) string {

	q := (&url.URL{}).Query()
	for k, v := range params {
		fmt.Println(k, v)
		q.Add(k, fmt.Sprintf("%v", v))
	}
	return q.Encode()
}
