package risk

import "fmt"

type searchReq struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type SearchReq struct {
	searchReq
	UserName      string `json:"userName,omitempty"`
	ServerIp      string `json:"serverIp,omitempty"`
	WeakType      string `json:"weakType,omitempty"`
	Level         int    `json:"level,omitempty"`         //风险等级：低 中 高 危机：1,2,3,4
	ProcessStatus int    `json:"processStatus,omitempty"` //处理状态 待处理 已处理 不适用 1,2,3
}

func (this *SearchReq) Check() (bool, error) {
	if this.Level != 1 && this.Level != 2 && this.Level != 3 && this.Level != 4 {
		return false, fmt.Errorf("风险等级参数错误")
	}
	if this.ProcessStatus != 1 && this.ProcessStatus != 2 && this.ProcessStatus != 3 {
		return false, fmt.Errorf("处理态度参数错误")
	}
	if this.WeakType != "" && this.WeakType != "0201" && this.WeakType != "0202" && this.WeakType != "0203" {
		return false, fmt.Errorf("弱口令类型参数错误")
	}
	return true, nil
}

type RiskSearchReq struct {
	searchReq
	UserName    string `json:"userName,omitempty"`
	ServerIp    string `json:"serverIp,omitempty"`
	IsProcessed bool   `json:"isProcessed,omitempty"` //待处理 false，已处理 true
	OnLine      bool   `json:"onLine,omitempty"`      //待处理 false，已处理 true	针对异常登录
	State       int    `json:"state,omitempty"`       //未处理 已关闭 0 、 1、7
}

type DashboardResp struct {
	Host       int `json:"host"`        //主机
	OnlineHost int `json:"online_host"` //在线主机数
	Invade     int `json:"invade"`      //入侵
	Risk       int `json:"risk"`        //高危漏洞
	TodayRisk  int `json:"today_risk"`  //今日高危漏洞
}

type searchResp struct {
	PageNo    int `json:"pageNo"`
	PageSize  int `json:"pageSize"`
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
}

type SearchResp struct {
	searchResp
	SystemRiskInfoList []map[string]interface{} `json:"systemRiskInfoList"`
}

type RiskSearchResp struct {
	searchResp
	VirusCountInfoList           []map[string]interface{} `json:"virusCountInfoList"`           //木马病毒列表
	WebshellCountInfoList        []map[string]interface{} `json:"webshellCountInfoList"`        //网页后门列表
	ReboundshellCountInfoList    []map[string]interface{} `json:"reboundshellCountInfoList"`    //反弹shell数量列表
	AbnormalAccountCountInfoList []map[string]interface{} `json:"abnormalAccountCountInfoList"` //异常账号数量列表
	LogDeleteCountInfoList       []map[string]interface{} `json:"logDeleteCountInfoList"`       //日志异常删除
}

type SystemDistributedResp struct {
	Low      int `json:"low"`      // 低危险
	Middle   int `json:"middle"`   // 中危险
	High     int `json:"high"`     // 高危险
	Critical int `json:"critical"` //危急漏洞
	Total    int `json:"total"`    //漏洞总数
	Host     int `json:"host"`     //受影响主机数
}

type ProcessResp struct {
	Opt     string   `json:"opt"`
	MacCode string   `json:"macCode"`
	RiskIds []string `json:"riskIds"`
	ItemIds []string `json:"ItemIds"`
}

func (this *ProcessResp) Check() (bool, error) {

	if this.Opt != "ignore" && this.Opt != "cancel_ignore" && this.Opt != "repair" {
		return false, fmt.Errorf("opt参数有误")
	}
	if this.MacCode == "" {
		return false, fmt.Errorf("macCode不能为空")
	}
	if len(this.RiskIds) == 0 {
		return false, fmt.Errorf("风险项id不能为空")
	}
	if len(this.ItemIds) == 0 {
		return false, fmt.Errorf("报告id集合不能为空")
	}
	return true, nil
}
