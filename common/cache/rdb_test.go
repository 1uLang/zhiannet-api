package cache

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	InitClient()
}
func Test_get(t *testing.T) {
	//fmt.Println(InitClient())
	go CheckCache("a", getA, 10, true)
	go CheckCache("a", getA, 10, true)
	//fmt.Println(d)
	time.Sleep(10 * time.Second)
}

func Test_setnx(t *testing.T) {
	fmt.Println(SetNx("", time.Second*51))
}

func Test_Incr(t *testing.T) {
	fmt.Println(Incr("", time.Second*51))
}

func Test_Int(t *testing.T) {
	fmt.Println(GetInt(""))
}

func Test_Del(t *testing.T) {
	fmt.Println(DelKey(""))
}
