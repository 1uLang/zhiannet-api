package risk

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/server"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
	"sync"
	"time"
)

//系统漏洞 - 入侵威胁

//SystemRiskList 服务器系统漏洞信息列表
func SystemRiskList(args *SearchReq) (list SearchResp, err error) {

	list = SearchResp{}

	ok, err := args.Check()
	if err != nil || !ok {
		return list, err
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
	req.Path = _const.Risk_system_info_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	fmt.Println(resp)
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//Dashboard 首页概览数据 主机总数 - 待处理入侵事件 - 待处理高危漏洞 - 已开启安全监控主机
func Dashboard(userName string) (info DashboardResp, err error) {

	info = DashboardResp{}
	dashboardWg := sync.WaitGroup{}
	go func() {
		dashboardWg.Add(1)
		//获取主机数  在线主机数
		hostLock := sync.RWMutex{}
		hostWg := sync.WaitGroup{}
		go func() {
			hostWg.Add(1)
			sr, _ := server.List(&server.SearchReq{UserName: userName})
			hostLock.Lock()
			info.Host = sr.TotalData
			hostLock.Unlock()
			hostWg.Done()
		}()
		go func() {
			hostWg.Add(1)
			sr, _ := server.List(&server.SearchReq{UserName: userName, ServerStatus: "1"})
			hostLock.Lock()
			info.Host = sr.TotalData
			hostLock.Unlock()
			hostWg.Done()
		}()
		hostWg.Wait()
		dashboardWg.Done()
	}()

	//获取八大入侵事件总数
	go func() {
		dashboardWg.Add(1)
		invadeLock := sync.RWMutex{}
		invadeWg := sync.WaitGroup{}
		fns := []func(*RiskSearchReq) (RiskSearchResp, error){
			VirusList,
			WebShellList,
			ReboundList,
			AbnormalAccountList,
			LogDeleteList,
			AbnormalLoginList,
			AbnormalProcessList,
			SystemCmdList,
		}
		args := &RiskSearchReq{UserName: userName}

		for _, fn := range fns {
			go func() {
				invadeWg.Add(1)
				risk, _ := fn(args)
				invadeLock.Lock()
				info.Invade += risk.TotalData
				invadeLock.Unlock()
				invadeWg.Done()
			}()
		}
		invadeWg.Wait()
		dashboardWg.Done()
	}()

	//待处理高危漏洞
	{
		args := &SearchReq{UserName: userName, Level: 3, ProcessState: 1}
		args.PageSize = 100
		args.PageNo = 1
		today := 0
		nowStr := time.Now().Format("2006-01-02")
	Do:
		di, err := SystemRiskList(args)
		if err != nil {
			return info, err
		}

		//统计今天高危漏洞
		for _, v := range di.SystemRiskInfoList {
			if v["time"].(string) > nowStr {
				today++
			}
		}
		if di.TotalData/args.PageSize+1 > args.PageNo {
			goto Do
		}

		info.Risk = di.TotalData
		info.TodayRisk = today
	}

	dashboardWg.Wait()
	return info, nil
}

func riskList(path string, args *RiskSearchReq) (list RiskSearchResp, err error) {

	list = RiskSearchResp{}
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
	req.Path = path
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	fmt.Println(resp)
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//VirusList 入侵威胁 - 病毒木马列表
func VirusList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 1 && args.State != 2 && args.State != 3 && args.State != 101 &&
		args.State != 201 && args.State != 301 && args.State != -1 && args.State != -2 && args.State != -3 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_Virus_api_url, args)
}

//WebShellList 网页后门列表
func WebShellList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 1 && args.State != 2 && args.State != 3 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_webshell_api_url, args)
}

//ReboundList 反弹shell数据列表
func ReboundList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 7 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_reboundshell_api_url, args)
}

//AbnormalAccountList 异常账号数量列表
func AbnormalAccountList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 7 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_abnormal_account_api_url, args)
}

