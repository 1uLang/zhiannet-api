package agents

import (
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"testing"
)

func init() {
	err := redis_cache.InitTestClient()
	if err != nil {
		panic(err)
	}
}

func TestStatistics(t *testing.T) {

	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := Statistics(req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)

}
func TestList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := List(req, &ListReq{IP: "192.168.137.6"})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.AffectedItems) > 0 {
		err = Delete(req, []string{list.AffectedItems[0].ID})
		if err != nil {
			t.Fatal(err)
		}
	}
}
func TestScan(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	err = Scan(req, []string{"30889"})
	{
		t.Fatal(err)
	}
}
func TestSCAList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := SCAList(req, "003")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestSCADetailsList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := SCADetailsList(req, "003", "cis_centos7_linux")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestSysCheckList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := SysCheckList(req, "003")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestCiscatList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := CiscatList(req, "003")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestVulnerabilityList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168:55000")
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
	list, err := VulnerabilityList(req, "003")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}
