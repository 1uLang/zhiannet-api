package risk

import "fmt"

type searchReq struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type SearchReq struct {
	searchReq
	UserName     string `json:"userName,omitempty"`
	MacCode      string `json:"macCode,omitempty"`
	ServerIp     string `json:"serverIp,omitempty"`
	WeakType     string `json:"weakType,omitempty"`
	Level        int    `json:"level,omitempty"`        //风险等级：低 中 高 危机：1,2,3,4
	ProcessState int    `json:"processState,omitempty"` //处理状态 待处理 已处理 不适用 1,2,3
}

func (this *SearchReq) Check() (bool, error) {
	if this.Level != 0 && this.Level != 1 && this.Level != 2 && this.Level != 3 && this.Level != 4 {
		return false, fmt.Errorf("风险等级参数错误")
	}
	if this.ProcessState != 0 && this.ProcessState != 1 && this.ProcessState != 2 && this.ProcessState != 3 {
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
	AbnormalLoginCountInfoList   []map[string]interface{} `json:"abnormalLoginCountInfoList"`   //异常登录
	AbnormalProcessCountInfoList []map[string]interface{} `json:"abnormalProcessCountInfoList"` //异常进程
	SystemCmdInfoList            []map[string]interface{} `json:"systemCmdInfoList"`            //命令篡改
}

type SystemDistributedResp struct {
	Low      int                      `json:"low"`      // 低危险
	Middle   int                      `json:"middle"`   // 中危险
	High     int                      `json:"high"`     // 高危险
	Critical int                      `json:"critical"` //危急漏洞
	Total    int                      `json:"total"`    //漏洞总数
	Host     int                      `json:"host"`     //受影响主机数
	List     []map[string]interface{} //列表
}

type ProcessReq struct {
	Opt string `json:"opt"`
	Req struct {
		MacCode string   `json:"macCode"`
		RiskIds []int    `json:"riskIds"`
		ItemIds []string `json:"itemIds"`
	}
}

var opts = map[string]bool{
	"add_trust":     true,
	"cancel_trust":  true,
	"isolate":       true,
	"revert":        true,
	"delete":        true,
	"close":         true,
	"cancel_close":  true,
	"ignore":        true,
	"cancel_ignore": true,
	"repair":        true,
}

func (this *ProcessReq) Check() (bool, error) {

	if _, isExist := opts[this.Opt]; !isExist {
		return false, fmt.Errorf("opt参数有误")
	}

	if this.Req.MacCode == "" {
		return false, fmt.Errorf("macCode不能为空")
	}
	if len(this.Req.RiskIds) == 0 {
		return false, fmt.Errorf("风险项id不能为空")
	}
	if len(this.Req.ItemIds) == 0 {
		return false, fmt.Errorf("报告id集合不能为空")
	}
	return true, nil
}

type DetailReq struct {
	MacCode string `json:"macCode"`
	Req     struct {
		UserName     string `json:"userName"`
		PageNo       int    `json:"pageNo"`
		PageSize     int    `json:"pageSize"`
		Level        int    `json:"level,omitempty"`        //漏洞风险 特有字段
		ProcessState int    `json:"ProcessState,omitempty"` //漏洞风险 特有字段
		State        int    `json:"state,omitempty"`        //入侵威胁 特有字段
	}
}

func (this *DetailReq) Check() (bool, error) {
	if this.MacCode == "" {
		return false, fmt.Errorf("请输入机器码")
	}
	if this.Req.Level < 0 || this.Req.Level > 4 {
		return false, fmt.Errorf("风险等级无效")
	}
	if this.Req.ProcessState < 0 || this.Req.ProcessState > 2 {
		return false, fmt.Errorf("处理状态无效")
	}
	return true, nil
}

type DetailResp struct {
	PageNo    int `json:"pageNo"`
	PageSize  int `json:"pageSize"`
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
	//弱口令列表
	WeakInfoList []map[string]interface{} `json:"weakInfoList"`
	//高危账号列表
	DangerAccountList []map[string]interface{} `json:"dangerAccountList"`
	//配置缺陷列表
	ConfigDefectList []map[string]interface{} `json:"configDefectList"`

	//木马病毒列表
	ServerVirusInfoList []map[string]interface{} `json:"serverVirusInfoList"`
	//网页后门列表
	WebshellInfoLis []map[string]interface{} `json:"webshellInfoLis"`
	//反弹shell列表
	ReboundshellInfoList []map[string]interface{} `json:"reboundshellInfoList"`
	//异常账号列表
	AbnormalAccountInfoList []map[string]interface{} `json:"abnormalAccountInfoList"`
	//日志异常删除列表
	LogDeleteInfoList []map[string]interface{} `json:"logDeleteInfoList"`
	//异常登录列表
	ServerAbnormalLoginInfoList []map[string]interface{} `json:"serverAbnormalLoginInfoList"`
	//异常进程列表
	AbnormalProcessInfoList []map[string]interface{} `json:"abnormalProcessInfoList"`
	//系统命令篡改列表
	SystemCmdInfoList []map[string]interface{} `json:"systemCmdInfoList"`
}
