package examine

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//List 体检列表
func List(args *SearchReq, online ...bool) (list SearchResp, err error) {

	list = SearchResp{}

	ok, err := args.Check()
	if err != nil || !ok {
		return list, fmt.Errorf("参数错误：%v", err)
	}

	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.Examine_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	if len(online) > 0 {
		req.Params = model.ToMap(OnlineSearchReq{SearchReq: *args, Online: online[0]})
	} else {
		req.Params = model.ToMap(args)
	}
	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}
