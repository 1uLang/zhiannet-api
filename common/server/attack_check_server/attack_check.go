package attack_check_server

import (
	"encoding/json"
	"fmt"
	awvs_request "github.com/1uLang/zhiannet-api/awvs/request"
	awvs_server "github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/1uLang/zhiannet-api/common/server/edge_server_server"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
)

type AttackCheckRequest struct{}

func (*AttackCheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("teaweb-入侵事件检测----------------------------------------------", err)
		}
	}()

	//ddos 入侵检测 将入侵ip 加入ddos黑名单
	//func() {
	//	err := ddos{}.AttackCheck()
	//	if err != nil {
	//		fmt.Println("ddos 攻击ip自动加入 黑名单错误：", err)
	//		return
	//	}
	//}()
	//hids 入侵检测 将入侵ip 加入ddos黑名单
	func() {

		info, err := server.GetHideInfo()
		if err != nil {
			fmt.Println("获取hids配置信息失败:", err)
			return
		}
		//初始化 hids 服务器地址
		err = server.SetUrl(info.Addr)
		if err != nil {
			fmt.Println("设置hids域名错误:", err)
			return
		}
		//初始化 hdis 系统管理员账号apikeys
		err = server.SetAPIKeys(&request.APIKeys{
			AppId:  info.AppId,
			Secret: info.Secret,
		})
		ips, err := hids{}.AttackCheck()
		if err != nil {
			fmt.Println("hids入侵检测失败:", err)
			return
		}
		fmt.Println("add ddos black ip : ", ips)
		err = ddos{}.AddBlackIP(ips)
		if err != nil {
			fmt.Println("hids入侵检测 自动添加ddos黑名单失败:", err)
			return
		}
	}()
	//hids 入侵检测 网页后门 加自动进行web漏洞扫描
	func() {
		ips, err := hids{}.WebShellAttackCheck()
		if err != nil {
			fmt.Println("hids 网页后门入侵检测失败:", err)
			return
		}

		info, err := awvs_server.GetWebScan()
		if err != nil {
			fmt.Printf("web漏洞扫描获取配置信息失败:%v\n", err)
			return
		}
		err = awvs_server.SetUrl(info.Addr)
		if err != nil {
			fmt.Printf("web漏洞扫描配置url失败:%v\n", err)
			return
		}
		err = awvs_server.SetAPIKeys(&awvs_request.APIKeys{XAuth: info.Key})
		if err != nil {
			fmt.Printf("web漏洞扫描配置api、key失败:%v\n", err)
			return
		}
		err = webscan{}.DoWebScan(ips)
		if err != nil {
			fmt.Printf("webscan 自动创建并扫描[%v]失败:%v\n", ips, err)
			return
		}
	}()
	//waf入侵检测 自动加入waf黑名单
	func() {
		err := waf{}.WAFAttackCheck()
		if err != nil {
			fmt.Printf("waf入侵检测失败:%v\n", err)
			return
		}
	}()
	func() {
		info, err := awvs_server.GetWebScan()
		if err != nil {
			fmt.Printf("web漏洞扫描获取配置信息失败:%v\n", err)
			return
		}
		err = awvs_server.SetUrl(info.Addr)
		if err != nil {
			fmt.Printf("web漏洞扫描配置url失败:%v\n", err)
			return
		}
		err = awvs_server.SetAPIKeys(&awvs_request.APIKeys{XAuth: info.Key})
		if err != nil {
			fmt.Printf("web漏洞扫描配置api、key失败:%v\n", err)
			return
		}
		servers, err := edge_server_server.GetServerList()
		if err != nil {
			fmt.Printf("waf服务列表失败:%v\n", err)
			return
		}
		checkAddress := []string{}
		urlHttpsConfig := map[string]int{}
		updateServerIdxs := map[int]bool{}

		for idx, v := range servers {
			urls := []struct {
				Name string `json:"name"`
			}{}
			_ = json.Unmarshal(v.ServerNames, &urls)
			for _, addr := range urls {
				urlHttpsConfig[addr.Name] = idx
				updateServerIdxs[idx] = false
				checkAddress = append(checkAddress, addr.Name)
			}
		}
		addres, err := webscan{}.WebScanCheckTLSVul(checkAddress)
		if err != nil {
			fmt.Printf("web TLS版本过低检测失败:%v\n", err)
			return
		}
		for _, v := range addres {
			idx := urlHttpsConfig[v]
			if !updateServerIdxs[idx] {
				//---- 更新 tls配置
				checkAndUpdateHttpsConfig(servers[idx].HttpsJSON)
				updateServerIdxs[idx] = true
				//通知更新
				if chanServer != nil {
					*chanServer <- servers[idx].ID
				}

			}
		}
	}()
}
