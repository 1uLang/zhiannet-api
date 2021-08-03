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

func (this *assets) Authorize(args *assets_model.AuthorizeReq) error {
	return assets_model.Authorize(this.req, args)
}

func (this *assets) DelAuthorize(args *assets_model.DelAuthorizeReq) error {
	return assets_model.DelAuthorize(this.req, args)
}

func (this *assets) AuthorizeList(args *assets_model.AuthorizeListReq) ([]map[string]interface{},error ){
	return assets_model.AuthorizeList(this.req, args)
}

func (this *assets) Link(id string) (string,string,error ){
	return assets_model.Link(this.req, id)
}

func (this *assets) Info(id string) (map[string]interface{}, error) {
	return assets_model.Info(this.req, id)
}
func (this *assets) Refresh(id string) (string,error)  {
	return assets_model.Refresh(this.req, id)
}
func (this *assets) CheckLink(id string)(string,error)  {
	return assets_model.CheckLink(this.req, id)
}