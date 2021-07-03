package subassemblynode

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
}

func Test_node_num(t *testing.T) {
	total, e := GetNodeNum(&NodeNumReq{
		Type:  1,
		State: "1",
	})
	fmt.Println(total, e)
}
