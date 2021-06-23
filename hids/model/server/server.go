package server

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//List 主机列表
func List(args *SearchReq) (list SearchResp, err error) {

	list = SearchResp{}
	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.ServerList_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	fmt.Println(resp)
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//Info 主机信息
func Info(serverIp string) (info map[string]interface{}, err error) {
	list, err := List(&SearchReq{ServerIp: serverIp})
	if err != nil {
		return info, err
	}
	if len(list.ServerInfoList) == 0 {
		return info, fmt.Errorf("无该主机信息")
	}
	return list.ServerInfoList[0], nil
}

var osTypeName = map[string]string{
	"1":  "Windows2003 ",
	"2":  "Windows2008 ",
	"3":  "Windows2016 ",
	"4":  "Centos ",
	"5":  "Ubuntu ",
	"6":  "Debian ",
	"7":  "OpenSUSE ",
	"8":  "SUSE ",
	"9":  "RedHat ",
	"10": "Windows2012 ",
	"11": "RedFlag ",
	"12": "NeoKylin ",
	"13": "WindowsVista ",
	"14": "WIN7 ",
	"15": "WIN8 ",
	"16": "Windows10 ",
	"17": "FreeBSD ",
	"18": "Fedora ",
	"19": "Scientific ",
}
