package scans

import "fmt"

type ListReq struct {
	Limit       int    `json:"l,omitempty"` //限制条数
	C           int    `json:"c,omitempty"` //偏移量
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}

type AddReq struct {
	TargetId         string `json:"target_id"`                    //目标id
	ProfileId        string `json:"profile_id"`                   //扫描类型
	ReportTemplateId string `json:"report_template_id,omitempty"` //报表类型
	Schedule         struct { //计划设置
		Disable       bool    `json:"disable"`        //启用
		StartDate     *string `json:"start_date"`     //生效时间规则：DTSTART:20210607T134000\nFREQ=MONTHLY;INTERVAL=1
		TimeSensitive bool    `json:"time_sensitive"` //时间敏感
	} `json:"schedule"`
	Incremental bool `json:"incremental"` //增量式
}

func (this *AddReq) Check() (bool, error) {

	if this.TargetId == "" {
		return false, fmt.Errorf("目标id不能为空")
	}
	if this.ProfileId == "" {
		return false, fmt.Errorf("扫描类型不能为空")
	}
	if this.ReportTemplateId != "" {
		list, err := ScanningProfiles()
		if err != nil {
			return false, err
		}
		profifles := map[string]bool{}
		for _, v := range list["scanning_profiles"].([]map[string]interface{}) {
			profifles[v["profile_id"].(string)] = true
		}
		_, isExist := profifles[this.ReportTemplateId]
		if !isExist {
			return false, fmt.Errorf("报表类型无效")
		}
	}

	return true, nil
}

type VulnerabilitiesReq struct {
	ScanId        string
	ScanSessionId string
	VulId         string
}

func (this *VulnerabilitiesReq) Check() (bool, error) {

	if this.ScanId == "" {
		return false,fmt.Errorf("扫描id不能为空")
	}
	if this.ScanSessionId == "" {
		return false,fmt.Errorf("扫描会话id不能为空")
	}
	if this.VulId == "" {
		return false,fmt.Errorf("漏洞id不能为空")
	}
	return true,nil
}