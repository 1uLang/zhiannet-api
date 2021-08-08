package server

import (
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/hids/model/user"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/request"
)

/*
	nessus 主机扫描api对接 server 层
*/
type CheckRequest struct{}

// SetUrl 初始化 Nessus APIKeys
func SetUrl(url string) error {
	return request.InitServerUrl(url)
}

// SetAPIKeys 初始化 Nessus APIKeys
func SetAPIKeys(req *request.APIKeys) error {
	return request.InitRequestAPIKeys(req)
}

func GetNessus() (resp *model.NessusResp, err error) {
	return model.GetNessusInfo()
}

//检测nessus 配置访问是否异常

func Check() (bool, uint64, error) {
	info, err := GetNessus()
	if err != nil {
		return false, 0, err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false, info.Id, err
	}
	err = SetAPIKeys(&request.APIKeys{info.Access, info.Secret})
	if err != nil {
		return false, info.Id, err
	}
	_, err = user.List(&user.SearchReq{})
	if err != nil {
		return false, info.Id, err
	}
	return true, info.Id, nil
}
func (this *CheckRequest) Run() {
	var conn int = 1
	res, id, _ := Check()
	if !res {
		conn = 0
	}
	if id > 0 {
		subassemblynode.UpdateConnState(id, conn)
	}

}
