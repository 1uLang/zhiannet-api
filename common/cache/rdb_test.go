package cache

import (
	"fmt"
	"testing"
)

func Test_get(t *testing.T) {
	fmt.Println(InitClient())
	go CheckCache("a", getA, 10, true)
	res, err, d := CheckCache("a", getA, 10, true)
	fmt.Println(res)
	fmt.Println(err)
	fmt.Println(d)
	//time.Sleep(10*time.Second)
}
