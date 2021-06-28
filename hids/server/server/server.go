package server

import "github.com/1uLang/zhiannet-api/hids/model/server"

func List(req *server.SearchReq) (server.SearchResp, error) {
	return server.List(req)
}
func Info(serverIp, userName string) (info map[string]interface{}, err error) {
	return server.Info(serverIp, userName)
}
