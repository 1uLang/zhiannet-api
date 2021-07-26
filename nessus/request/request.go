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
	Access string `json:"access"`
	Secret string `json:"secret"`
}

func (this *APIKeys) Check() (bool, error) {

	if this.Access == "" {
		return false, fmt.Errorf("请输入accessKey")
	}
	if this.Secret == "" {
		return false, fmt.Errorf("请输入secretKey")
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
		"Accept":       "application/json",
		"Content-Type": "application/json",
	},
}

func InitServerUrl(url string) error {
	req.Url = url
	return nil
}
func InitRequestAPIKeys(api *APIKeys) error {

	ok, err := api.Check()
	if err != nil {
		return err
	}
	if ok {
		req.Headers["x-apikeys"] = "accessKey=" + api.Access + ";secretKey=" + api.Secret
		return nil
	} else {
		return fmt.Errorf("参数错误")
	}
}
func NewRequest() (*request, error) {
	if req.Url == "" {
		return nil, fmt.Errorf("未配置nessus 服务器地址")
	}
	if _, isExist := req.Headers["x-apikeys"]; !isExist {
		return nil, fmt.Errorf("未配置APIKey")
	}

	return &request{Headers: req.Headers, Url: req.Url}, nil
}
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
		Timeout:   5 * time.Second,
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
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
