package session

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	"github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

var session_path = "/sessions"
//会话列表
func List(req *request.Request ,args *ListReq)([]map[string]interface{},error  ){

	if args.Status !="connected" && args.Status != "disconnected"{
		return nil,fmt.Errorf("参数错误")
	}

	req.Path = session_path+"/paging"
	req.Method = "get"
	req.Params = model.ToMap(args)

	resp,err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, fmt.Errorf("服务器异常：%v",resp.Message)
	}
	assets,total,err := asset.GetList(&asset.ListReq{})
	if err != nil || total == 0 {
		return nil,err
	}
	contain := map[string]bool{}
	for _,v := range assets {
		contain[v.AssetsId] = true
	}
	fmt.Println(contain)
	fmt.Println(resp.Data)
	ret := make([]map[string]interface{},0)
	for _,v := range resp.Data.(map[string]interface{})["items"].([]interface{}){
		item := v.(map[string]interface{})
		if _,isExist := contain[item["assetId"].(string)];isExist{
			ret = append(ret, item)
		}
	}
	return ret,nil
}
//删除历史会话
func Delete(req *request.Request ,args *DeleteReq)error  {
	req.Path = session_path+"/"+args.Id
	req.Method = "delete"
	req.Params = model.ToMap(args)

	resp,err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v",resp.Message)
	}
	return nil
}
//回放
func Replay(req *request.Request ,args *ReplayReq)([]byte,error)  {
	req.Path = session_path+"/"+args.Id+"/recording"
	req.Method = "get"
	req.Params = nil

	return req.Do()
}
//监控
func Monitor(req *request.Request ,args *MonitorReq)(*MonitorResp,error)  {

	return nil,nil
}
//断开
func DisConnect(req *request.Request ,args *DisConnectReq)error  {
	req.Path = session_path+"/"+args.Id+"/disconnect"
	req.Method = "post"
	req.Params = nil

	resp,err :=  req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v",resp.Message)
	}
	return nil
}