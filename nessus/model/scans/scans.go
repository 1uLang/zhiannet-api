package scans

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/util"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/request"
	"math/big"
	"strconv"
	"strings"
	"time"
)

//扫描目标模板

const (
	scan_templates_url = "/editor/scan/templates"
	scan_url           = "/scans"
	folders_url        = "/folders"
)

func getScanInfo(id string, history ...string) (map[string]interface{}, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += scan_url + "/" + id

	req.Params = map[string]interface{}{
		"limit":                              2500,
		"includeHostDetailsForHostDiscovery": true,
	}
	if len(history) > 0 {
		req.Params["history_id"] = history[0]
	}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}

	return model.ParseResp(resp)
}
func getScanFoldersId(name string) (float64, error) {

	req, err := request.NewRequest()
	if err != nil {
		return 0, err
	}

	req.Method = "get"
	req.Url += folders_url
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	var id float64
	templates := ret["folders"].([]interface{})
	for _, node := range templates {
		template := node.(map[string]interface{})
		if template["name"] == name { //主机漏洞扫描 模板
			id = template["id"].(float64)
		}
	}
	return id, nil

}
func getScanTemplateUUid() (string, error) {

	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}

	req.Method = "get"
	req.Url += scan_templates_url
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return "", err
	}
	var id string
	templates := ret["templates"].([]interface{})
	for _, node := range templates {
		template := node.(map[string]interface{})
		if template["title"] == "Host Discovery" { //主机漏洞扫描 模板
			id = template["uuid"].(string)
		}
	}
	return id, nil
}

