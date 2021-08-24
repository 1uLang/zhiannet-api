package request

import (
	"testing"

	"github.com/1uLang/zhiannet-api/common/model"

	param "github.com/1uLang/zhiannet-api/resmon/const"
)

func init() {
	model.InitMysqlLink()
}
func TestAgentList(t *testing.T) {
	param.BASE_URL = "http://127.0.0.1:7777"
	param.TEA_KEY = "63b467f790de84a3588651d7dc04c25f"

	agents, err := AgentList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(agents.Total)
	for _, v := range agents.List {
		t.Logf("%v\n", v)
	}
}

func TestAddAgent(t *testing.T) {
	param.BASE_URL = "http://127.0.0.1:7777"
	param.TEA_KEY = "63b467f790de84a3588651d7dc04c25f"

	err := AddAgent("test", "127.0.0.1", true, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateAgent(t *testing.T) {
	param.BASE_URL = "http://127.0.0.1:7777"
	param.TEA_KEY = "63b467f790de84a3588651d7dc04c25f"
	agentID := "d8ee9357a53ca11d"

	err := UpdateAgent("test1", "127.0.0.2", agentID, true, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteAgent(t *testing.T) {
	param.BASE_URL = "http://127.0.0.1:7777"
	param.TEA_KEY = "63b467f790de84a3588651d7dc04c25f"
	agentID := "722b27e31a364703"

	err := DeleteAgent(agentID)
	if err != nil {
		t.Fatal(err)
	}
}

// func Test_check(t *testing.T) {
// 	check := new(CheckNode)
// 	check.Run()
// }

func TestGetAgentState(t *testing.T) {
	param.BASE_URL = "https://182.150.0.107:7443"
	param.TEA_KEY = "6d71e77857c1ddf791c65d44982a6d0a"

	as, err := getAgentState("51e831b35da771a9")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(as)
}
