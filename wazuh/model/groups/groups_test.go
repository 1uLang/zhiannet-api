package groups

import (
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"testing"
)

func TestCreate(t *testing.T) {

	err := redis_cache.InitTestClient()
	if err != nil {
		panic(err)
	}
	err = request.InitServerUrl("https://156.240.95.168:55000")
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
	for i := 0; i < 100; i++ {
		err = Create(req, fmt.Sprintf("user_%v", i))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestList(t *testing.T) {

	err := redis_cache.InitTestClient()
	if err != nil {
		panic(err)
	}
	err = request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := List(req)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range list.AffectedItems {
		if v.Count == 0 {
			fmt.Println("delete : ", v.Name)
			err = Delete(req, []string{v.Name})
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
