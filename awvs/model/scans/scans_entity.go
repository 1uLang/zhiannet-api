package scans

import "fmt"

type AddReq struct {
	TargetId         string   `json:"target_id"`                    //目标id
	ProfileId        string   `json:"profile_id"`                   //扫描类型
	ReportTemplateId string   `json:"report_template_id,omitempty"` //报表类型
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
