package user

import "github.com/1uLang/zhiannet-api/hids/model/user"

func Add(req *user.AddReq) (uint64, error) {
	return user.Add(req)
}
func List(req *user.SearchReq)(user.SearchResp,error)  {
	return user.List(req)
}
