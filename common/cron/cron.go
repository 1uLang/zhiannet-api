package cron

import (
	audit_request "github.com/1uLang/zhiannet-api/audit/request"
	ddos_request "github.com/1uLang/zhiannet-api/ddos/request"
	monitor_cron "github.com/1uLang/zhiannet-api/monitor/cron"
	awvs_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	hids_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	nessus_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	nextcloud_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	opnsense_request "github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/robfig/cron/v3"
)

func InitCron() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	c := cron.New(cron.WithParser(parser))

	//c.AddFunc("*/1 * * * * *", func() {
	//	fmt.Println("tick every 1 second")
	//	time.Sleep(time.Second * 5)
	//})
	//运行结构体中的Run 方法,每次任务之间不重复 每5分钟执行一次
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&monitor_cron.PortCheck{}))
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&monitor_cron.CodeCheck{}))
	//安全审计组件定时检测状态是否可用
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&audit_request.LoginReq{}))
	//ddos金盾 组件定时检测状态是否可用
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&ddos_request.LoginReq{}))
	//下一代防火墙 组件定时检测状态是否可用
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&opnsense_request.ApiKey{}))
	//数据备份系统 组件定时检测状态是否可用
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&nextcloud_request.CheckRequest{}))
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&awvs_request.CheckRequest{}))
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&nessus_request.CheckRequest{}))
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&hids_request.CheckRequest{}))

	c.Start()
}