//LogDeleteList 日志异常删除数量列表
func LogDeleteList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 7 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_log_delete_api_url, args)
}

//AbnormalLoginList 异常登录数量列表
func AbnormalLoginList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 7 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_abnormal_login_api_url, args)
}

//AbnormalProcessList 异常进程数量列表
func AbnormalProcessList(args *RiskSearchReq) (list RiskSearchResp, err error) {
	if args.State != 0 && args.State != 7 {
		return list, fmt.Errorf("处理状态参数错误")
	}
	return riskList(_const.Risk_abnormal_process_api_url, args)
}

//SystemCmdList 命令篡改数量列表
func SystemCmdList(args *RiskSearchReq) (list RiskSearchResp, err error) {

	return riskList(_const.Risk_system_cmd_api_url, args)
}

//SystemDistributed 风险概览 ： 系统漏洞总数
func SystemDistributed(args *SearchReq) (info SystemDistributedResp, err error) {

	info = SystemDistributedResp{}
	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = _const.Risk_distributed_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	ret, err := model.ParseResp(resp)
	if _, isExist := ret["data"]; !isExist || err != nil {
		return info, err
	}
	ret = ret["data"].(map[string]interface{})
	info.Host, _ = util.Interface2Int(ret["totalData"])
	for _, v := range ret["systemRiskDistributionInfoList"].([]interface{}) {
		node := v.(map[string]interface{})
		low, _ := util.Interface2Int(node["lowRiskCount"])
		middle, _ := util.Interface2Int(node["middleRiskCount"])
		high, _ := util.Interface2Int(node["highRiskCount"])
		critical, _ := util.Interface2Int(node["criticalCount"])

		info.Low += low
		info.Middle += middle
		info.High += high
		info.Critical += critical
		info.List = append(info.List, node)
	}
	info.Total = info.Low + info.Middle + info.High + info.Critical
	return info, err
}

//WeakList 弱口令分布列表
func WeakList(args *SearchReq) (info SystemDistributedResp, err error) {

	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = _const.Risk_weak_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	ret, err := model.ParseResp(resp)
	if _, isExist := ret["data"]; !isExist || err != nil {
		return info, err
	}
	ret = ret["data"].(map[string]interface{})
	info.Host, _ = util.Interface2Int(ret["totalData"])
	for _, v := range ret["weakDistributionInfoList"].([]interface{}) {
		node := v.(map[string]interface{})
		low, _ := util.Interface2Int(node["lowRiskCount"])
		middle, _ := util.Interface2Int(node["middleRiskCount"])
		high, _ := util.Interface2Int(node["highRiskCount"])
		critical, _ := util.Interface2Int(node["criticalCount"])

		info.Low += low
		info.Middle += middle
		info.High += high
		info.Critical += critical
		info.List = append(info.List, node)
	}
	info.Total = info.Low + info.Middle + info.High + info.Critical
	return info, err
}

//DangerAccountList 高危账号分布列表
func DangerAccountList(args *SearchReq) (info SystemDistributedResp, err error) {

	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = _const.Risk_danger_account_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	ret, err := model.ParseResp(resp)

	if _, isExist := ret["data"]; !isExist || err != nil {
		return info, err
	}
	ret = ret["data"].(map[string]interface{})
	info.Host, _ = util.Interface2Int(ret["totalData"])
	for _, v := range ret["dangerAccountDistributionInfoList"].([]interface{}) {
		node := v.(map[string]interface{})
		low, _ := util.Interface2Int(node["lowRiskCount"])
		middle, _ := util.Interface2Int(node["middleRiskCount"])
		high, _ := util.Interface2Int(node["highRiskCount"])
		critical, _ := util.Interface2Int(node["criticalCount"])

		info.Low += low
		info.Middle += middle
		info.High += high
		info.Critical += critical
		info.List = append(info.List, node)
	}
	info.Total = info.Low + info.Middle + info.High + info.Critical
	return info, err
}

