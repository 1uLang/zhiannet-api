package util

import (
	"fmt"
	"testing"
)

func TestEncodePassword(t *testing.T) {
	pswd := "123456"
	fmt.Println(EncodePassword(pswd))
}
