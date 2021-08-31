package login

import "testing"

func TestManagerLogin(t *testing.T) {
	ap := ApiKey{}
	ap.Addr = "https://hids.zhiannet.com/manager/main"
	ap.IsSsl = true
	ap.Username = "admin"
	ap.Password = "aqgDdfi72hv1!r!WNsafedog"
	ma := Manager{}
	_,err := ma.Login(&ap)
	if err != nil {
		t.Fatal(err)
	}
}

