package monitor_list

import (
	"github.com/1uLang/zhiannet-api/monitor/model/edge_users"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_list"
	"time"
)

type (
	//列表请求参数
	ListReq struct {
		MonitorType int `json:"monitor_type"` // 1端口监控 2web监控
		PageNum     int `json:"page_num"`     //页数
		PageSize    int `json:"page_size"`    //每页条数
	}
	//列表响应参数
	ListResp struct {
		Id       uint64 `json:"id"`
		Status   int    `json:"status"`   //1状态启用 0状态关闭
		Addr     string `json:"addr"`     //IP地址 或者 url
		Port     int    `json:"port"`     //端口
		Code     int    `json:"code"`     //http响应码
		Username string `json:"username"` //归属用户名称
	}

	//添加监控请求参数
	AddReq struct {
		MonitorType int    `json:"monitor_type"` // 1端口监控 2web监控
		UserId      uint64 `json:"user_id"`      //所属用户ID
		Addr        string `json:"addr"`         //IP地址 url
		Port        int    `json:"port"`         //端口
		Code        int    `json:"code"`         //http 响应码
	}

	//删除监控请求参数
	DelReq struct {
		Id uint64 `json:"id"`
	}
)

func GetList(req *ListReq) (list []*ListResp, total int64, err error) {
	var res []*monitor_list.MonitorList
	res, total, err = monitor_list.GetList(&monitor_list.ListReq{
		MonitorType: req.MonitorType,
		PageNum:     req.PageNum,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return
	}
	if len(res) > 0 {
		var uids = make([]uint64, len(res))
		for k, v := range res {
			uids[k] = v.UserId
		}
		//通过用户ID  获取用户信息
		var uidList map[uint64]*edge_users.EdgeUsers
		uidList, _, err = edge_users.GetListByUid(uids)
		if err != nil {
			return
		}
		list = make([]*ListResp, len(res))
		for k, v := range res {
			username := ""
			if user, ok := uidList[v.UserId]; ok {
				username = user.Username
			}
			li := &ListResp{
				Id:       v.Id,
				Status:   v.Status,
				Addr:     v.Addr,
				Port:     v.Port,
				Code:     v.Code,
				Username: username,
			}
			list[k] = li
		}
	}
	return
}

//添加
func Add(req *AddReq) (id uint64, err error) {
	id, err = monitor_list.Add(&monitor_list.MonitorList{
		Addr:        req.Addr,
		Port:        req.Port,
		Code:        req.Code,
		Status:      1,
		UserId:      req.UserId,
		CreateTime:  int(time.Now().Unix()),
		MonitorType: req.MonitorType,
	})

	return
}

//删除
func Del(req *DelReq) (res bool, err error) {
	ids := []uint64{req.Id}
	err = monitor_list.DeleteByIds(ids)

	return
}
