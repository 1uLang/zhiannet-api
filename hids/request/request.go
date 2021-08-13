package request

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type request struct {
	Method  string
	url     string
	Path    string
	Params  map[string]interface{}
	Headers map[string]string
	Secret  string
}

var req_mutex sync.RWMutex
var req = request{
	Method: "get",
	Headers: map[string]string{
		"signMethod":  _const.Sign_method,
		"signVersion": _const.Sign_version,
		"version":     _const.Version,
	},
}

type APIKeys struct {
	AppId  string `json:"appid"`
	Secret string `json:"secret"`
}

func (this *APIKeys) Check() (bool, error) {

	if this.AppId == "" {
		return false, fmt.Errorf("请输入appid")
	}
	if this.Secret == "" {
		return false, fmt.Errorf("请输入secret")
	}
	return true, nil
}
func InitServerUrl(url string) error {
	req_mutex.Lock()
	defer req_mutex.Unlock()
	req.url = url
	return nil
}
func InitRequestAPIKeys(api *APIKeys) error {

	ok, err := api.Check()
	if err != nil {
		return err
	}
	if ok {
		req_mutex.Lock()
		req.Headers["appId"] = api.AppId
		req.Secret = api.Secret
		req_mutex.Unlock()
		return nil
	} else {
		return fmt.Errorf("参数错误")
	}
}

func NewRequest() (*request, error) {
	req_mutex.RLock()
	if req.url == "" {
		req_mutex.RUnlock()
		return nil, fmt.Errorf("未配置主机防护 服务器地址")
	}
	if _, isExist := req.Headers["appId"]; !isExist {
		req_mutex.RUnlock()
		return nil, fmt.Errorf("未配置appId")
	}
	req_mutex.RUnlock()
	//删除签名
	req_mutex.Lock()
	delete(req.Headers, "sign")
	req_mutex.Unlock()

	req_mutex.RLock()
	defer req_mutex.RUnlock()
	return &request{Headers: req.Headers, url: req.url, Secret: req.Secret}, nil
}
func NewRequest2() (*request, error) {
	if req.url == "" {
		return nil, fmt.Errorf("未配置主机防护 服务器地址")
	}
	if _, isExist := req.Headers["appId"]; !isExist {
		return nil, fmt.Errorf("未配置appId")
	}
	r := &request{Headers: req.Headers, url: req.url, Secret: req.Secret}

	r.Headers["version"] = _const.Version2
	r.Headers["signVersion"] = _const.Sign_version2
	return r, nil
}

//sign  rsa HmacSHA256 签名函数
//参数：
//	params 请求参数
//	secret 密钥
func (this *request) sign() string {

	keyList := []string{}
	//排序
	for k := range this.Params {
		keyList = append(keyList, k)
	}
	for k := range this.Headers {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	stringSignTemp := ""
	for _, k := range keyList {
		v, isExist := this.Params[k]
		value := ""
		if this.Method == "POST" {
			//if k == "itemIds" || k == "riskIds" {
			tmpbuf, _ := json.Marshal(v)
			value = fmt.Sprintf("%s", string(tmpbuf))
			//} else {
			//	value = fmt.Sprintf("[\"%v\"]", v)
			//}
		} else {
			value = fmt.Sprintf("[\"%v\"]", v)
		}

		if !isExist {
			value = "\"" + this.Headers[k] + "\""
		}
		stringSignTemp += fmt.Sprintf("%s=%v&", k, value)
	}
	//加上私钥key
	stringSignTemp += "key=" + this.Secret
	//rsa加密
	h := hmac.New(sha256.New, []byte(this.Secret))
	h.Write([]byte(stringSignTemp))

	return hex.EncodeToString(h.Sum(nil))
}

//sign2  rsa HmacSHA256 签名函数
//参数：
//	params 请求参数
//	secret 密钥
func (this *request) sign2() string {

	keyList := []string{}
	addData := false
	//排序
	for k := range this.Params {
		if !addData && k == "data" {
			for key := range this.Params["data"].(map[string]interface{}) {
				keyList = append(keyList, key)
			}
			addData = true
			continue
		}
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	stringSignTemp := ""
	for _, k := range keyList {
		value, isExist := this.Params[k]
		if !isExist {
			value = this.Params["data"].(map[string]interface{})[k]
		}
		stringSignTemp += fmt.Sprintf("%s=%v&", k, value)
	}
	//加上私钥key
	stringSignTemp += "key=" + this.Secret
	//rsa加密
	h := hmac.New(sha256.New, []byte(this.Secret))
	h.Write([]byte(stringSignTemp))

	return hex.EncodeToString(h.Sum(nil))
}
func (this *request) ToString() string {
	str := ""
	str += fmt.Sprintf("url : %v \n", this.url+this.Path)
	str += fmt.Sprintf("Method : %v \n", this.Method)
	str += fmt.Sprintf("Params : %v \n", this.Params)
	str += fmt.Sprintf("Header : %v \n", this.Headers)
	return str
}
func (this *request) Do2() (respBody []byte, err error) {

	this.Method = strings.ToUpper(this.Method)
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	//将header参数放在param中
	//appId
	this.Params["appid"] = this.Headers["appId"]
	this.Params["version"] = this.Headers["version"]
	this.Params["signVersion"] = this.Headers["signVersion"]
	this.Params["signMethod"] = this.Headers["signMethod"]
	//毫秒时间戳
	this.Params["time"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	//签名
	this.Params["sign"] = this.sign2()

	var body io.Reader
	//非get 参数设置在body中 以json形式传输
	if strings.ToUpper(this.Method) != "GET" {
		buf, _ := json.Marshal(this.Params)
		body = bytes.NewReader(buf)
	}
	req, err := http.NewRequest(this.Method, this.url+this.Path, body)
	if err != nil {
		return nil, err
	}

	//非get 设置content-type
	if this.Method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	} else {
		q := req.URL.Query()
		for k, v := range this.Params {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}
	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
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
	if strings.ToUpper(this.Method) != "GET" {
		buf, _ := json.Marshal(this.Params)
		body = bytes.NewReader(buf)
	}
	req, err := http.NewRequest(this.Method, this.url+this.Path, body)
	if err != nil {
		return nil, err
	}

	//非get 设置content-type
	if this.Method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	} else {
		q := req.URL.Query()
		for k, v := range this.Params {
			q.Add(k, fmt.Sprintf("%v", v))
		}
		req.URL.RawQuery = q.Encode()
	}
	//毫秒时间戳
	this.Headers["time"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	//签名
	this.Headers["sign"] = this.sign()
	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	//请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
