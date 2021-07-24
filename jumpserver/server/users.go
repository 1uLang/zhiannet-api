package server

import (
	users_model "github.com/1uLang/zhiannet-api/jumpserver/model/users"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type users struct{ req *request.Request }

func (this *users) List(args *users_model.ListReq) ([]map[string]interface{}, error) {
	return users_model.List(this.req, args)
}
func (this *users) Create(args *users_model.CreateReq) (map[string]interface{}, error) {
	return users_model.Create(this.req, args)
}
