package edge_ssl_policies_server

import "github.com/1uLang/zhiannet-api/common/model/edge_ssl_policies"

func CheckAndUpdate(id uint64,check []string,update string)error  {
	return edge_ssl_policies.CheckAndUpdate(id,check,update)
}