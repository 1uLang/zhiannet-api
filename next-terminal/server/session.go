package server

import (
	session_model "github.com/1uLang/zhiannet-api/next-terminal/model/session"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

type session struct{ req *request.Request }

func (this *session)List(args *session_model.ListReq) ([]map[string]interface{}, error) {
	return session_model.List(this.req,args)
}
func (this *session)Delete(args *session_model.DeleteReq)error  {
	return session_model.Delete(this.req,args)
}
func (this *session)Replay(args *session_model.ReplayReq)([]byte,error  ){
	return session_model.Replay(this.req,args)
}
func (this *session)Monitor(args *session_model.MonitorReq)(*session_model.MonitorResp,error  ){
	return session_model.Monitor(this.req,args)
}
func (this *session)DisConnect(args *session_model.DisConnectReq)error  {
	return session_model.DisConnect(this.req,args)
}