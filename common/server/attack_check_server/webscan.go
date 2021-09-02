package attack_check_server

import (
	"crypto/tls"
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	scans_server "github.com/1uLang/zhiannet-api/awvs/server/scans"
	targets_server "github.com/1uLang/zhiannet-api/awvs/server/targets"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"net/http"
	"time"
)

var web_server_port = 8088

type webscan struct{}

//判断存在ssl不安全版本的地址
func (webscan) WebScanCheckTLSVul(address []string) (adds []string, err error) {
	for _, addr := range address {
		list, err := targets_server.Search(&targets.ListReq{Limit: 100, Query: "text_search:*" + addr + ";"})
		if err != nil {
			return nil, err
		}
		vulFflag := false
		for _, item := range list {
			target := item.(map[string]interface{})
			scs, err := scans_server.Search(&scans.ListReq{Limit: 100, Query: "target_id:" + target["target_id"].(string) + ";"})
			if err != nil {
				return nil, err
			}
			for _, sitem := range scs {
				scan := sitem.(map[string]interface{})
				vuls, err := scans_server.VulnerabilitiesList(&scans.VulnerabilitiesListReq{scan["scan_id"].(string),
					scan["current_session"].(map[string]interface{})["scan_session_id"].(string)})
				if err != nil {
					return nil, err
				}
				for _, vitem := range vuls {
					vul := vitem.(map[string]interface{})
					if vul["vt_name"].(string) == "TLS 1.1 已启用" {
						vulFflag = true
						break
					}
				}
				if vulFflag {
					break
				}
			}
			if vulFflag {
				break
			}
		}
		//存在tls漏洞
		if vulFflag {
			adds = append(adds, addr)
		}
	}
	return adds, nil
}

func SetWebServerPort(port int) {
	web_server_port = port
}

func (webscan) checkWebServer(ip string) bool {

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%v:%v", ip, web_server_port), nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	return true
}

//对指定地址进行web漏洞扫描
func (this webscan) DoWebScan(ips []string) error {

	for _, ip := range ips {
		if this.checkWebServer(ip) { //判断是否存在web服务，有则进行web scan
			//1、找到该agent 所属用户
			users, admins, err := getAgentIpUsers(ip)
			if err != nil {
				return err
			}
			for _, user := range users {
				//该用户下创建目标并扫描
				_ = this.createAndScanWebServer(ip, user, 0)
			}
			for _, admin := range admins {
				//该用户下创建目标并扫描
				_ = this.createAndScanWebServer(ip, 0, admin)
			}

		}
	}
	return nil
}
func (webscan) createAndScanWebServer(ip string, user, admin uint64) error {

	addr := fmt.Sprintf("https://%v:%v", ip, web_server_port)
	if user == 0 && admin == 0 {
		return nil
	}
	target_id, ok, err := targets.CheckAddr(&targets.CheckAddrReq{Addr: addr, UserId: user, AdminUserId: admin})
	if err != nil {
		return err
	}
	if !ok { //不存在则创建
		target_id, err = targets_server.Add(&targets.AddReq{
			UserId: user, AdminUserId: admin,
			Address: addr,
		})
		if err != nil {
			return err
		}
	}

	if target_id != "" {
		err = scans_server.Add(&scans.AddReq{TargetId: target_id, ProfileId: "11111111-1111-1111-1111-111111111111"})
		if err != nil {
			return err
		}
	}

	edge_messages.Add(&edge_messages.Edgemessages{
		Level:     "error",
		Subject:   "入侵检测",
		Userid:    uint(user),
		Adminid:   uint(admin),
		Body:      fmt.Sprintf("主机防护漏洞检测网页后门[%s],已自动创建并对其进行web漏洞扫描", addr),
		Type:      "IntrusionDetection",
		Params:    "{}",
		Createdat: uint64(time.Now().Unix()),
		Day:       time.Now().Format("20060102"),
		Hash:      "",
		Role:      "admin",
	})
	return nil
}
