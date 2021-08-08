package server

import (
	cert_model "github.com/1uLang/zhiannet-api/next-terminal/model/cert"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

type cert struct{ req *request.Request }

func (this *cert)List(args *cert_model.ListReq) ([]interface{}, int64, error) {
	return cert_model.List(this.req,args)
}
func (this *cert)Create(args *cert_model.CreateReq)error  {
	return cert_model.Create(this.req,args)
}
func (this *cert)Update(args *cert_model.UpdateReq)error  {
	return cert_model.Update(this.req,args)
}
func (this *cert)Delete(args *cert_model.DeleteReq)error  {
	return cert_model.Delete(this.req,args)
}
func (this *cert)Details(args *cert_model.DetailsReq)(map[string]interface{},error  ){
	return cert_model.Details(this.req,args)
}
func (this *cert)Authorize(args *cert_model.AuthorizeReq) error  {
	return cert_model.Authorize(this.req,args)
}
func (this *cert)AuthorizeUserList(args *cert_model.AuthorizeUserListReq) (*cert_model.AuthorizeUserListResp, error) {
	return cert_model.AuthorizeUserList(args)
}