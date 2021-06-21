package examine

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//主机体检

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

//ScanServerNow 立即体检
func ScanServerNow(args *ScanReq) (err error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "post"
	req.Path = _const.Examine_scan_server_now_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	info, err := model.ParseResp(resp)
	fmt.Println(info)
	return err
}

//ScanServerCancel 取消体检
func ScanServerCancel(macCodes []string) error {

	if len(macCodes) == 0 {
		return fmt.Errorf("参数错误：机器码集合不能为空")
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "post"
	req.Path = _const.Examine_scan_server_cancel_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = map[string]interface{}{
		"macCodes": macCodes,
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)

	return err

}

//Details 体检详情
func Details(macCode string) (info DetailsResp, err error) {

	//获取系统漏洞
	req := &risk.SearchReq{MacCode: macCode}
	riskInfo, err := risk.SystemDistributed(req)
	if err != nil {
		return info, fmt.Errorf("获取系统漏洞失败：%v", err)
	}
	info.Risk = riskInfo.Total
	//弱口令
	weakInfo, err := risk.WeakList(req)
	if err != nil {
		return info, fmt.Errorf("获取弱口令失败：%v", err)
	}
	info.Weak = weakInfo.Total
	//危险账号
	dangerAccountInfo, err := risk.DangerAccountList(req)
	if err != nil {
		return info, fmt.Errorf("获取危险账号失败：%v", err)
	}
	info.DangerAccount = dangerAccountInfo.Total
	//配置缺陷
	ConfigDefectInfo, err := risk.ConfigDefectList(req)
	if err != nil {
		return info, fmt.Errorf("获取配置缺陷失败：%v", err)
	}
	info.ConfigDefect = ConfigDefectInfo.Total
	return info, nil
}
