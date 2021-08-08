package server

import (
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

type asset struct{ req *request.Request }

func (this *asset)List(args *asset_model.ListReq) ([]interface{}, uint64, error) {
	return asset_model.List(this.req,args)
}
func (this *asset)Create(args *asset_model.CreateReq)error  {
	return asset_model.Create(this.req,args)
}
func (this *asset)Update(args *asset_model.UpdateReq)error  {
	return asset_model.Update(this.req,args)
}
func (this *asset)Delete(args *asset_model.DeleteReq)error  {
	return asset_model.Delete(this.req,args)
}
func (this *asset)Details(args *asset_model.DetailsReq)(map[string]interface{},error  ){
	return asset_model.Details(this.req,args)
}
func (this *asset)Authorize(args *asset_model.AuthorizeReq) error  {
	return asset_model.Authorize(this.req,args)
}
func (this *asset)Connect(args *asset_model.ConnectReq) (string,error  ){
	return asset_model.Connect(this.req,args)
}
func (this *asset)AuthorizeUserList(args *asset_model.AuthorizeUserListReq) ([]string, error) {
	return asset_model.AuthorizeUserList(this.req,args)
}