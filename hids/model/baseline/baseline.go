package baseline

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
	"strconv"
)

//合规基线

//List 合规基线列表
func List(args *SearchReq) (list SearchResp, err error) {

	agentList := make([]map[string]interface{},0)
	agents,total ,err := agent.GetList(&agent.ListReq{UserId: args.UserId,AdminUserId: args.AdminUserId})
	if err != nil || total == 0{
		return list,err
	}
	contain := map[string]int{}
	for k,v := range agents{
		contain[v.IP] = k
		agentList = append(agentList, map[string]interface{}{"serverIp":v.IP})
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.Baseline_check_list_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	args.PageSize = 100
	args.PageNo = 1
	args.UserName = model.HidsUserNameAPI
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)

	for _,item := range list.List{
		if idx,isExist := contain[item["serverIp"].(string)];isExist{
			agentList[idx] = item
		}
	}
	list.List = agentList
	return list, err
}

//Check 基线检测
func Check(args *CheckReq) (err error) {

	if len(args.MacCodes) == 0 {
		return fmt.Errorf("机器码不能为空")
	}
	//检测模板是否存在
	info, err := TemplateDetail(&TemplateDetailReq{ID: strconv.Itoa(args.TemplateId)})
	if err != nil {
		return fmt.Errorf("获取合规基线模板信息失败：%v", err)
	}
	if info.TotalData == 0 {
		return fmt.Errorf("无效的合规基线模板，该模板不存在")
	}
	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "post"
	req.Path = _const.Baseline_check_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}

//TemplateList 模板列表
func TemplateList(args *TemplateSearchReq) (list TemplateSearchResp, err error) {
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
	req.Path = _const.Baseline_template_list_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//TemplateDetail 模板详情
func TemplateDetail(args *TemplateDetailReq) (list TemplateSearchResp, err error) {

	if args.ID == "" {
		return list, fmt.Errorf("合规基线模板id不能为空")
	}
	if args.Req.PageSize == 0 {
		args.Req.PageSize = 10
	}
	if args.Req.PageNo == 0 {
		args.Req.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Baseline_template_detail_api_url, args.ID)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args.Req)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//Detail 合规基线详情
func Detail(args *DetailReq) (list DetailResp, err error) {

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
	req.Path = fmt.Sprintf(_const.Baseline_check_detail_api_url, args.MacCode)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = map[string]interface{}{
		"pageNo":   args.PageNo,
		"pageSize": args.PageSize,
	}

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}
