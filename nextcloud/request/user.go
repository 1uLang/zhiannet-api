package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

// CreateUser 创建用户
func CreateUser(token, user, passwd string) error {
	uRL := fmt.Sprintf("%s/%s", param.BASE_URL, param.CREATE_USER)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("POST", uRL, nil)
	if err != nil {
		return fmt.Errorf("创建用户接口请求失败：%w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("OCS-APIRequest", "true")
	reqQuery := req.URL.Query()
	reqQuery.Add("userid", user)
	reqQuery.Add("password", passwd)
	req.URL.RawQuery = reqQuery.Encode()
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("请求创建用户接口失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取创建用户响应失败：%w", err)
	}
	cuRsp := model.CreateUserResp{}
	err = xml.Unmarshal(body, &cuRsp)
	if err != nil {
		return fmt.Errorf("xml解析失败：%w", err)
	}

	// 说明用户已经被创建过，避免脏数据，这里不将这个错误抛出
	if cuRsp.Meta.Statuscode == 102 {
		return nil
	}

	if cuRsp.Meta.Status != "ok" || cuRsp.Meta.Statuscode != 100 {
		return errors.New(cuRsp.Meta.Message)
	}

	return nil
}

func CreateUserV2(token, user, pwd string) error {
	cUserURL := fmt.Sprintf("%s/%s", param.BASE_URL, param.CREATE_USER_V2)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	cur := model.CreateUserReq{
		Userid:      user,
		DisplayName: user,
		Password:    pwd,
		Quota:       "1 GB",
		Language:    "zh_CN",
	}
	rb, err := json.Marshal(cur)
	if err != nil {
		return fmt.Errorf("JSON编码失败：%w", err)
	}
	req, err := http.NewRequest("POST", cUserURL, bytes.NewReader(rb))
	req.Header.Set("Authorization", token)
	req.Header.Add("OCS-APIRequest", "true")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return fmt.Errorf("新建请求失败：%w", err)
	}
	rsp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败：%w", err)
	}
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败：%w", err)
	}
	cuRsp := model.CreateUserResp{}
	err = xml.Unmarshal(body, &cuRsp)
	if err != nil {
		return fmt.Errorf("xml解析失败：%w", err)
	}

	// 说明用户已经被创建过，避免脏数据，这里不将这个错误抛出
	if cuRsp.Meta.Statuscode == 102 {
		return nil
	}

	if cuRsp.Meta.Status != "ok" || cuRsp.Meta.Statuscode != 200 {
		return errors.New(cuRsp.Meta.Message)
	}

	return nil
}
