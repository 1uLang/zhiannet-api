package attack_message_server

import "fmt"

//hids 主机防护告警

type AttackMessageRequest struct{}

func (*AttackMessageRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hids-入侵事件告警----------------------------------------------", err)
		}
	}()

}
