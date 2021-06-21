package cron

import (
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
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&PortCheck{}))
	c.AddJob("0 */5 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&CodeCheck{}))

	c.Start()
}
