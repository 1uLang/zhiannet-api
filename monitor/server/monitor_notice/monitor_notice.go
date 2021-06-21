package monitor_notice

import (
	"github.com/1uLang/zhiannet-api/monitor/model/edge_users"
	"github.com/1uLang/zhiannet-api/monitor/model/monitor_notice"
)

type (
	ListReq struct {
		Message  string `json:"message"`
		PageNum  int    `json:"page_num"`  //页数
		PageSize int    `json:"page_size"` //每页条数
	}
	ListResp struct {
		Message    string `json:"message"`
		Id         uint64 `json:"id"`
		Username   string `json:"username"` //归属用户名称
		CreateTime int    `json:"create_time"`
	}
	//删除监控请求参数
	DelReq struct {
		Id uint64 `json:"id"`
	}
)

func GetList(req *ListReq) (list []*ListResp, total int64, err error) {
	var res []*monitor_notice.MonitorNotice
	res, total, err = monitor_notice.GetList(&monitor_notice.ListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
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
				Id:         v.Id,
				Message:    v.Message,
				CreateTime: v.CreateTime,
				Username:   username,
			}
			list[k] = li
		}
	}
	return
}

func Del(req *DelReq) (res bool, err error) {
	ids := []uint64{req.Id}
	err = monitor_notice.DeleteByIds(ids)

	return
}
