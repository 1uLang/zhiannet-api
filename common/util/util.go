package util

import (
	"fmt"
	"strconv"
)

func Interface2Int(i interface{}) (int,error){
	return strconv.Atoi(fmt.Sprintf("%v",i))
}

func Interface2Uint64(i interface{})  (uint64,error){
	return strconv.ParseUint(fmt.Sprintf("%v", i), 10, 64)
}