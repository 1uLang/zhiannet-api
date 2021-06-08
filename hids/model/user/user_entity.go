package user

import "fmt"

type AddReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Role     int    `json:"roleId"`
	OrgId    int    `json:"orgId"`
}

func (this *AddReq) Check() (bool, error) {

	if this.UserName == "" {
		return false, fmt.Errorf("用户账号不能为空")
	}
	//TODO:需添加密码复杂度检测
	if this.Password == "" {
		return false, fmt.Errorf("用户密码不能为空")
	}
	if this.Role == 0 {
		return false, fmt.Errorf("用户角色不能为空")
	}
	return true, nil
}

type SearchReq struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	UserName string `json:"userName"`
	userID   string `json:"userId"`
	phoneNum string `json:"phoneNum"`
	email    string `json:"email"`
	state    int    `json:"state"`
}
type SearchResp struct {
	PageNo       int `json:"pageNo"`
	PageSize     int `json:"pageSize"`
	TotalData    int `json:"totalData"`
	TotalPage    int `json:"totalPage"`
	UserInfoList []struct {
		UserName string `json:"userName"`
		userID   string `json:"userId"`
		phoneNum string `json:"phoneNum"`
		email    string `json:"email"`
		state    int    `json:"state"`
		QQ       string `json:"qq"`
		Weibo    string `json:"weibo"`
	} `json:"userInfoList"`
}