//ConfigDefectList 配置缺陷分布列表
func ConfigDefectList(args *SearchReq) (info SystemDistributedResp, err error) {

	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = _const.Risk_config_defect_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	ret, err := model.ParseResp(resp)
	if _, isExist := ret["data"]; !isExist || err != nil {
		return info, err
	}
	ret = ret["data"].(map[string]interface{})
	info.Host, _ = util.Interface2Int(ret["totalData"])
	for _, v := range ret["configDefectDistributionInfoList"].([]interface{}) {
		node := v.(map[string]interface{})
		low, _ := util.Interface2Int(node["lowRiskCount"])
		middle, _ := util.Interface2Int(node["middleRiskCount"])
		high, _ := util.Interface2Int(node["highRiskCount"])
		critical, _ := util.Interface2Int(node["criticalCount"])

		info.Low += low
		info.Middle += middle
		info.High += high
		info.Critical += critical
		info.List = append(info.List, node)
	}
	info.Total = info.Low + info.Middle + info.High + info.Critical
	return info, err
}

func process(args *ProcessReq, path string) error {
	if ok, err := args.Check(); err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "post"
	//忽略
	req.Path = path + "/" + args.Opt
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args.Req)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}

//ProcessRisk 处置服务器系统漏洞
func ProcessRisk(args *ProcessReq) error {
	return process(args, _const.Risk_process_api_url)
}

//ProcessWeak 处置服务器弱口令
func ProcessWeak(args *ProcessReq) error {
	return process(args, _const.Risk_weak_process_api_url)
}

//ProcessDangerAccount 处置服务器弱口令
func ProcessDangerAccount(args *ProcessReq) error {
	return process(args, _const.Risk_danger_account_process_api_url)
}

//ProcessConfigDefect 处置服务器配置缺陷
func ProcessConfigDefect(args *ProcessReq) error {
	return process(args, _const.Risk_config_defect_process_api_url)
}

func detail(args *DetailReq, path string) (info DetailResp, err error) {
	ok, err := args.Check()
	if err != nil || !ok {
		return info, err
	}

	if args.Req.PageSize == 0 {
		args.Req.PageSize = 10
	}
	if args.Req.PageNo == 0 {
		args.Req.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = path + "/" + args.MacCode + "/detail/list"
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	_, err = model.ParseResp(resp, &info)
	return info, err
}

//SystemRiskDetail 系统漏洞详情
func SystemRiskDetail(macCode, riskId string) (info map[string]interface{}, err error) {

	if macCode == "" || riskId == "" {
		return info, fmt.Errorf("参数错误：机器码和漏洞id不能为空")
	}

	req, err := request.NewRequest()
	if err != nil {
		return info, err
	}
	req.Method = "get"
	req.Path = _const.Risk_system_detail_api_url + "/" + macCode + "/" + riskId
	req.Headers["signNonce"] = util.RandomNum(10)

	resp, err := req.Do()
	if err != nil {
		return info, err
	}
	_, err = model.ParseResp(resp, &info)
	return info, err
}

//WeakDetail 弱口令详情
func WeakDetail(args *DetailReq) (info DetailResp, err error) {

	return detail(args, _const.Risk_weak_server_api_url)
}

//DangerAccountDetail 高危账号详情
func DangerAccountDetail(args *DetailReq) (info DetailResp, err error) {

	return detail(args, _const.Risk_danger_account_detail_api_url)
}

//ConfigDefectDetail 配置缺陷详情
func ConfigDefectDetail(args *DetailReq) (info DetailResp, err error) {

	return detail(args, _const.Risk_config_defect_detail_api_url)
}

//VirusDetailList 入侵威胁病毒木马详情列表
func VirusDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_Virus_detail_list_api_url)
}

//WebShellDetailList 入侵威胁网页后门详情列表
func WebShellDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_webshell_detail_list_api_url)
}

