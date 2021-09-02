package attack_check_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/bwlist"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/1uLang/zhiannet-api/hids/request"
	hids_server "github.com/1uLang/zhiannet-api/hids/server"
	hids_agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	"gopkg.in/fatih/set.v0"
	"strings"
)
import hids_risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"

type hids struct{}

//hids 主机防护入侵检测ip

func (hids) AttackCheck() (ips []string, err error) {

	info, err := hids_server.GetHideInfo()
	if err != nil {
		return nil, err
	}
	err = hids_server.SetUrl(info.Addr)
	if err != nil {
		return nil, err
	}
	err = hids_server.SetAPIKeys(&request.APIKeys{info.AppId, info.Secret})
	if err != nil {
		return nil, err
	}

	resp, err := hids_agent_server.ListAll()
	if err != nil {
		return nil, err
	}
	setMap := set.New(set.NonThreadSafe)
	for _, item := range resp.List {
		macCode := item["macCode"].(string)

		//查询改macCode的登录异常列表
		abl, err := hids_risk_server.AbnormalLoginDetailList(&risk.DetailReq{
			MacCode: macCode,
		})
		if err != nil {
			return nil, err
		}
		//查询该异常是否包含ip 有加入到ddos黑名单
		fmt.Println(macCode, abl.ServerAbnormalLoginInfoList)
		for _, abnormal := range abl.ServerAbnormalLoginInfoList {
			loginInfo := abnormal["loginInfo"].(string)
			fmt.Println(loginInfo)
			if idx := strings.Index(loginInfo, "IP："); idx > 0 {
				ip := strings.TrimPrefix(loginInfo[idx:], "IP：")
				setMap.Add(ip)
			}
		}
	}
	return set.StringSlice(setMap), nil
}

//hids 主机防火入侵检测 网页后门
//自动对其进行web漏洞扫描
func (hids) WebShellAttackCheck() (ips []string, err error) {
	info, err := hids_server.GetHideInfo()
	if err != nil {
		return nil, err
	}
	err = hids_server.SetUrl(info.Addr)
	if err != nil {
		return nil, err
	}
	err = hids_server.SetAPIKeys(&request.APIKeys{info.AppId, info.Secret})
	if err != nil {
		return nil, err
	}

	list, err := risk.WebShellAllList()
	if err != nil {
		return nil, err
	}
	for _, v := range list.WebshellCountInfoList {
		ips = append(ips, v["serverIp"].(string))
	}
	return ips, err
}

//获取改agent ip 所属用户列表
func getAgentIpUsers(ip string) (users []uint64, admins []uint64, err error) {
	return hids_agent_server.GetUserListByAgentIP(ip)
}
func addBlackList(ips []string) error {
	fmt.Println("add hids black ip : ", ips)
	for _, ip := range ips {
		users, admins, err := getAgentIpUsers(ip)
		if err != nil {
			return err
		}
		for _, user := range users {
			_ = bwlist.AddBWList(&bwlist.HIDSBWList{Black: true, IP: ip, UserId: user})
		}
		for _, admin := range admins {
			_ = bwlist.AddBWList(&bwlist.HIDSBWList{Black: true, IP: ip, AdminUserId: admin})
		}
	}
	return nil
}
