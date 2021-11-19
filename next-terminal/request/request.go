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
	//req.Header.Set("X-Auth-Token", "3776e8f7-4cb5-48a4-8568-4c0f1cb4f5d35daacfeb-ee38-4267-9136-f41d15954c79b7d54370-0c67-458c-87e0-d790f813c683f63c883c-6a1c-46c0-8c78-6904f4238183")
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
