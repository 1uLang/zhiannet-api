package request

import (
	"crypto/tls"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/common/util"
	_const "github.com/1uLang/zhiannet-api/maltrail/const"
	"github.com/1uLang/zhiannet-api/utils"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type (
	LoginReq struct {
		Name     string
		Password string
		Addr     string
		Port     string
		IsSsl    bool
		Cookie   string
		Nonce    string

		ReqType     string            //请求方式
		QueryParams map[string]string //请求参数
	}
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取token
func Login(req *LoginReq) (token string, err error) {
	req.Nonce, req.Password = GetNonce(req.Password)

	client := GetHttpClient(req)
	url := utils.CheckHttpUrl(req.Addr+_const.MALTRAIL_LOGIN, req.IsSsl)
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("X-Requested-With", "XMLHttpRequest").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		//SetHeader("Accept", "text/plain, */*; q=0.01").
		//SetHeader("Accept-Encoding", "gzip, deflate").
		//SetHeader("Accept-Language", "zh-CN,zh;q=0.9").
		//SetHeader("Connection", "keep-alive").
		//SetHeader("Cookie", "maltrail_sessid=c8088ac7c7d7ffeb776e71082a804e5e").
		//SetHeader("Host", "156.240.95.221:8338").
		//SetHeader("Origin", "http://156.240.95.221:8338").
		//SetHeader("Referer", "http://156.240.95.221:8338/").
		//SetBody("username=admin&hash=9d9c3e66d1f2ca9f3bf834f05f0dc7fcf05dfaf29d2dc9a64e1ef29c5efd6560&nonce=RjhhNHbycyj7").
		SetQueryParams(map[string]string{
			"username": req.Name,
			"hash":     req.Password,
			"nonce":    req.Nonce,
		}).
		Post(url)
	//Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error("err1=", err)
		return token, err
	}
	if resp.StatusCode() == 200 {
		//获取cookie
		Cookies := resp.Cookies()
		if len(Cookies) > 0 {
			token = Cookies[0].Name + "=" + Cookies[0].Value
		}

		//fmt.Println("login in Cookie=", Cookie)
	}
	return token, err
}

func GetLoginInfo() (logReq *LoginReq, err error) {
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		State:    "1",
		Type:     11,
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取apt检测节点信息失败")
		return logReq, err
	}
	node := nodes[0]

	logReq = &LoginReq{
		Name:     node.Key,
		Password: node.Secret,
		Addr:     node.Addr,
		IsSsl:    node.IsSsl == 1,
	}

	key := fmt.Sprintf("apt-get-token-%v:%v", logReq.Addr, logReq.Name)

	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(logReq)
	}, 600, true)
	if err != nil {
		return
	}
	logReq.Cookie = fmt.Sprintf("%v", resp)
	return
}

//获取请求客户端
func GetHttpClient(req *LoginReq) *resty.Client {
	if req.IsSsl {
		Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return Client
}

func GetNonce(password string) (nonce, pass string) {
	NONCE_ALPHABET := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 12; i++ {

		num := rand.Intn(62)
		if num > 62 {
			num = 62
		}

		nonce += string(NONCE_ALPHABET[num])

	}
	//nonce = "4JAuDFEXxr77"
	//pa := util.GetSHA256HashCode([]byte(password))
	pass = util.GetSHA256HashCode([]byte(password + nonce))
	//fmt.Println(nonce)
	return
}

//检测是否可用
func (this *LoginReq) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("maltrail-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     11,
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取apt节点信息失败")
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
				Body:      "apt节点状态不可用",
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
					Body:      "apt节点恢复可用状态",
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
