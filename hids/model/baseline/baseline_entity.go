package baseline

type SearchReq struct {
	UserName    string `json:"userName,omitempty"`    //用户名
	MacCode     string `json:"macCode,omitempty"`     //机器码
	PageNo      int    `json:"pageNo"`                //页码
	PageSize    int    `json:"pageSize"`              //显示条数
	State       *int   `json:"state,omitempty"`       //检查状态 0未检查 1 检查中  2 已完成 3 检查失败
	ResultState int    `json:"resultState,omitempty"` //检查结论 1 基线异常 2 基线正常
}

type SearchResp struct {
	PageNo    int                      `json:"pageNo"`
	PageSize  int                      `json:"pageSize"`
	TotalData int                      `json:"totalData"`
	TotalPage int                      `json:"totalPage"`
	List      []map[string]interface{} `json:"serverBaselineCheckInfoList"`
}

type CheckReq struct {
	MacCodes   []string `json:"macCodes"`
	TemplateId int      `json:"templateId"`
}

type TemplateSearchReq struct {
	UserName string `json:"userName,omitempty"` //用户名
	PageNo   int    `json:"pageNo"`             //页码
	PageSize int    `json:"pageSize"`           //显示条数
}
type TemplateSearchResp struct {
	PageNo    int                      `json:"pageNo"`
	PageSize  int                      `json:"pageSize"`
	TotalData int                      `json:"totalData"`
	TotalPage int                      `json:"totalPage"`
	List      []map[string]interface{} `json:"templateList"`
	Items     []map[string]interface{} `json:"templateItemList"`
}
type TemplateDetailReq struct {
	ID  string `json:"id"`
	Req TemplateSearchReq
}

type DetailReq struct {
	MacCode  string
	PageNo   int `json:"pageNo"`   //页码
	PageSize int `json:"pageSize"` //显示条数
}
type DetailResp struct {
	PageNo    int                      `json:"pageNo"`
	PageSize  int                      `json:"pageSize"`
	TotalData int                      `json:"totalData"`
	TotalPage int                      `json:"totalPage"`
	List      []map[string]interface{} `json:"templateItemList"`
}
