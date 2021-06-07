package user

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/nessus/model/user"
)

//Add 创建用户
func Add(req *user.AddReq) error {

	id, err := user.Add(req)
	if err != nil {
		return err
	}
	fmt.Println("create nessus account success : ", id)
	return nil
}

//Delete 删除用户
func Delete(id string) error {
	err := user.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//Update 修改用户信息
func Update(req *user.UpdateReq) error {
	err := user.Update(req)
	if err != nil {
		return err
	}
	return nil
}

//ChangePassword 修改用户密码
func ChangePassword(req *user.ChangePasswordReq) error {
	err := user.ChangePassword(req)
	if err != nil {
		return err
	}
	return nil
}

//Enabled 启用/禁用账户
func Enabled(req *user.EnableReq) error {
	err := user.Enable(req)
	if err != nil {
		return err
	}
	return nil
}

//APIKeys 获取用户apikeys
func APIKeys(id string) (accessKey, secretKey string, err error) {
	return user.APIKeys(id)
}
