package request

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	authorization_template = "Bearer "

	token_path = "/security/user/authenticate?raw=true"

	server_api_port = ":55000"

	kbn_version = "7.10.2"

	server_api_username = "wazuh"
	server_api_password = "wazuh"
)

type Request struct {
	Method   string
	url      string
	Path     string
	Params   map[string]interface{}
	Headers  map[string]string
	UserName string
	Password string
}
type Response struct {
	Data struct {
		Affected      []interface{} `json:"affected_items"`
		TotalAffected float64       `json:"total_affected_items"`
		Failed        []interface{} `json:"failed_items"`
		TotalFailed   float64       `json:"total_failed_items"`
	}
	Error int `json:"error"`
}
type RetInfo struct {
	Title   string      `json:"title"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   int         `json:"error"`
}

var req_mutex sync.RWMutex
var req = Request{
	Method: "get",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

//InitServerUrl 初始化wazuh服务器地址
func InitServerUrl(url string) error {
	req_mutex.Lock()
	req.url = url
	req_mutex.Unlock()
	return nil
}

//InitToken 初始化 token 所需的当前用户的账号密码
func InitToken(username, password string) error {

	//将改用户的账号密码存入程序缓存中。重启时自动清空
	req_mutex.Lock()
	req.UserName = username
	req.Password = password
	req_mutex.Unlock()
	_, err := req.token()
	if err != nil {
		return fmt.Errorf("初始化token失败：%v", err)
	}
	return nil
}

//NewRequest 创建一个请求头
func NewRequest() (*Request, error) {
	req_mutex.Lock()
	defer req_mutex.Unlock()
	if req.url == "" {
		return nil, fmt.Errorf("未配置主机防护 服务器地址")
	}

	if req.UserName == "" {
		return nil, fmt.Errorf("未配置主机防护系统管理员账号")
	}
	if req.Password == "" {
		return nil, fmt.Errorf("未配置主机防护系统管理员密码")
	}

	return &Request{Headers: req.Headers, url: req.url, UserName: req.UserName, Password: req.Password}, nil
}

//updateToken 获取更新 token
func (this *Request) updateToken() (string, error) {

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	auth := base64.StdEncoding.EncodeToString([]byte(server_api_username + ":" + server_api_password))

	req, err := http.NewRequest("GET", this.url+server_api_port+token_path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tokenBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	token := string(tokenBytes)
	err = redis_cache.SetCache(server_api_username+"_wazuh_token", token, 3600)

	if err != nil {
		return "", fmt.Errorf("token存入Redis缓存失败：%v", err)
	}
	return token, nil
}

//token 获取token
func (this *Request) token() (string, error) {

	value, err := redis_cache.GetCache(server_api_username + "_wazuh_token")
	token := value.(string)
	if err == nil && len(token) != 0 {
		return token, nil
	}
	return this.updateToken()
}
func (this *Request) GetToken() (token string, err error) {
	return this.token()
}
func (this *Request) cookie() (cookie string, err error) {

	value, err := redis_cache.GetCache(this.UserName + "_wazuh_cookie")
	token := value.(string)
	if err == nil && len(token) != 0 {
		return token, nil
	}
	return this.updateCookie()
}
func (this *Request) GetCookie() (cookie string, err error) {

	value, err := redis_cache.GetCache(this.UserName + "_wazuh_cookie")
	cookie = value.(string)
	if err == nil && len(cookie) != 0 {
		return cookie, nil
	}
	return this.updateCookie()
}
func (this *Request) updateCookie() (cookie string, err error) {

	url := this.url + "/auth/login"

	client := resty.New().SetTimeout(time.Second * 60)
	if strings.HasPrefix(this.url, "https://") {
		client = client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	resp, err := client.SetDebug(false).R().
		SetHeader("Content-Type", "application/json").
		SetHeader("kbn-version", kbn_version).
		SetBody(map[string]string{
			"username": req.UserName,
			"password": req.Password,
		}).
		Post(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() == 200 {
		//获取cookie
		Cookies := resp.Cookies()
		if len(Cookies) > 0 {
			for _, v := range Cookies {
				if v.Name == "security_authentication" {
					_ = redis_cache.SetCache(this.UserName+"_wazuh_cookie", v.Value, 3600)
					return v.Value, nil
				}
			}
		}
	}
	return "", nil
}

//ToString 打印
func (this *Request) ToString() string {
	str := ""
	str += fmt.Sprintf("url : %v \n", this.url)
	str += fmt.Sprintf("Method : %v \n", this.Method)
	str += fmt.Sprintf("Params : %v \n", this.Params)
	str += fmt.Sprintf("Header : %v \n", this.Headers)
	return str
}

//Do 执行请求
func (this *Request) Do() (respBody []byte, err error) {
	var body io.Reader
	var updateTokenFlags bool

	this.Method = strings.ToUpper(this.Method)
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	//非get 参数设置在body中 以json形式传输
	//if this.Method != "GET" && len(this.Params) > 0 {
	//	buf, _ := json.Marshal(this.Params)
	//	body = bytes.NewReader(buf)
	//}
	req, err := http.NewRequest(this.Method, this.url+server_api_port+this.Path, body)

	if err != nil {
		return nil, err
	}
	//if this.Method == "GET" {
	q := req.URL.Query()
	for k, v := range this.Params {
		q.Add(k, fmt.Sprintf("%v", v))
	}
	req.URL.RawQuery = q.Encode()
	//}
	//todo:获取token
	token, err := this.token()
	if err != nil {
		return nil, fmt.Errorf("获取token失败：%v", err)
	}

	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}

doRequest:
	req.Header.Set("Authorization", authorization_template+token)
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	retBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &RetInfo{}
	_ = json.Unmarshal(retBuf, &ret)

	//验证是否token过期
	if len(ret.Title) != 0 && ret.Title == "Unauthorized" {
		if updateTokenFlags {
			return nil, fmt.Errorf("token失效")
		} else {
			token, err = this.updateToken()
			updateTokenFlags = true
			goto doRequest
		}
	}
	return retBuf, nil
}

//Do 执行请求
func (this *Request) Do2(param interface{}) (respBody []byte, err error) {
	var body io.Reader
	var updateTokenFlags bool

	this.Method = strings.ToUpper(this.Method)
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	//非get 参数设置在body中 以json形式传输
	if this.Method != "GET" && param != nil {
		buf, _ := json.Marshal(param)
		body = bytes.NewReader(buf)
		fmt.Println(string(buf))
	}
	req, err := http.NewRequest(this.Method, this.url+this.Path, body)
	if err != nil {
		return nil, err
	}

	if this.Method == "GET" {
		q := req.URL.Query()
		for k, v := range this.Params {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}
	//todo:获取token
	token, err := this.token()
	if err != nil {
		return nil, fmt.Errorf("获取token失败：%v", err)
	}
	cookie, err := this.GetCookie()
	if err != nil {
		return nil, fmt.Errorf("获取cookie失败：%v", err)
	}
	this.Headers["Content-Type"] = "application/json"
	this.Headers["kbn-version"] = kbn_version
	this.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
doRequest:
	req.Header.Set("Cookie", fmt.Sprintf("wz-user=wazuh; wz-api=default; wz-token=%s;security_authentication=%s", token, cookie))
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	retBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &RetInfo{}
	_ = json.Unmarshal(retBuf, &ret)
	//验证是否token过期
	if len(ret.Title) != 0 && ret.Title == "Unauthorized" {
		if updateTokenFlags {
			return nil, fmt.Errorf("token失效")
		} else {
			token, err = this.updateToken()
			if err != nil {
				return nil, fmt.Errorf("获取token失败：%v", err)
			}
			cookie, err = this.GetCookie()
			if err != nil {
				return nil, fmt.Errorf("获取cookie失败：%v", err)
			}
			updateTokenFlags = true
			goto doRequest
		}
	}
	return retBuf, nil
}

//DoAndParseResp 执行请求并解析
func (this *Request) DoAndParseResp() (ret *RetInfo, err error) {
	resp, err := this.Do()
	if err != nil {
		return nil, err
	}
	ret = &RetInfo{}
	err = json.Unmarshal(resp, &ret)
	if ret.Detail != "" {
		return ret, fmt.Errorf(ret.Detail)
	}
	return
}
