package attack_message_server

import (
	"fmt"
	"time"
)

//hids 主机防护告警

type AttackMessageRequest struct {
	Interval time.Duration
}

func (this *AttackMessageRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("wazuh 暴力破解、文件完整性、病毒邮箱告警----------------------------------------------", err)
		}
	}()
	//err := hids{}.AttackCheck()
	//if err != nil {
	//	fmt.Println("hids 入侵检测告警失败：", err)
	//}

	err := wazuh{}.AttackCheck(this.Interval)
	if err != nil {
		fmt.Println("wazuh 邮箱告警失败：", err)
	}
}
