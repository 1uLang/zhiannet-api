package targets

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
}
func Test_add(t *testing.T) {
	res, err := AddAddr(&WebscanAddr{
		UserId:     1,
		TargetId:   "target_id",
		CreateTime: int(time.Now().Unix()),
	})
	fmt.Println(res, err)
}

func Test_del(t *testing.T) {
	err := DeleteByTargetIds([]string{"target_id"})
	fmt.Println(err)
}
