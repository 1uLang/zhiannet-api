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
	"time"
)

type request struct {
	Method  string
	Url     string
	Params  map[string]interface{}
	Headers map[string]string
	Secret  string
}

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
	req.Url = url
	return nil
}
func InitRequestAPIKeys(api *APIKeys) error {

	ok, err := api.Check()
	if err != nil {
		return err
	}
	if ok {
		req.Headers["appId"] = api.AppId
		req.Secret = api.Secret
		return nil
	} else {
		return fmt.Errorf("参数错误")
	}
}

func NewRequest() (*request, error) {
	if req.Url == "" {
		return nil, fmt.Errorf("未配置主机防护 服务器地址")
	}
	if _, isExist := req.Headers["appId"]; !isExist {
		return nil, fmt.Errorf("未配置appid")
	}
	return &request{Headers: req.Headers, Url: req.Url, Secret: req.Secret}, nil
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
		value := fmt.Sprintf("[\"%v\"]", v)
		if !isExist {
			value = "\"" + this.Headers[k] + "\""
		}
		stringSignTemp += fmt.Sprintf("%s=%v&", k, value)
	}
	//加上私钥key
	stringSignTemp += "key=" + this.Secret
	fmt.Println(stringSignTemp)
	//rsa加密
	h := hmac.New(sha256.New, []byte(this.Secret))
	h.Write([]byte(stringSignTemp))

	return hex.EncodeToString(h.Sum(nil))
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
	if strings.ToUpper(this.Method) != "GET" {
		buf, _ := json.Marshal(this.Params)
		body = bytes.NewReader(buf)
	}
	req, err := http.NewRequest(this.Method, this.Url, body)
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
	fmt.Println("http request :" + req.URL.String())
	//毫秒时间戳
	this.Headers["time"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	//签名
	this.Headers["sign"] = this.sign()
	for k, v := range this.Headers {
		req.Header.Add(k, v)
	}
	fmt.Println("request info : \n", this.ToString())
	//请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
