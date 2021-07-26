package server

import (
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type assets struct{ req *request.Request }

func (this *assets) List(args *assets_model.ListReq) ([]map[string]interface{}, error) {
	return assets_model.List(this.req, args)
}

func (this *assets) Create(args *assets_model.CreateReq) (map[string]interface{}, error) {
	return assets_model.Create(this.req, args)
}
func (this *assets) Update(args *assets_model.UpdateReq) (map[string]interface{}, error) {
	return assets_model.Update(this.req, args)
}
func (this *assets) Delete(id string) error {
	return assets_model.Delete(this.req, id)
}
