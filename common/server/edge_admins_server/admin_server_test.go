package edge_admins_server

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
}
func Test_initFiled(t *testing.T) {
	InitField()
}
