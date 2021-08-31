package attack_check_server

import (
	"encoding/json"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/1uLang/zhiannet-api/common/server/edge_server_server"
	"testing"
)

func TestWebScanCheckTLSVul(t *testing.T) {
	InitTestDB()
	info, err := server.GetWebScan()
	if err != nil {
		t.Fatal(err)
	}
	err = server.SetUrl(info.Addr)
	if err != nil {
		t.Fatal(err)
	}
	err = server.SetAPIKeys(&request.APIKeys{XAuth: info.Key})
	if err != nil {
		t.Fatal(err)
	}
	servers, err := edge_server_server.GetServerList()
	if err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}
	for _, v := range addres {
		idx := urlHttpsConfig[v]
		if !updateServerIdxs[idx] {
			//---- 更新 tls配置
			checkAndUpdateHttpsConfig(servers[idx].HttpsJSON)
			updateServerIdxs[idx] = true
		}
	}
}

func TestDoWebScan(t *testing.T) {
	InitTestDB()
	info, err := server.GetWebScan()
	if err != nil {
		t.Fatal(err)
	}
	err = server.SetUrl(info.Addr)
	if err != nil {
		t.Fatal(err)
	}
	err = server.SetAPIKeys(&request.APIKeys{XAuth: info.Key})
	if err != nil {
		t.Fatal(err)
	}
	SetWebServerPort(443)
	err = webscan{}.DoWebScan([]string{"182.150.0.83"})
	if err != nil {
		t.Fatal(err)
	}
}
