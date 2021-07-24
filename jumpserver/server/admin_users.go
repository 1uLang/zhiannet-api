package server

import (
	admin_users_model "github.com/1uLang/zhiannet-api/jumpserver/model/admin_users"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type admin_users struct{ req *request.Request }

func (this *admin_users) List(args *admin_users_model.ListReq) ([]map[string]interface{}, error) {
	return admin_users_model.List(this.req, args)
}

func (this *admin_users) Create(args *admin_users_model.CreateReq) (map[string]interface{}, error) {
	return admin_users_model.Create(this.req, args)
}
func (this *admin_users) Update(args *admin_users_model.UpdateReq) (map[string]interface{}, error) {
	return admin_users_model.Update(this.req, args)
}
func (this *admin_users) Delete(id string) error {
	return admin_users_model.Delete(this.req, id)
}
