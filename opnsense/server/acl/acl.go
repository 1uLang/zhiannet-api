package acl

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/acl"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	//获取详情请求参数
	InfoReq struct {
		NodeId uint64 `json:"node_id"`
		ID     string `json:"id"`
	}
	//更新请求参数
	SaveAclReq struct {
		NodeId     uint64 `json:"node_id"`
		ID         string `json:"id"`
		Type       string `json:"type"`       //操作
		Disabled   bool   `json:"disabled"`   //启用
		Interface  string `json:"interface"`  //接口
		Direction  string `json:"direction"`  //方向
		Ipprotocol string `json:"ipprotocol"` //tcp/IP版本
		Protocol   string `json:"protocol"`   //协议
		//Srcnot    bool     `json:"srcnot"`    //源 反转
		Src     string `json:"src"`     //内部地址
		Srcmask string `json:"srcmask"` //内部地址 掩码
		//Dstnot    bool     `json:"dstnot"`    //目标反转
		Dst     string `json:"dst"`     //目的地
		Dstmask string `json:"dstmask"` //目的地掩码
		//Category  []string `json:"category"`  //分类 多选
		Descr string `json:"descr"` //描述
	}
	//启动停止请求参数
	StartAclReq struct {
		NodeId    uint64 `json:"node_id"`
		Interface string `json:"interface"`
		ID        string `json:"id"`
	}
	//删除请求参数
	DelAclReq struct {
		NodeId    uint64 `json:"node_id"`
		Interface string `json:"interface"`
		ID        string `json:"id"`
	}
)

//获取列表
func GetAclList(nodeId uint64) (res []*acl.AclListResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: nodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	list1, err := acl.GetAclList("lan", loginInfo)
	if err != nil {
		return res, err
	}
	res = append(res, list1...)
	list2, err := acl.GetAclList("lo0", loginInfo)
	if err != nil {
		return res, err
	}
	res = append(res, list2...)
	list3, err := acl.GetAclList("wan", loginInfo)
	if err != nil {
		return res, err
	}
	res = append(res, list3...)
	return res, err
}

//获取详情
func GetAclInfo(req *InfoReq) (res *acl.AclInfoResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return acl.GetAclInfo(&acl.AclInfoReq{
		ID: req.ID,
	}, loginInfo)
}

//添加修改
func SaveAcl(req *SaveAclReq) (res []string, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}

	//先获取详细数据填充
	var info *acl.AclInfoResp
	info, err = acl.GetAclInfo(&acl.AclInfoReq{
		ID: req.ID,
	}, loginInfo)
	if err != nil {
		return res, err
	}
	//填充数据
	reqMap := map[string]string{}
	reqMap, err = supplyData(req, info)
	if req.ID == "" {
		res, err = acl.AddAcl(reqMap, loginInfo)
	} else {
		res, err = acl.EditAcl(reqMap, loginInfo)
	}
	if len(res) == 0 && err == nil { //应用修改
		acl.Apply(req.Interface, loginInfo)
	}
	return res, err
}

//合并数据
func supplyData(req *SaveAclReq, info *acl.AclInfoResp) (res map[string]string, err error) {
	res = make(map[string]string)
	//TODO 获取的详情
	//源数据
	{
		//操作类型
		if len(info.Type) > 0 {
			for _, v := range info.Type {
				if v.Selected {
					res["type"] = v.Value
					break
				}
			}
		}
		//状态
		if info.Disabled {
			res["disabled"] = "yes"
		}
		//快速
		if info.Quick {
			res["quick"] = "yes"
		}
		////接口
		if len(info.Interface) > 0 {
			for _, v := range info.Interface {
				if v.Selected {
					res["interface"] = v.Value
					break
				}
			}
		}
		//方向
		if len(info.Direction) > 0 {
			for _, v := range info.Direction {
				if v.Selected {
					res["direction"] = v.Value
					break
				}
			}
		}
		//tcp版本
		if len(info.Ipprotocol) > 0 {
			for _, v := range info.Ipprotocol {
				if v.Selected {
					res["ipprotocol"] = v.Value
					break
				}
			}
		}
		//协议
		if len(info.Protocol) > 0 {
			for _, v := range info.Protocol {
				if v.Selected {
					res["protocol"] = v.Value
					break
				}
			}
		}
		//反转
		if info.Srcnot {
			res["srcnot"] = "yes"
		}
		//源
		if len(info.Src) > 0 {
			for _, v := range info.Src {
				if v.Selected {
					res["src"] = v.Value
					break
				}
			}
		}
		//源掩码
		if len(info.Srcmask) > 0 {
			for _, v := range info.Srcmask {
				if v.Selected {
					res["srcmask"] = v.Value
					break
				}
			}
		}
		//目标
		if len(info.Dst) > 0 {
			for _, v := range info.Dst {
				if v.Selected {
					res["dst"] = v.Value
					break
				}
			}
		}
		//目标掩码
		if len(info.Dstmask) > 0 {
			for _, v := range info.Dstmask {
				if v.Selected {
					res["dstmask"] = v.Value
					break
				}
			}
		}
		//日志
		if info.Log {
			res["log"] = "yes"
		}
	}

	//提交数据
	res["type"] = req.Type
	res["direction"] = req.Direction
	res["interface"] = req.Interface
	res["ipprotocol"] = req.Ipprotocol
	res["protocol"] = req.Protocol
	res["src"] = req.Src
	res["srcmask"] = req.Srcmask
	res["dst"] = req.Dst
	res["dstmask"] = req.Dstmask
	res["descr"] = req.Descr
	res["id"] = req.ID
	return res, err
}

//启动停止
func StartUpAcl(req *StartAclReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return acl.StartUpAcl(req.ID, req.Interface, loginInfo)
}

//删除
func DelAcl(req *DelAclReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return acl.DelAcl(req.ID, req.Interface, loginInfo)
}
