package util

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	enc := Encode([]byte(fmt.Sprintf("%x",md5.Sum([]byte("Cloud123!@#safedog")))))
	fmt.Println(enc)
}
