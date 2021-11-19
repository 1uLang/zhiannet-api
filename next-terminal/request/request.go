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

const (
	authorization_template = ""

	token_path = "/login"
)

type Request struct {
	Method  string
	url     string
	Path    string
	Params  map[string]interface{}
	Headers map[string]string
}
type RetInfo struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var req = Request{
	Method: "get",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

//NewRequest 创建一个请求头
func NewRequest(url string) (*Request, error) {

	return &Request{Headers: req.Headers, url: url}, nil
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

	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	//req.Header.Set("X-Auth-Token", "85a69e08-bf9b-4d05-bd31-841238a090c1e6cf1f0a-3392-465a-b89f-8c0c9542e88f2637fae9-5229-407b-ae77-67282415ea571abcc44e-37d3-438b-91a9-f366a3bb396c")
	req.Header.Set("X-Auth-Token", "0bff649f-6c51-4b29-ab03-3d1f298e66cc870cc9fc-c21b-4b38-ac01-2ef31aeb41edf04b4a84-785f-4a37-ad0e-397ba95b1b2f38aca310-a9c9-4f9f-89b2-39bd2ba4fa29")
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
	return
}
