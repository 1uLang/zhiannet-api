package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	authorization_template = ""

	token_path = "/login"
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
type RetInfo struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var req = Request{
	Method: "get",
	Headers: map[string]string{
		"X-Auth-Token": "",
		"Content-Type": "application/json",
	},
}

//InitServerUrl 初始化next-terminal服务器地址
func InitServerUrl(url string) error {
	req.url = url
	return nil
}

//InitToken 初始化 token 所需的当前用户的账号密码
func InitToken(username, password string) error {

	//将改用户的账号密码存入程序缓存中。重启时自动清空
	req.UserName = username
	req.Password = password
	_, err := req.token()
	if err != nil {
		return fmt.Errorf("初始化token失败：%v", err)
	}
	return nil
}

//NewRequest 创建一个请求头
func NewRequest() (*Request, error) {
	if req.url == "" {
		return nil, fmt.Errorf("未配置堡垒机 服务器地址")
	}

	if req.UserName == "" {
		return nil, fmt.Errorf("未配置堡垒机系统管理员账号")
	}
	if req.Password == "" {
		return nil, fmt.Errorf("未配置堡垒机系统管理员密码")
	}

	return &Request{Headers: req.Headers, url: req.url, UserName: req.UserName, Password: req.Password}, nil
}

//updateToken 获取更新 token
func (this *Request) updateToken() (string, error) {
	params := map[string]interface{}{
		"username": this.UserName,
		"password": this.Password,
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	buf, _ := json.Marshal(params)
	body := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", this.url+token_path, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	userInfo := map[string]interface{}{}
	_ = json.Unmarshal(b, &userInfo)
	token := userInfo["data"].(string)
	err = redis_cache.SetCache(this.UserName+"_next-terminal_token", token, 3600)

	if err != nil {
		return "", fmt.Errorf("token存入Redis缓存失败：%v", err)
	}
	return token, nil
}

//token 获取token
func (this *Request) token() (string, error) {

	value, err := redis_cache.GetCache(this.UserName + "_next-terminal_token")
	token := value.(string)
	if err == nil && len(token) != 0 {
		return token, nil
	}
	return this.updateToken()
}
func (this *Request) GetToken() (token string, err error) {
	return this.token()
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

	this.Method = strings.ToUpper(this.Method)
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	//非get 参数设置在body中 以json形式传输
	if this.Method != "GET" && len(this.Params) > 0 {
		buf, _ := json.Marshal(this.Params)
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

	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("X-Auth-Token", authorization_template+token)
	//请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	retBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	retInfo := map[string]interface{}{}
	_ = json.Unmarshal(retBuf, &retInfo)
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
	fmt.Println(ret)
	return
}
