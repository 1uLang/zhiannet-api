package examine

import "fmt"

type SearchReq struct {
	PageNo       int      `json:"pageNo"`
	PageSize     int      `json:"pageSize"`
	UserName     string   `json:"userName,omitempty"`
	ServerIp     string   `json:"serverIp,omitempty"`
	State        int      `json:"state"`                  //体检状态：全部 -1、未体检 0、待体检 3、体检中 1、已完成 2
	Score        int      `json:"score"`                  //体检分数 全部 -1、不健康(0-59) 0、亚健康(60-89) 1、健康(90-100) 2
	Type         int      `json:"type"`                   //体检类型 全部 -1、非定时体检 1、定时体检2
	StartTime    string   `json:"startTime,omitempty"`    //体检开始时间
	EndTime      string   `json:"endTime,omitempty"`      //体检结束时间
	ExamineItems []string `json:"examineItems,omitempty"` //体检项目集合
}

func (this *SearchReq) Check() (bool, error) {

	if this.Type != -1 && this.Type != 1 && this.Type != 2 {
		return false, fmt.Errorf("体检类型参数错误")
	}
	if this.Score != -1 && this.Score != 0 && this.Score != 1 && this.Score != 2 {
		return false, fmt.Errorf("体检分数参数错误")
	}
	if this.State != -1 && this.State != 0 && this.State != 1 && this.State != 2 && this.State != 3 {
		return false, fmt.Errorf("体检状态参数错误")
	}
	if len(this.ExamineItems) > 0 {
		check := func(n string) bool {
			if n != "01" && n != "02" && n != "03" && n != "04" && n != "11" && n != "12" && n != "13" && n != "14" && n != "15" && n != "16" && n != "17" {
				return false
			}
			return true
		}
		for _, v := range this.ExamineItems {
			ok := check(v)
			if !ok {
				return ok, fmt.Errorf("体检项目参数错误")
			}
		}
	}
	return true, nil
}

type OnlineSearchReq struct {
	SearchReq
	Online bool `json:"online"` //是否在线
}
type SearchResp struct {
	PageNo                      int                      `json:"pageNo"`
	PageSize                    int                      `json:"pageSize"`
	TotalData                   int                      `json:"totalData"`
	TotalPage                   int                      `json:"totalPage"`
	ServerExamineResultInfoList []map[string]interface{} `json:"serverExamineResultInfoList"`
}
