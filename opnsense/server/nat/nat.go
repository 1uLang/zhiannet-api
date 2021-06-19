package nat

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/nat"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	ListReq struct {
		NodeId uint64 `json:"node_id"`
	}
	InfoReq struct {
		NodeId uint64 `json:"node_id"`
		Id     string `json:"id"`
	}

	SaveNat1To1Req struct {
		NodeId        uint64   `json:"node_id"`
		ID            string   `json:"id"`
		Disabled      bool     `json:"disabled"`      //启用
		Interface     string   `json:"interface"`     //接口
		Type          string   `json:"type"`          //类型
		External      string   `json:"external"`      //外部地址
		Srcnot        bool     `json:"srcnot"`        //源 反转
		Src           string   `json:"src"`           //内部地址
		Srcmask       string   `json:"srcmask"`       //内部地址 掩码
		Dstnot        bool     `json:"dstnot"`        //目标反转
		Dst           string   `json:"dst"`           //目的地
		Dstmask       string   `json:"dstmask"`       //目的地掩码
		Category      []string `json:"category"`      //分类 多选
		Descr         string   `json:"descr"`         //描述
		Natreflection string   `json:"natreflection"` //NAT回流
	}
	//启动 停止 请求参数
	StartNat1To1Req struct {
		NodeId uint64 `json:"node_id"`
		Id     string `json:"id"`
	}

	//删除请求参数
	DelNat1To1Req struct {
		NodeId uint64 `json:"node_id"`
		Id     string `json:"id"`
	}
)

//获取nat 1：1 列表
func GetNat1To1List(req *ListReq) (list []*nat.Nat1To1ListResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}

	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	return nat.GetNat1To1List(loginInfo)
}

//获取 nat 1：1详情
func GetNat1To1Info(req *InfoReq) (res *nat.Nat1To1InfoResp, err error) {
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
	return nat.GetNat1To1Info(&nat.Nat1To1InfoReq{
		Id: req.Id,
	}, loginInfo)
}

//添加修改
func SaveNat1To1(req *SaveNat1To1Req) (res []string, err error) {
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
	var info *nat.Nat1To1InfoResp
	info, err = nat.GetNat1To1Info(&nat.Nat1To1InfoReq{
		Id: req.ID,
	}, loginInfo)
	if err != nil {
		return res, err
	}
	//填充数据
	reqMap := map[string]string{}
	reqCateMap := map[string][]string{}
	reqMap, reqCateMap, err = supplyData(req, info)
	if req.ID == "" {
		res, err = nat.AddNat1To1(reqMap, reqCateMap, loginInfo)
	} else {
		res, err = nat.EditNat1To1(reqMap, reqCateMap, loginInfo)
	}
	if len(res) == 0 && err == nil { //应用修改
		nat.Apply(loginInfo)
	}
	return res, err
}

//填充数据
//参数1编辑提交的参数
//参数2当前ID的数据详情
//返回参数 组合起来的map数据（不丢失未提交原本有的数据的数据）
func supplyData(saveReq *SaveNat1To1Req, info *nat.Nat1To1InfoResp) (res map[string]string, cates map[string][]string, err error) {
	res = make(map[string]string)
	{
		//TODO 获取的详情
		//启用状态
		if info.Disabled {
			res["disabled"] = "yes"
		}
		//接口
		if len(info.Interface) > 0 {
			for _, v := range info.Interface {
				if v.Selected {
					res["interface"] = v.Value
					break
				}
			}
		}
		//类型
		if len(info.Type) > 0 {
			for _, v := range info.Type {
				if v.Selected {
					res["type"] = v.Value
					break
				}
			}
		}
		//外部网络
		res["external"] = info.External
		//源反转
		if info.Srcnot {
			res["srcnot"] = "yes"
		}
		//源
		res["src"] = info.Src
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
		//类别
		cates = make(map[string][]string)
		if len(info.Category) > 0 {
			cate := []string{}
			for _, v := range info.Category {
				if v.Selected {
					cate = append(cate, v.Value)
				}
			}
			cates["category"] = cate
		}
		//描述
		res["descr"] = info.Descr

		//nat 回流
		if len(info.Natreflection) > 0 {
			for _, v := range info.Natreflection {
				if v.Selected {
					res["natreflection"] = v.Value
					break
				}
			}
		}
	}
	//提交的数据
	res["interface"] = saveReq.Interface
	res["type"] = saveReq.Type
	res["external"] = saveReq.External
	res["src"] = saveReq.Src
	res["srcmask"] = saveReq.Srcmask
	res["dst"] = saveReq.Dst
	res["dstmask"] = saveReq.Dstmask
	res["descr"] = saveReq.Descr
	res["id"] = saveReq.ID

	return
}

//启动停止
func StartUpNat1To1(req *StartNat1To1Req) (res bool, err error) {
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
	return nat.StartUpNat1To1(req.Id, loginInfo)
}

//删除
func DelNat1To1(req *DelNat1To1Req) (res bool, err error) {
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
	return nat.DelNat1To1(req.Id, loginInfo)
}

//应用 使修改生效
func ApplyNat1To1(nodeId uint64) (res bool, err error) {
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
	return nat.Apply(loginInfo)
}
