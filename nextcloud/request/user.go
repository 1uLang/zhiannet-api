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
	"strconv"

	cm_model "github.com/1uLang/zhiannet-api/common/model"
	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

// CreateUser 创建用户
func CreateUser(token, user, passwd string) error {
	getNCInfo()
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
	getNCInfo()
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

// DeleteNCUser 删除用户
func DeleteNCUser(user string) error {
	getNCInfo()
	token := GetAdminToken()
	uRL := fmt.Sprintf("%s/"+param.DELETE_USER, param.BASE_URL, user)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("DELETE", uRL, nil)
	if err != nil {
		return fmt.Errorf("删除户接口请求失败：%w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("OCS-APIRequest", "true")
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

	//if cuRsp.Meta.Status != "ok" || cuRsp.Meta.Statuscode != 200 {
	//	return errors.New(cuRsp.Meta.Message)
	//}

	return nil
}

// DeleteUser 删除用户
func DeleteUser(uid, kid int64) error {
	nct := model.NextCloudToken{}
	cm_model.MysqlConn.First(&nct, "uid = ? AND kind = ?", uid, kid)
	if nct.ID == 0 {
		return nil
	}

	err := DeleteNCUser(nct.User)
	if err != nil {
		return err
	}
	ddb := cm_model.MysqlConn.Delete(&model.NextCloudToken{}, nct.ID)
	if ddb.Error != nil {
		return fmt.Errorf("删除用户失败：%w", ddb.Error)
	}

	return nil
}

// GetNCUserInfo 获取用户信息
func GetNCUserInfo(token, user string) (quota, used, percent string) {
	// getNCInfo()
	percent = "1%"
	uRL := fmt.Sprintf("%s/"+param.USER_INFO, param.BASE_URL, user)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", uRL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", token)
	req.Header.Add("OCS-APIRequest", "true")
	rsp, err := cli.Do(req)
	if err != nil {
		return
	}
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	userInfo := model.NCUserInfo{}
	err = xml.Unmarshal(body, &userInfo)
	if err != nil {
		return
	}

	quota = userInfo.Data.Quota.Quota
	used = userInfo.Data.Quota.Used
	q, _ := strconv.ParseInt(quota, 10, 64)
	u, _ := strconv.ParseInt(used, 10, 64)
	if q <= 0 {
		quota = "无限"
	} else {
		quota = FormatBytes(quota)
		percent = fmt.Sprintf("%d", u/q) + "%"
	}
	used = FormatBytes(used)

	return
}

// UpdateUserPassword 更新用户密码
func UpdateUserPassword(newPassword, un string) error {
	getNCInfo()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	cli := &http.Client{
		Transport: tr,
	}

	cupURL := fmt.Sprintf("%s/"+param.UPDATE_USERPWD, param.BASE_URL, un)
	cb := model.ChangeUPwdReq{
		Key:   "password",
		Value: newPassword,
	}
	cbb, err := json.Marshal(cb)
	if err != nil {
		return fmt.Errorf("密码更新失败：%w", err)
	}
	req, err := http.NewRequest("PUT", cupURL, bytes.NewReader(cbb))
	if err != nil {
		return fmt.Errorf("密码更新失败：%w", err)
	}
	req.Header.Set("Authorization", GetAdminToken())
	req.Header.Add("OCS-APIREQUEST", "true")
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("密码更新失败：%w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("密码更新失败：%w", err)
	}
	cpr := model.CreateUserResp{}
	err = xml.Unmarshal(respBody, &cpr)
	if err != nil {
		return fmt.Errorf("密码更新失败：%w", err)
	}

	if cpr.Meta.Statuscode != 200 {
		return errors.New("密码更新失败：" + cpr.Meta.Message)
	}

	return nil
}
