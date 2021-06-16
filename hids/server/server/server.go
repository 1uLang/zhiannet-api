package server

import "github.com/1uLang/zhiannet-api/hids/model/server"

func List(req *server.SearchReq) (server.SearchResp, error) {
	return server.List(req)
}
func Info(serverIp string) (info server.InfoResp, err error) {
	return server.Info(serverIp)
}
