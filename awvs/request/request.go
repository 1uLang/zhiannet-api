package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type APIKeys struct {
	XAuth string `json:"x_auth"`
}

func (this *APIKeys) Check() (bool, error) {

	if this.XAuth == "" {
		return false, fmt.Errorf("请输入X-Auth")
	}
	return true, nil
}

type request struct {
	Method  string
	Url     string
	Params  map[string]interface{}
	Headers map[string]string
}

var req = request{
	Method: "get",
	Headers: map[string]string{
		"Content-Type": "application/json;charset=utf8",
		"Accept":       "application/json",
	},
}

//InitServerUrl 初始化awvs服务器地址
func InitServerUrl(url string) error {
	req.Url = url
	return nil
}

//InitRequestXAuth 初始化X-Auth
func InitRequestXAuth(api *APIKeys) error {
	req.Headers["X-Auth"] = api.XAuth
	return nil
}

//NewRequest 创建一个请求头
func NewRequest() (*request, error) {
	if req.Url == "" {
		return nil, fmt.Errorf("未配置awvs 服务器地址")
	}
	if _, isExist := req.Headers["X-Auth"]; !isExist {
		return nil, fmt.Errorf("未配置X-Auth")
	}

	return &request{Headers: req.Headers, Url: req.Url}, nil
}

//ToString 打印
func (this *request) ToString() string {
	str := ""
	str += fmt.Sprintf("Url : %v \n", this.Url)
	str += fmt.Sprintf("Method : %v \n", this.Method)
	str += fmt.Sprintf("Params : %v \n", this.Params)
	str += fmt.Sprintf("Header : %v \n", this.Headers)
	return str
}

//Do 执行请求
func (this *request) Do() (respBody []byte, err error) {

	this.Method = strings.ToUpper(this.Method)
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	var body io.Reader
	//非get 参数设置在body中 以json形式传输
	if strings.ToUpper(this.Method) != "GET" && len(this.Params) > 0 {
		buf, _ := json.Marshal(this.Params)
		body = bytes.NewReader(buf)
	}
	req, err := http.NewRequest(this.Method, this.Url, body)
	if err != nil {
		return nil, err
	}

	if strings.ToUpper(this.Method) == "GET" {
		q := req.URL.Query()
		for k, v := range this.Params {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}
	fmt.Println("http audit_db :" + req.URL.String())

	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	fmt.Println("audit_db dashboard : \n", this.ToString())
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
