package login

import (
	"testing"
)

func TestLogin(t *testing.T) {
	ap := ApiKey{}
	ap.Addr = "https://hids.zhiannet.com"
	ap.IsSsl = true
	ap.Username = "dengbao"
	ap.Password = "Cloud123!@#"
	pp := Passport{}
	_,err := pp.Login(&ap)
	if err != nil {
		t.Fatal(err)
	}
}
