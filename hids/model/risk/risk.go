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

//RiskList 服务器系统漏洞信息列表
func RiskList(args *SearchReq) (list SearchResp, err error) {

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
			WebshellList,
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
		args := &SearchReq{UserName: userName, Level: 3, ProcessStatus: 1}
		args.PageSize = 100
		args.PageNo = 1
		today := 0
		nowStr := time.Now().Format("2006-01-02")
	Do:
		di, err := RiskList(args)
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

//WebshellList 网页后门列表
func WebshellList(args *RiskSearchReq) (list RiskSearchResp, err error) {
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
	for _, node := range ret["systemRiskDistributionInfoList"].([]map[string]interface{}) {
		low, _ := util.Interface2Int(node["lowRiskCount"])
		middle, _ := util.Interface2Int(node["middleRiskCount"])
		high, _ := util.Interface2Int(node["highRiskCount"])
		critical, _ := util.Interface2Int(node["criticalCount"])

		info.Low += low
		info.Middle += middle
		info.High += high
		info.Critical += critical
	}
	info.Total = info.Low + info.Middle + info.High + info.Critical
	return info, err
}

//WeakList 弱口令分布列表
func WeakList(args *SearchReq) (list map[string]interface{}, err error) {

	list = map[string]interface{}{}
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
	req.Path = _const.Risk_weak_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//DangerAccountList 高危账号分布列表
func DangerAccountList(args *SearchReq) (list map[string]interface{}, err error) {

	list = map[string]interface{}{}
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
	req.Path = _const.Risk_danger_account_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//ConfigDefectList 配置缺陷分布列表
func ConfigDefectList(args *SearchReq) (list map[string]interface{}, err error) {

	list = map[string]interface{}{}
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
	req.Path = _const.Risk_config_defect_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//ProcessRisk 处置服务器系统漏洞
func ProcessRisk(args *ProcessResp) error {

	if ok, err := args.Check(); err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "post"
	req.Path = _const.Risk_process_api_url + "/ignore" //忽略
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}
