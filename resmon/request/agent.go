package request

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	param "github.com/1uLang/zhiannet-api/resmon/const"
	"github.com/1uLang/zhiannet-api/resmon/model"
	"github.com/1uLang/zhiannet-api/resmon/server"
)

func generatepost(reqURL, name, host, agentID string, on bool) ([]byte, error) {
	// 获取teaweb节点信息
	server.GetNodeInfo()
	if param.BASE_URL == "" {
		return nil, errors.New("该节点暂未添加，请添加后重试")
	}
	url := fmt.Sprintf("%s/%s", param.BASE_URL, reqURL)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败：%w", err)
	}
	reqQuery := req.URL.Query()
	reqQuery.Add("TeaKey", param.TEA_KEY)
	reqQuery.Add("name", name)
	reqQuery.Add("host", host)
	reqQuery.Add("on", strconv.FormatBool(on))
	if reqURL == param.UPDATE_AGENT {
		reqQuery.Add("agentId", agentID)
	}
	req.URL.RawQuery = reqQuery.Encode()

	rsp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("获取请求响应失败：%w", err)
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取请求响应体失败：%w", err)
	}

	return body, nil
}

// AddAgent 添加Agent节点
func AddAgent(name, host string, on bool, aid uint8) error {
	body, err := generatepost(param.ADD_AGENT, name, host, "", on)
	if err != nil {
		return err
	}

	br := &model.AddAgentResp{}
	err = json.Unmarshal(body, br)
	if err != nil {
		return fmt.Errorf("json解析失败：%w", err)
	}

	if br.Code != 200 {
		return fmt.Errorf("添加监控主机失败：%s", br.Message)
	}

	if err = server.AddResmon(br.Data.AgentID, aid); err != nil {
		return err
	}

	return nil
}

// UpdateAgent 更新agent信息
func UpdateAgent(name, host, agentID string, on bool,aid uint8) error {
	body, err := generatepost(param.UPDATE_AGENT, name, host, agentID, on)
	if err != nil {
		return err
	}

	br := &model.BaseResp{}
	err = json.Unmarshal(body, br)
	if err != nil {
		return  fmt.Errorf("json解析失败：%w", err)
	}

	if br.Code != 200 {
		return fmt.Errorf("修改监控主机失败：%s", br.Message)
	}

	if err := server.AddResmon(agentID,aid);err != nil {
		return err
	}

	return nil
}

// DeleteAgent 删除代理主机
func DeleteAgent(agentID string) error {
	// 获取teaweb节点信息
	server.GetNodeInfo()
	if param.BASE_URL == "" {
		return errors.New("该节点暂未添加，请添加后重试")
	}

	body, err := createReq(param.DELETE_AGENT, agentID)
	if err != nil {
		return err
	}

	br := &model.BaseResp{}
	json.Unmarshal(body, br)

	if br.Code != 200 && br.Code != 0 {
		return fmt.Errorf("删除监控主机失败：%s", br.Message)
	}

	if err = server.DeleteResmon(agentID); err != nil {
		return err
	}

	return nil
}
