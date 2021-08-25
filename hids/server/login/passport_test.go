package request

import (
	"testing"
)

func TestLogin(t *testing.T) {
	ap := ApiKey{}
	ap.Addr = "https://hids.zhiannet.com/manager/main"
	ap.IsSsl = true
	ap.Username = "admin"
	ap.Password = "aqgDdfi72hv1!r!WNsafedog"
	_,err := Login(&ap)
	if err != nil {
		t.Fatal(err)
	}
}
