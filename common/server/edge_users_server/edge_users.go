package edge_users_server

import (
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/model/edge_logins"
	"github.com/1uLang/zhiannet-api/common/model/edge_node_clusters"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"github.com/1uLang/zhiannet-api/common/server/edge_logins_server"
	"gorm.io/gorm"
	"time"
)

type (
	Users struct {
		edge_users.EdgeUsers `json:"users"`
		Cluster              Cluster  `json:"cluster"`   //节点
		Channels             Channels `json:"channels"`  //渠道信息
		SubTotal             int64    `json:"sub_total"` //子账号数
		OtpLoginIsOn         bool     `json:"otpLoginIsOn"`
	}

	Cluster struct {
		Id   uint64 `json:"id"`
		Name string `json:"name"`
	}

	Channels struct {
		Id   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

//判断用户密码是否过期 有效90天
func CheckPwdInvalid(id uint64) (res bool, err error) {
	info, err := edge_users.GetInfoById(id)
	if err != nil || info == nil {
		return
	}
	if int64(info.PwdAt) < time.Now().Add(-time.Second*60*60*24*90).Unix() {
		res = true
	}
	return
}

//更新密码
func UpdatePwd(id uint64, pwd string) (res int64, err error) {
	return edge_users.UpdatePwd(id, pwd)

}

//更新密码时间
func UpdatePwdAt(id uint64) (res int64, err error) {
	return edge_users.UpdatePwdAt(id)

}

//更新渠道ID
func UpdateChannel(id, chanId uint64) (res int64, err error) {
	return edge_users.UpdateChannel(id, chanId)
}

func GetUserInfo(id uint64) (info *edge_users.EdgeUsers, err error) {
	info, err = edge_users.GetInfoById(id)

	return
}

func GetUserInfoByName(name string) (info *edge_users.EdgeUsers, err error) {
	info, err = edge_users.GetInfoByUsername(name)
	if err == gorm.ErrRecordNotFound { //可能找不到数据
		err = nil
	}
	return
}

func GetChannelUserTotal(channelId []uint64, subUse bool) (total map[uint64]int64, err error) {
	total = make(map[uint64]int64, 0)
	t, err := edge_users.GetChannelUserTotal(channelId, subUse)
	if err != nil {
		return total, err
	}
	if len(t) > 0 {
		for _, v := range t {
			total[v.ChannelId] = v.Total
		}
	}

	return total, err
}

func GetList(req *edge_users.ListReq) (list []*Users, total int64, err error) {
	var lists []*edge_users.EdgeUsers
	lists, total, err = edge_users.GetList(req)
	if err != nil {
		return
	}
	if len(lists) > 0 {

		clusterMap, clusterids := map[uint64]Cluster{}, []uint64{0}
		channelMap, channelids := map[uint64]Channels{}, []uint64{0}
		subMap, subids := map[uint64]int64{}, []uint64{0}
		uidMap := map[uint64]*edge_logins.EdgeLogins{}
		for _, v := range lists {
			clusterids = append(clusterids, v.Clusterid)
			channelids = append(channelids, v.ChannelId)
			if v.ID > 0 {
				subids = append(subids, v.ID)
			}
		} //获取节点信息
		clusters, _, err := edge_node_clusters.GetList(&edge_node_clusters.ListReq{
			PageNum: 1, PageSize: 999, Ids: clusterids,
		})
		if err != nil {
			return list, total, err
		}
		if len(clusters) > 0 {
			//处理节点信息
			for _, v := range clusters {
				clusterMap[v.Id] = Cluster{
					Id:   v.Id,
					Name: v.Name,
				}
			}
		}
		//获取渠道列表
		channel, _, err := channels.GetList(&channels.ChannelReq{
			PageNum: 1, PageSize: 999, Ids: channelids,
		})
		if err != nil {
			return list, total, err
		}
		if len(channel) > 0 {
			//处理渠道信息
			for _, v := range channel {
				channelMap[v.Id] = Channels{
					Id:   v.Id,
					Name: v.Name,
				}
			}
		}

		//统计子账号数
		sub, err := edge_users.GetSubUserTotal(subids)
		if err != nil {
			return list, total, err
		}
		if len(sub) > 0 {
			//处理渠道信息
			for _, v := range sub {
				subMap[v.ParentId] = v.Total
			}
		}

		//获取otp信息
		uidMap, _, err = edge_logins_server.GetListByUid(subids)
		if err != nil {
			return list, total, err
		}

		for _, v := range lists {
			clu := Cluster{}
			if c, ok := clusterMap[v.Clusterid]; ok { //匹配节点
				clu = c
			}
			cha := Channels{}
			if c, ok := channelMap[v.ChannelId]; ok { //匹配渠道
				cha = c
			}
			subTotal := int64(0)
			if t, ok := subMap[v.ID]; ok { //匹配子账号数量
				subTotal = t
			}

			//匹配otp
			var IsOn bool
			if len(uidMap) > 0 {
				if info, ok := uidMap[v.ID]; ok && info.IsOn == 1 && info.Type == "otp" {
					IsOn = true
				}
			}

			list = append(list, &Users{
				Cluster:      clu,
				Channels:     cha,
				SubTotal:     subTotal,
				EdgeUsers:    *v,
				OtpLoginIsOn: IsOn,
			})

		}
	}
	return list, total, err
}
