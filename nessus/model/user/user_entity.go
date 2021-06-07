package user

import "fmt"

type AddReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Permissions string `json:"permissions"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

func (this *AddReq) Check() (bool, error) {

	if this.Permissions != "32" && this.Permissions != "128" {
		return false, fmt.Errorf("permissions 参数错误")
	}
	if this.Username == "" || this.Password == "" {
		return false, fmt.Errorf("账户密码不能为空")
	}
	return true, nil
}

type UpdateReq struct {
	ID          string `json:"id"`
	Permissions string `json:"permissions"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Enabled     bool   `json:"enabled"` //启用、禁用 账户 true/false
}

func (this *UpdateReq) Check() (bool, error) {

	if this.Permissions != "32" && this.Permissions != "128" {
		return false, fmt.Errorf("权限设置错误")
	}

	if this.ID == "" {
		return false, fmt.Errorf("用户id不能为空")
	}
	return true, nil
}

type ChangePasswordReq struct {
	ID              string `json:"id"`
	CurrentPassword string `json:"current_password"`
	Password        string `json:"password"`
}

func (this *ChangePasswordReq) Check() (bool, error) {

	if this.ID == "" {
		return false, fmt.Errorf("用户id不能为空")
	}
	if this.CurrentPassword == "" {
		return false, fmt.Errorf("当前密码不能为空")
	}
	if this.Password == "" {
		return false, fmt.Errorf("新密码不能为空")
	}

	return true, nil
}

type EnableReq struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (this *EnableReq) Check() (bool, error) {

	if this.ID == "" {
		return false, fmt.Errorf("用户id不能为空")
	}

	return true, nil
}
