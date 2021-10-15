package server

import (
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"testing"
)

func TestInfo(t *testing.T) {

	err := redis_cache.InitTestClient()
	if err != nil {
		panic(err)
	}
	err = request.InitServerUrl("https://156.240.95.168")
	if err != nil {
		t.Fatal(err)
	}
	err = request.InitToken("wazuh", "wazuh")
	if err != nil {
		t.Fatal(err)
	}
	req, err := request.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	info, _ := Info(req)
	fmt.Println(info.ApiVersion)
}
