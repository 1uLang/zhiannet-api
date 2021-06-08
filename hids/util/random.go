package util

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

//RandomNum 生成随机数
//参数：
//	l 随机数位数
func RandomNum(l int) string {

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		b := r.Intn(10) + 0x30
		bytes[i] = byte(b)
	}
	return string(bytes)
}
