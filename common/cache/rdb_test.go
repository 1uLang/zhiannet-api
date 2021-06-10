package cache

import (
	"fmt"
	"testing"
	"time"
)

func Test_get(t *testing.T) {
	fmt.Println(InitClient())
	go CheckCache("a", getA, 10, true)
	go CheckCache("a", getA, 10, true)
	//fmt.Println(d)
	time.Sleep(10 * time.Second)
}
