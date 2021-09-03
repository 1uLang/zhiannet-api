package request

import (
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"testing"
)

func init() {
	err := redis_cache.InitTestClient()
	if err != nil {
		panic(err)
	}
}

func TestRequest_GetToken(t *testing.T) {

	err := InitServerUrl("https://156.240.95.168:55000")
	if err != nil {
		t.Fatal(err)
	}
	err = InitToken("wazuh", "wazuh")
	if err != nil {
		t.Fatal(err)
	}
	req, err := NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	token, err := req.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
