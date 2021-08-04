package server

import "github.com/1uLang/zhiannet-api/edgeUsers/model"

func ListEnabledUsers(req *model.ListReq) ([]*model.Edgeusers, error) {
	return model.GetList(req)
}
func CountAllEnabledUsers(req *model.GetNumReq) (int64, error) {
	return model.GetNum(req)
}
func CheckUserUsername(req *model.CheckUserNameReq)(bool,error)  {
	return model.CheckUserUsername(req)
}
func UpdateUser(req *model.UpdateUserReq)  error{
	return model.UpdateUser(req)
}
func CreateUser(req *model.CreateUserReq)(uint64,error)  {
	return model.CreateUser(req)
}
func DeleteUser(req *model.DeleteUserReq)error  {
	return model.DeleteUser(req)
}
func FindUserFeatures(req *model.FindUserFeaturesReq)([]string,error)  {
	return model.FindUserFeatures(req)
}
func UpdateUserFeatures(req *model.UpdateUserFeaturesReq) error {
	return model.UpdateUserFeatures(req)
}