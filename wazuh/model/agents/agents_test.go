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

	err := request.InitServerUrl("https://156.240.95.168")
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
	err := request.InitServerUrl("https://156.240.95.168")
	if err != nil {
		t.Fatal(err)
	}
	err = request.InitToken("wazuh", "AgI_kwQ2GQ8v354EQtd6pSpT7bDjdaNJ")
	if err != nil {
		t.Fatal(err)
	}
	req, err := request.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	list, err := List(req, &ListReq{})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.AffectedItems) > 0 {
		//err = Delete(req, []string{list.AffectedItems[0].ID})
		//if err != nil {
		//	t.Fatal(err)
		//}
		item := list.AffectedItems[0]
		fmt.Println(item.Name,
			item.Os.Name+" "+item.Os.Version,
			item.DateAdd,
			item.LastKeepAlive,
			item.Status,
		)
	}
}
func TestScan(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
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
	err := request.InitServerUrl("https://156.240.95.168")
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
	list, err := SCAList(req, SCAListReq{Agent: "001"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestSCADetailsList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
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
	list, err := SCADetailsList(req, SCADetailsListReq{Agent: "001", Policy: "cis_centos7_linux"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestSysCheckList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
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
	err := request.InitServerUrl("https://156.240.95.168")
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
	list, err := CiscatList(req, "001")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list.TotalAffectedItems, list.TotalFailedItems)
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}

func TestVulnerabilityList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
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
	list, err := VulnerabilityList(req, "006")
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range list.AffectedItems {
		fmt.Println(v)
	}
}
func TestVulnerabilityESList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
	if err != nil {
		t.Fatal(err)
	}
	err = request.InitToken("wazuh", "AgI_kwQ2GQ8v354EQtd6pSpT7bDjdaNJ")
	if err != nil {
		t.Fatal(err)
	}
	req, err := request.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	breakF := true
	for breakF {
		list, err := VulnerabilityESList(req, ESListReq{
			//Agent: "001", Severity: "Critical",
			Start: 1630982235, End: 1631068635,
			//Limit: 5,
			Offset: 20,
		})
		if err != nil {
			t.Fatal(err)
		} else {
			fmt.Println(list.Total)
			for k, v := range list.Hits {
				fmt.Println(k, v)
				fmt.Println("=================================")
			}
			breakF = false
		}
	}
}
func TestVirusESList(t *testing.T) {
	err := request.InitServerUrl("https://156.240.95.168")
	if err != nil {
		t.Fatal(err)
	}
	err = request.InitToken("wazuh", "AgI_kwQ2GQ8v354EQtd6pSpT7bDjdaNJ")
	if err != nil {
		t.Fatal(err)
	}
	req, err := request.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	breakF := true
	for breakF {
		list, err := VirusESList(req, ESListReq{Agent: "001", Limit: 400, Offset: 5})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(len(list.Hits), list.Total)
			breakF = false
		}
	}
}