//创建目标
func Create(args *AddReq) (uint64, error) {

	req, err := request.NewRequest()
	if err != nil {
		return 0, err
	}

	args.UUID, err = getScanTemplateUUid()
	if err != nil {
		return 0, err
	}
	req.Method = "post"
	req.Url += scan_url
	req.Params = model.ToMap(*args)

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}

	id, _ := util.Interface2Uint64(ret["scan"].(map[string]interface{})["id"])
	//写入数据库
	_, err = AddScans(&NessusScans{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		ScansId:     id,
		Description: args.Settings.Description,
		Addr:        args.Settings.Name,
		CreateTime:  int(time.Now().Unix()),
	})
	return id, err
}
func List(args *ListReq) ([]interface{}, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	folder_id, err := getScanFoldersId("My Scans")
	if err != nil {
		return nil, err
	}
	req.Method = "get"
	req.Params = map[string]interface{}{
		"folder_id": folder_id,
	}
	req.Url += scan_url
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	result := make(map[string]interface{}, 0)
	fmt.Println(string(ret))
	err = json.Unmarshal(ret, &result)
	if err != nil {
		return nil, err
	}
	scansList, total, err := GetList(&ScansListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return []interface{}{}, err
	}
	contain := map[uint64]string{}
	for _, v := range scansList {
		contain[v.ScansId] = v.Description
	}

	resList := make([]interface{}, 0)
	if result["scans"] == nil {
		return nil,nil
	}
	for _, v := range result["scans"].([]interface{}) {
		scan := v.(map[string]interface{})
		id, _ := util.Interface2Uint64(scan["id"])
		//对应web扫描
		if desc, isExist := contain[id]; isExist {

			if args.Scan { //扫描任务列表
				scan["scan_id"] = fmt.Sprintf("%v-host", scan["id"])
				scan["target"] = map[string]interface{}{
					"address":     scan["name"],
					"description": desc,
				}
				scan["current_session"] = map[string]interface{}{
					"status":     scan["status"],
					"start_date": float2TimeStr(scan["creation_date"].(float64)),
				}
			} else if args.Targets { //扫描目标列表

				scan["target_id"] = fmt.Sprintf("%v-host", scan["id"])

				scan["address"] = scan["name"]
				scan["description"] = desc

				scan["last_scan_session_status"] = scan["status"]
				scan["last_scan_date"] = float2TimeStr(scan["last_modification_date"].(float64))
			}
			//获取漏洞信息 扫描完成
			if scan["status"].(string) == "completed" {

				if args.Report { //扫描报告列表
					scan["report_id"] = fmt.Sprintf("%v-host", scan["id"])
					scan["generation_date"] = float2TimeStr(scan["last_modification_date"].(float64))
					scan["address"] = scan["name"]
				}

				info, err := getScanInfo(fmt.Sprintf("%v", scan["id"]))
				if err != nil {
					return nil, err
				}
				if len(info["hosts"].([]interface{})) > 0 {
					severity_counts := map[string]float64{}
					severity_counts["high"] = info["hosts"].([]interface{})[0].(map[string]interface{})["high"].(float64) + info["hosts"].([]interface{})[0].(map[string]interface{})["critical"].(float64)
					severity_counts["medium"] = info["hosts"].([]interface{})[0].(map[string]interface{})["medium"].(float64)
					severity_counts["low"] = info["hosts"].([]interface{})[0].(map[string]interface{})["low"].(float64)
					severity_counts["info"] = info["hosts"].([]interface{})[0].(map[string]interface{})["info"].(float64)

					scan["severity_counts"] = severity_counts
					if args.Scan {
						scan["current_session"].(map[string]interface{})["severity_counts"] = severity_counts
					}
				}
			} else if (scan["status"].(string) == "empty" && args.Scan) || args.Report { // 当是scan list 同时没有扫描过 不返回 当是报表列表 且没有完成扫描 不返回
				continue
			}
			resList = append(resList, scan)
		}
	}
	return resList, err
}
func float2TimeStr(f float64) string {
	//2021-07-13T17:28:55.095708+08:00
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(f)
	t, _ := strconv.ParseInt(newNum.FloatString(0), 10, 64)
	if t == 0 {
		return ""
	} else {
		return time.Unix(t, 0).Format("2006-01-02T15:04:05") + ".095708+08:00"
	}
}
func Scans(args *ScanReq) error {
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "post"
	req.Url += scan_url + "/" + args.ID + "/launch"

	req.Params = map[string]interface{}{
		"limit":                              2500,
		"includeHostDetailsForHostDiscovery": true,
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func Pause(args *PauseReq) error {
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "post"
	req.Url += scan_url + "/" + args.ID + "/stop"

	req.Params = map[string]interface{}{
		"limit":                              2500,
		"includeHostDetailsForHostDiscovery": true,
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func Resume(args *ResumeReq) error {
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "post"
	req.Url += scan_url + "/" + args.ID + "/resume"

	req.Params = map[string]interface{}{
		"limit":                              2500,
		"includeHostDetailsForHostDiscovery": true,
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func Export(args *ExportReq) (*ExportResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "post"
	req.Url += scan_url + "/" + args.ID + "/export?limit=2500"
	if args.HistoryId != "" {
		req.Url += "&history_id=" + args.HistoryId
	}
	args.Format = strings.ToLower(args.Format)
	var reportContents interface{}
	if args.Format == "csv" {
		reportContents = map[string]interface{}{
			"csvColumns": map[string]interface{}{
				"cve":                  true,
				"cvss":                 true,
				"cvss3_base_score":     false,
				"cvss3_temporal_score": false,
				"cvss_temporal_score":  false,
				"description":          true,
				"exploitable_with":     false,
				"hostname":             true,
				"id":                   true,
				"plugin_information":   false,
				"plugin_name":          true,
				"plugin_output":        true,
				"port":                 true,
				"protocol":             true,
				"references":           false,
				"risk":                 true,
				"risk_factor":          false,
				"see_also":             true,
				"solution":             true,
				"stig_severity":        false,
				"synopsis":             true,
			},
		}

	} else if args.Format == "html" {
		reportContents = map[string]interface{}{
			"csvColumns":            map[string]interface{}{},
			"formattingOptions":     map[string]interface{}{},
			"hostSections":          map[string]interface{}{},
			"vulnerabilitySections": map[string]interface{}{},
		}
	}
	req.Params = map[string]interface{}{
		"format":   args.Format,
		"chapters": "",
		"extraFilters": map[string]interface{}{
			"host_ids":   []string{},
			"plugin_ids": []string{},
		},
		"reportContents": reportContents,
	}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}

	export := &ExportResp{}
	err = json.Unmarshal(resp, export)
	return export, err
}
func Vulnerabilities(args *VulnerabilitiesReq) ([]interface{}, error) {

	info, err := getScanInfo(args.ID, args.HistoryId)
	if err != nil {
		return nil, err
	}
	vuls := info["vulnerabilities"].([]interface{})
	list := []interface{}{}
	for _, v := range vuls {
		vul := v.(map[string]interface{})
		vul["name"] = vul["plugin_name"]
		vul["target_info"] = map[string]interface{}{"host": info["info"].(map[string]interface{})["name"]}
		vul["time"] = float2TimeStr(info["info"].(map[string]interface{})["timestamp"].(float64))
		list = append(list, vul)
	}
	return list, nil
}

//漏洞详情
func Plugins(args *PluginsReq) (map[string]interface{}, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += scan_url + "/" + args.ID + "/plugins/" + args.VulId + "?history_id=" + args.HistoryId + "&limit=2500"

	req.Params = map[string]interface{}{}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}

	info := map[string]interface{}{}
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	ret := map[string]interface{}{}
	ret["vt_name"] = info["info"].(map[string]interface{})["plugindescription"].(map[string]interface{})["pluginname"]
	ret["description"] = info["info"].(map[string]interface{})["plugindescription"].
	(map[string]interface{})["pluginattributes"].(map[string]interface{})["description"]
	ret["details"] = info["info"].(map[string]interface{})["plugindescription"].
	(map[string]interface{})["pluginattributes"].(map[string]interface{})["synopsis"]
	//ret["impact"] = info["info"].(map[string]interface{})["plugindescription"].
	//(map[string]interface{})["pluginattributes"].(map[string]interface{})["synopsis"]
	ret["recommendation"] = info["info"].(map[string]interface{})["plugindescription"].
	(map[string]interface{})["pluginattributes"].(map[string]interface{})["solution"]
	return ret, nil
}
func Delete(args *DeleteReq) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "put"
	req.Url += scan_url + "/" + args.ID + "/folder"

	folder_id, err := getScanFoldersId("Trash")
	if err != nil {
		return err
	}
	req.Params = map[string]interface{}{
		"folder_id": folder_id,
	}
	resp, err := req.Do()
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

//reset 重置扫描
func Reset(args *ResetReq) error {
	//先删除该扫描 然后创建
	//获取desc
	var desc, name string
	entiry, _ := GetInfo(args.ID)
	desc = entiry.Description
	info, err := getScanInfo(args.ID)
	if err != nil {
		fmt.Println("get scan info error ", err.Error())
	} else {
		name = info["info"].(map[string]interface{})["name"].(string)
	}
	err = Delete(&DeleteReq{ID: args.ID})
	if err != nil {
		return err
	}
	//创建
	req := &AddReq{}
	req.UserId = args.UserId
	req.AdminUserId = args.AdminUserId
	req.Settings.Name = name
	req.Settings.Text_targets = name
	req.Settings.Description = desc
	_, err = Create(req)
	if err != nil {
		fmt.Println("create scan error ", err)
	}
	return nil
}

func History(args *HistoryReq) ([]interface{}, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	folder_id, err := getScanFoldersId("My Scans")
	if err != nil {
		return nil, err
	}
	req.Method = "get"
	req.Params = map[string]interface{}{
		"folder_id": folder_id,
	}
	req.Url += scan_url
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	result := make(map[string]interface{}, 0)

	err = json.Unmarshal(ret, &result)
	if err != nil {
		return nil, err
	}
	scansList, total, err := GetList(&ScansListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return nil, err
	}
	contain := map[uint64]string{}
	for _, v := range scansList {
		contain[v.ScansId] = v.Description
	}

	retList := make([]interface{}, 0)

	for _, v := range result["scans"].([]interface{}) {
		scan := v.(map[string]interface{})
		id, _ := util.Interface2Uint64(scan["id"])
		//对应web扫描
		if desc, isExist := contain[id]; isExist {
			info, err := getScanInfo(fmt.Sprintf("%v", id))
			if err != nil {
				return nil, err
			}
			if info["history"] == nil {
				return nil, nil
			}
			for _, v := range info["history"].([]interface{}) {
				item := v.(map[string]interface{})
				node := map[string]interface{}{}
				node["target_id"] = fmt.Sprintf("%v-host", scan["id"])
				node["owner"] = "host"
				node["scan_id"] = fmt.Sprintf("%v-host", item["history_id"])
				node["target"] = map[string]interface{}{
					"address":     scan["name"],
					"description": desc,
				}
				node["current_session"] = map[string]interface{}{
					"status":     item["status"],
					"start_date": float2TimeStr(item["creation_date"].(float64)),
				}
				history_info, err := getScanInfo(fmt.Sprintf("%v", scan["id"]), fmt.Sprintf("%v", item["history_id"]))
				if err != nil {
					return nil, err
				}
				if len(history_info["hosts"].([]interface{})) > 0 {
					severity_counts := map[string]float64{}
					severity_counts["high"] = history_info["hosts"].([]interface{})[0].(map[string]interface{})["high"].(float64) + history_info["hosts"].([]interface{})[0].(map[string]interface{})["critical"].(float64)
					severity_counts["medium"] = history_info["hosts"].([]interface{})[0].(map[string]interface{})["medium"].(float64)
					severity_counts["low"] = history_info["hosts"].([]interface{})[0].(map[string]interface{})["low"].(float64)
					severity_counts["info"] = history_info["hosts"].([]interface{})[0].(map[string]interface{})["info"].(float64)

					node["current_session"].(map[string]interface{})["severity_counts"] = severity_counts
				}
				retList = append(retList, node)
			}
		}
	}
	return retList, nil
}

func DelHistory(args *DelHistoryReq) error {
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "delete"
	req.Url += scan_url + "/" + args.ID + "/history/" + args.HistoryId

	resp, err := req.Do()
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	//删除对应报表
	_ = DeleteScansReportByHistoryId(args.ID, args.HistoryId)
	return nil
}

func CreateReport(args *CreateReportReq) error {

	info, err := GetInfo(args.ID)
	if err != nil {
		return err
	}
	report := &NessusScanReport{}
	report.HistoryId, _ = util.Interface2Uint64(args.HistoryId)
	report.ScansId, _ = util.Interface2Uint64(args.ID)
	report.UserId = args.UserId
	report.AdminUserId = args.AdminUserId
	report.Addr = info.Addr
	report.CreateTime = time.Now().Unix()

	return AddScansReport(report)
}
func ListReport(args *ScansListReq) ([]interface{}, error) {
	list, total, err := GetListReport(args)
	if err != nil || total == 0 {
		return nil, err
	}
	ret := make([]interface{}, 0)
	for _, v := range list {
		node := map[string]interface{}{}
		node["report_id"] = fmt.Sprintf("%v-host", v.Id)
		node["scan_id"] = fmt.Sprintf("%v", v.ScansId)
		node["history_id"] = fmt.Sprintf("%v", v.HistoryId)
		node["address"] = v.Addr
		node["generation_date"] = time.Unix(int64(v.CreateTime), 0).Format("2006-01-02T15:04:05") + ".095708+08:00"
		node["status"] = "completed"
		node["owner"] = "host"
		ret = append(ret, node)
	}
	return ret, nil
}

func DeleteReport(ids []string) (err error) {
	return DeleteScansReportById(ids)
}
