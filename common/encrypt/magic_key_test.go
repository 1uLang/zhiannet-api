package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

func TestMagicKeyEncode(t *testing.T) {
	var code string
	var year int
	var month int
	var day int
	var hour int
	var err error
	code = "1f0eea9db4384439b4c7e1d230d9f2af"
	year = 1
	now := time.Now()
	nowt := now.Unix()
	//addt := now.AddDate(year,month,day).Add(time.Duration(hour) * time.Hour)
	addt := now.AddDate(year,month,day).Add(time.Duration(hour) * time.Hour)
	renewal := addt.Unix() - nowt
	timeout := now.Add(5 * time.Minute).Unix()
	dst := MagicKeyEncode([]byte(fmt.Sprintf("%v,%v,%v",code,renewal,timeout)))

	dst = []byte(base64.StdEncoding.EncodeToString(dst))
	fmt.Println(string(dst))

	dst, err = base64.StdEncoding.DecodeString(string(dst))
	if err != nil {
		t.Fatal(err)
	}
	src := MagicKeyDecode(dst)
	fmt.Println(string(src))
}
