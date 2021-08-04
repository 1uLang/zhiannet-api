package audit

import (
	"fmt"
	"github.com/dlclark/regexp2"
	//"regexp"
	"testing"
)

func Test_login(t *testing.T) {
	//server.GetLoginInfo(&server.UserReq{
	//	AdminUserId: 1,
	//})
}

func Test_re(t *testing.T) {
	str := "12345678Aa45"
	reg, err := regexp2.Compile(
		`^(?![A-z0-9]+$)(?=.[^%&',;=?$\x22])(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).{8,30}$`, 0)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	if match, err := reg.FindStringMatch(str); err != nil || match == nil {
		fmt.Println(match, err)
	} else {
		fmt.Println("匹配正确", match)
	}

	//(\?![A-z0-9]+)(\?=.[^%&',;=?$\\x22])(\?=.*[a-z])(\?=.*[A-Z])(\?=.*[0-9])
}