//ReboundDetailList 入侵威胁详情列表
func ReboundDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_reboundshell_detail_list_api_url)
}

//AbnormalAccountDetailList 入侵威胁异常账号详情列表
func AbnormalAccountDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_abnormal_account_detail_list_api_url)
}

//LogDeleteDetailList 入侵威胁日志异常删除详情列表
func LogDeleteDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_log_delete_detail_list_api_url)
}

//AbnormalLoginDetailList 入侵威胁异常登录详情列表
func AbnormalLoginDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_abnormal_login_detail_list_api_url)
}

//AbnormalProcessDetailList 入侵威胁异常进程详情列表
func AbnormalProcessDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_abnormal_process_detail_list_api_url)
}

//SystemCmdDetailList 入侵威胁命令篡改详情列表
func SystemCmdDetailList(args *DetailReq) (info DetailResp, err error) {
	return detail(args, _const.Risk_system_cmd_detail_list_api_url)
}

//riskDetail 入侵威胁详情
func riskDetail(macCode, id, path string) (map[string]interface{}, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(path, macCode, id)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	ret := map[string]interface{}{}
	_, err = model.ParseResp(resp, &ret)
	return ret, err

}

//VirusDetail 入侵威胁病毒木马详情
func VirusDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_Virus_detail_api_url)
}

//WebShellDetail 入侵威胁网页后门详情
func WebShellDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_webshell_detail_api_url)
}

//ReboundShellDetail 入侵威胁 反弹shell详情
func ReboundShellDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_reboundshell_detail_api_url)
}

//AbnormalAccountDetail 入侵威胁 异常账号详情
func AbnormalAccountDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_abnormal_account_detail_api_url)
}

//LogDeleteDetail 入侵威胁 日志异常删除详情
func LogDeleteDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_log_delete_detail_api_url)
}

//AbnormalLoginDetail 入侵威胁 异常登录详情
func AbnormalLoginDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_abnormal_login_detail_api_url)
}

//AbnormalProcessDetail 入侵威胁 异常进程详情
func AbnormalProcessDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_abnormal_process_detail_api_url)
}

//SystemCmdDetail 入侵威胁 系统命令篡改详情
func SystemCmdDetail(macCode, id string) (map[string]interface{}, error) {
	return riskDetail(macCode, id, _const.Risk_system_cmd_detail_api_url)
}

//VirusProcess 入侵威胁 病毒木马处理
func VirusProcess(args *ProcessReq) error {
	return process(args, _const.Risk_Virus_process_api_url)
}

//WebShellProcess 入侵威胁 网页后门
func WebShellProcess(args *ProcessReq) error {
	return process(args, _const.Risk_webshell_process_api_url)
}

//ReboundShellProcess 入侵威胁 反弹shell
func ReboundShellProcess(args *ProcessReq) error {
	return process(args, _const.Risk_reboundshell_process_api_url)
}

//AbnormalAccountProcess 入侵威胁 异常账号
func AbnormalAccountProcess(args *ProcessReq) error {
	return process(args, _const.Risk_abnormal_account_process_api_url)
}

//LogDeleteProcess 入侵威胁 日志异常删除
func LogDeleteProcess(args *ProcessReq) error {
	return process(args, _const.Risk_log_delete_process_api_url)
}

//AbnormalLoginProcess 入侵威胁 异常登录
func AbnormalLoginProcess(args *ProcessReq) error {
	return process(args, _const.Risk_abnormal_login_process_api_url)
}

//AbnormalProcessProcess 入侵威胁 异常进程
func AbnormalProcessProcess(args *ProcessReq) error {
	return process(args, _const.Risk_abnormal_process_process_api_url)
}

//SystemCmdProcess 入侵威胁 系统命令篡改
func SystemCmdProcess(args *ProcessReq) error {
	return process(args, _const.Risk_system_cmd_process_api_url)
}
