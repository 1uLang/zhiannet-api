package scans

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/nessus/model/scans"
)

// 创建扫描
func Create(req *scans.AddReq) error {

	id, err := scans.Create(req)
	if err != nil {
		return err
	}
	fmt.Println("create nessus scans success : ", id)
	return nil
}

