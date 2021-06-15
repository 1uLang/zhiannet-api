package examine

import "github.com/1uLang/zhiannet-api/hids/model/examine"

func List(req *examine.SearchReq) (examine.SearchResp, error) {
	return examine.List(req)
}
