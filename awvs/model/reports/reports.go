package reports

import (
	"crypto/tls"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

//List 报表列表
func List(args *ListReq) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Reports_api_url
	args.Limit = 100
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	list, err = model.ParseResp(resp)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if args.UserId == 0 && args.AdminUserId == 0 {
		return list, err
	}
	//获取数据库 当前用户的扫描用户
	targetList, total, err := GetList(&ReportListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return map[string]interface{}{}, err
	}
	tarMap := map[string]int{}
	for _, v := range targetList {
		tarMap[v.ReportId] = 0
	}
	resList := gjson.ParseBytes(resp)
	list = map[string]interface{}{}
	if resList.Get("reports").Exists() {
		reports := []interface{}{}
		for _, v := range resList.Get("reports").Array() {
			if _, ok := tarMap[v.Get("report_id").String()]; ok {
				reports = append(reports, v.Value())
			}
		}
		list["reports"] = reports
	}
	return list, err
}

//Delete 删除报表
func Delete(report_id string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "DELETE"
	req.Url += _const.Reports_api_url + "/" + report_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)

	if err == nil {
		DeleteByTargetIds([]string{report_id})
	}
	return err
}

//Create 生成报表
func Create(args *CreateResp) (info map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "post"
	req.Url += _const.Reports_api_url
	req.Params = model.ToMap(args)
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	info, err = model.ParseResp(resp)
	if err != nil {
		return nil, err
	}
	//为生成的reportid 分配用户
	//if len(args.TargetIds) > 0 {
	if reportId, ok := info["report_id"]; ok {
		AddAddr(&WebscanReport{
			ReportId:    fmt.Sprintf("%v", reportId),
			AdminUserId: args.AdminUserId,
			UserId:      args.UserId,
			CreateTime:  int(time.Now().Unix()),
		})
	}
	//}
	return info, err
}

func Download(url string, pdf bool) ([]byte, string, error) {
	Client := resty.New().SetDebug(false).SetTimeout(time.Second * 60)
	Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	index, err := Client.R().Get(url)
	if err != nil {
		return nil, "", err
	}

	contents := index.Header().Get("Content-Disposition")
	var bytes string

	bytes = string(index.Body())

	//startIdx := strings.Index(bytes, "<img class=\"logo\"")
	//endIdx := strings.Index(bytes, "<h2 class=\"page-break ax-section-title ax-section-title--big\">")
	//bytes = bytes[:startIdx] + bytes[endIdx:]
	//
	//bytes = strings.Replace(bytes, ".98in", "0", 1)
	//bytes = strings.Replace(bytes, "Acunetix", "Zhiannet", 1)
	//bytes = strings.Replace(bytes, "Acunetix", "", 1)

	//综合报表
	bytes = strings.Replace(bytes, "class=\"col-xs-12 col-sm-6 end-sm\"", "class=\"col-xs-12 col-sm-6 end-sm\" style=\"display:none\"", 1)
	bytes = strings.Replace(bytes, "class=\"col-xs-12 col-sm-6\" id=\"logoContainer\"", "class=\"col-xs-12 col-sm-6\" id=\"logoContainer\" style=\"display:none\"", 1)
	bytes = strings.Replace(bytes, "Acunetix", "", 4)
	if pdf { //pdf
		bts, err := util.HTML2PDF(bytes)
		if err != nil {
			fmt.Println(err)
		} else {
			bytes = string(bts)
			contents = strings.Replace(contents, "html", "pdf", 1)
		}
	}

	return []byte(bytes), contents, nil
}
