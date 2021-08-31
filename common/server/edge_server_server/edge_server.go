package edge_server_server

import edge_server_model "github.com/1uLang/zhiannet-api/common/model/edge_server"
//waf服务列表
func GetServerList() ([]edge_server_model.EdgeServer,error) {
	return edge_server_model.GetList()

}
