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
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
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

func GetLoginInfo(nodeid uint64) (logReq *LoginReq, err error) {
	//nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
	//	State:    "1",
	//	Type:     11,
	//	PageNum:  1,
	//	PageSize: 1,
	//
	//})
	//if err != nil || len(nodes) == 0 {
	//	err = fmt.Errorf("获取apt检测节点信息失败")
	//	return logReq, err
	//}
	//node := nodes[0]

	node, err := subassemblynode.GetNodeInfoById(nodeid)
	if err != nil || node == nil {
		err = fmt.Errorf("获取apt检测节点信息失败")
		return logReq, err
	}
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
		PageSize: 99,
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

//登陆获取token
func Change() (err error) {
	//go  get  github.com/go-resty/resty/v2
	var Client = resty.New().SetDebug(true).SetTimeout(time.Second * 60)
	Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	loginUrl := "https://bptest.dengbao.cloud/login"
	url := "https://bptest.dengbao.cloud/settings/personal/changepassword"

	resp1, err := Client.R().Get(loginUrl)
	if err != nil {
		fmt.Println("err 1 ,", err)
		return
	}
	rgp := regexp.MustCompile(`data-requesttoken="(.*)"`)
	strs := rgp.FindStringSubmatch(string(resp1.Body()))
	if len(strs) < 2 {
		log.Fatal("更改密码错误")
	}
	reqToken := strs[1]
	fmt.Println(reqToken)
	//获取cookie
	Cookies := resp1.Cookies()
	cook := ""
	if len(Cookies) > 0 {
		fmt.Println(Cookies)
		for _, v := range Cookies {
			cook = fmt.Sprintf("%v;%v=%v", cook, v.Name, v.Value)
		}
		cook = strings.TrimLeft(cook, ";")
	}
	fmt.Println(cook)
	//return

	resp2, err := Client.R().SetCookies(Cookies).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		//SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		SetHeader("upgrade-insecure-requests", "1").
		SetQueryParams(map[string]string{
			"user":            "test_hanchan",
			"password":        "21ops.com123",
			"redirect_url":    "/settings/user/security",
			"timezone":        "Asia/Shanghai",
			"timezone_offset": "8",
			"requesttoken":    reqToken,
		}).
		//SetQueryString("requesttoken="+reqToken).
		Post(loginUrl)
	if err != nil {
		fmt.Println("err1=", err)
		return err
	}

	fmt.Println(resp2.StatusCode(), reqToken, string(resp2.Body()))
	if resp2.StatusCode() == 303 {
		//获取cookie
		fmt.Println("resp=", string(resp2.Body()))
	}
	return

	resp, err := Client.R().
		//SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		//SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		//SetHeader("method", "POST").
		//SetHeader("authority", "bptest.dengbao.cloud").
		//SetHeader("path", "/settings/personal/changepassword").
		//SetHeader("scheme", "https").
		//SetHeader("accept", "*/*").
		//SetHeader("accept-encoding", "gzip, deflate, br").
		//SetHeader("accept-language", "zh-CN,zh;q=0.9").
		//SetHeader("content-length", "72").
		//SetHeader("ocs-apirequest", "true").
		//SetHeader("origin", "https://bptest.dengbao.cloud").
		//SetHeader("sec-ch-ua", "\"Google Chrome\";v=\"93\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"93\"").
		//SetHeader("sec-ch-ua-mobile", "?0").
		//SetHeader("sec-ch-ua-platform", "\"macOS\"").
		//SetHeader("sec-fetch-dest", "empty").
		//SetHeader("sec-fetch-mode", "cors").
		//SetHeader("sec-fetch-site", "same-origin").
		SetHeader("x-requested-with", "XMLHttpRequest").
		SetHeader("requesttoken", "jEkk5om7ohaZGmI1EoEqL1UsiumLNKXsKrJuUkRLsN4=:wQsSjcuO0U7wQCZhZbNGeyYD3Iq6f927EvMMBhMp1K0=").
		SetHeader("cookie", "oc_sessionPassphrase=Xtey00gffPBNMWif%2Fh5uprJGuzFMHayQNNxhxSjkgjLs9tHj72hgdIlO1umypGNe8I9mgAtJ74%2BHFEgCj%2BRF1sh0QscCFBSruSZwyxOyyrDkcC7fFfPFxZOrjybjwDHl; __Host-nc_sameSiteCookielax=true; __Host-nc_sameSiteCookiestrict=true; nc_username=test_hanchan; ocdh9htx8nbo=53350d3d70240bafa1ca0f927e51a573; nc_token=r469Gqbv6XsF3S4eNfqNLkKNGWy0xxrl; nc_session_id=53350d3d70240bafa1ca0f927e51a573").
		SetFormData(map[string]string{
			"oldpassword":       "21pos.com.",
			"newpassword":       "21ops.com123",
			"newpassword-clone": "21ops.com123",
		}).
		Post(url)
	if err != nil {
		fmt.Println("err1=", err)
		return err
	}
	if resp.StatusCode() == 200 {
		//获取cookie
		fmt.Println("resp=", string(resp.Body()))
	}
	return err
}

func EditPwd(user, password, newPassword string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cli := &http.Client{
		Transport: tr,
	}
	loginUrl := "https://bptest.dengbao.cloud/login?redirect_url=/settings/user/security"
	//secURL := `http://localhost:8080/settings/user/security`
	req, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Authorization", "Basic YWRtaW46VGVzdEB0ZXN0MTIzNDU2")
	//req.Header.Set("OCS-APIRequest", "true")

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(body))
	rgp := regexp.MustCompile(`data-requesttoken="(.*)"`)
	strs := rgp.FindStringSubmatch(string(body))
	if len(strs) < 2 {
		log.Fatal("更改密码错误")
	}
	reqToken := strs[1]
	fmt.Println(reqToken)
	cookies1 := resp.Cookies()
	cookie1 := fmt.Sprintf("%v", cookies1)
	cookie1 = string([]rune(cookie1)[1 : len(cookie1)-1])
	//fmt.Println(cookies1)
	fmt.Println(cookie1)

	req, err = http.NewRequest("POST", "https://bptest.dengbao.cloud/login", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie", cookie1)

	queryParam := req.URL.Query()
	queryParam.Add("user", user)
	queryParam.Add("password", password)
	queryParam.Add("redirect_url", "/settings/user/security")
	queryParam.Add("timezone", "Asia/Shanghai")
	queryParam.Add("timezone_offset", "8")
	queryParam.Add("requesttoken", reqToken)
	req.URL.RawQuery = queryParam.Encode()
	fmt.Println(queryParam.Encode())
	resp, err = cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	if resp.StatusCode != 303 {
		loginBody, _ := io.ReadAll(resp.Body)
		fmt.Println("登陆失败", string(loginBody))
		return
	}

	cookies := resp.Cookies()
	cookie := fmt.Sprintf("%v", cookies)
	cookie = string([]rune(cookie)[1 : len(cookie)-1])
	fmt.Println(cookie)

	return
	req, err = http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	rgp = regexp.MustCompile(`data-requesttoken="(.*)"`)
	strs = rgp.FindStringSubmatch(string(body))
	if len(strs) < 2 {
		log.Fatal("更改密码错误22")
	}
	reqToken = strs[1]
	fmt.Println(reqToken)
	return

	changepwdURL := `https://bptest.dengbao.cloud/settings/personal/changepassword`
	req, err = http.NewRequest("POST", changepwdURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("requesttoken", reqToken)
	req.Header.Add("OCS-APIRequest", "true")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", cookie)
	queryParam = req.URL.Query()
	queryParam.Add("oldpassword", "Test@test123456")
	queryParam.Add("newpassword", "Test@test1234567")
	queryParam.Add("newpassword-clone", "Test@test1234567")
	req.URL.RawQuery = queryParam.Encode()

	resp, err = cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
}
