package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	LoginRes struct {
		Redirect `xml:"redirect"`
		Failure  `xml:"failure"`
	}
	Redirect struct {
		Page string `xml:"page"`
	}
	Failure struct {
		Info   string `xml:"info"`
		Url    string `xml:"url"`
		Params string `xml:"params"`
	}
)

//login 返回cookie
func Login(keys APIKeys) (cookie string, err error) {
	req := request{}
	req.Method = "POST"
	req.Params = map[string]interface{}{
		"param_type":     "login",
		"param_username": "cdadmin",
		"param_password": "A16pBIzVJOwHSC%23Q",
	}
	req.Headers = map[string]string{
		"Content-Type": "multipart/form-data; boundary=<calculated when request is sent>",
	}
	req.Url = _const.DDOS_HOST + _const.DDOS_LOGIN_URL
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	var body io.Reader
	//非get 参数设置在body中 以json形式传输
	buf, _ := json.Marshal(req.Params)
	body = bytes.NewReader(buf)
	reqs, err := http.NewRequest(req.Method, req.Url, body)
	//req.AddCookie()
	if err != nil {
		return "", err
	}
	fmt.Println(string(buf))
	//请求
	resp, err := client.Do(reqs)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	if err != nil {
		return "", err
	}
	var data = LoginRes{}
	err = xml.Unmarshal(res, &data)
	if err != nil {
		return "", err
	}
	//登陆失败原因
	if data.Failure.Info != "" {
		return "", fmt.Errorf(data.Failure.Info)
	}
	if len(resp.Cookies()) > 0 {
		cook := resp.Cookies()[0]
		cookie = cook.Value
	}
	return cookie, nil
}
