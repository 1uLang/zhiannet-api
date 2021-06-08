package request

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

var Cookie string = "%2D%88%D9%3A%19%87%9C%28%8A%AF%EA%F8v%9CB%95%87%9E%8A3%98%EF%ED%F0%8C%075KAPfR%8B%82%85%D0%CC%AC4%00%96%DB%AE%88f%3E%7C%D6v1%C5%2D%A1%BE%7BQ%E4%B4u%D9%F9%E5%EBo%F9Qp%133l%C6%93%97%99%F2%DDdI%5Bx%AC%89%F6%05ha%ADw%16%10%F873%ABq%E4%F3lT%7D%E7%7F9%11IxU%BA%21%B7%2C%DC"

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

//登陆获取cookie
func Login() (string, error) {
	var err error
	// https://182.131.30.171:28443/cgi-bin/login.cgi
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetHeader("Content-Type", "multipart/form-data; boundary=<calculated when request is sent>").
		SetFormData(map[string]string{
			"param_type":     "login",
			"param_username": _const.USERNAME,
			"param_password": _const.PASSWORD,
		}).
		Post(_const.DDOS_HOST + _const.DDOS_LOGIN_URL)
	if err != nil {
		logrus.Error(err)
		return Cookie, err
	}

	//logrus.Info(string(resp.Body()))
	var data = LoginRes{}
	err = xml.Unmarshal(resp.Body(), &data)
	if err != nil {
		logrus.Error(err)
		return Cookie, err
	}
	//logrus.Info(data)
	if data.Failure.Info != "" {
		logrus.Debug(data.Failure)
		err = fmt.Errorf(data.Failure.Info)
		return Cookie, err
	}
	if len(resp.Cookies()) > 0 {
		cook := resp.Cookies()[0]
		Cookie = cook.Value
	}
	fmt.Println(Cookie)
	return Cookie, err
	//logrus.Info( err)
}
