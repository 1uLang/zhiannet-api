package edge_users_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"

	"regexp"
	"testing"
)

func init() {
	model.InitMysqlLink()
}
func Test_user_total(t *testing.T) {
	total, err := GetChannelUserTotal([]uint64{1, 5}, false)
	fmt.Println(total, err)
}

//列表
func Test_list(t *testing.T) {
	list, total, err := GetList(&edge_users.ListReq{})
	fmt.Println(list[0].SubTotal)
	fmt.Println(list, total, err)
}

func Test_re(t *testing.T) {
	re := regexp.MustCompile(`[A-Z]`)
	str := "fsdfDSF34FFFF324#@$2fdD"
	all := re.FindAllString(str, -1)
	fmt.Println(all)
}
