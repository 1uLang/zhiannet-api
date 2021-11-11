package attack_message_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_message_media_instances"
	"github.com/1uLang/zhiannet-api/common/model/edge_message_recipients"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"github.com/1uLang/zhiannet-api/common/util"
	agents_model "github.com/1uLang/zhiannet-api/wazuh/model/agents"
	wazuh_server "github.com/1uLang/zhiannet-api/wazuh/server"
	"strconv"
	"strings"
	"time"
)

type wazuh struct{}

func (wazuh) AttackCheck(interval time.Duration) error {
	host, user, password, from, port, emailId := getEmailInfo()
	fmt.Println("获取管理端配置发送者邮箱：", host, user, password, from, port, emailId)
	if host == "" {
		return nil
	}
	info, err := wazuh_server.GetWazuhInfo()
	if err != nil {
		return err
	}
	err = wazuh_server.SetUrl(info.Addr)
	if err != nil {
		return err
	}
	err = wazuh_server.InitToken(info.Username, info.Password)
	if err != nil {
		return err
	}

	// Flag 0000 0 1 1 1
	//   暴力破解|病毒|文件篡改

	//文件篡改
	agents := map[string]uint8{}
	now := time.Now()
	filecomList, err := wazuh_server.SysCheckESList(agents_model.ESListReq{Start: now.Add(-interval).Unix(), End: now.Unix(), Limit: 10000})
	fmt.Println("文件篡改 ： ", len(filecomList.Hits))
	if err != nil {
		fmt.Println("文件完整性列表获取失败")
	} else {
		for _, c := range filecomList.Hits {
			if c.Source.Agent.Id == "000" { //wazuh agent 忽略
				continue
			}
			//添加|删除|修改
			if agents[c.Source.Agent.Id]&0x01 == 0 &&
				(c.Source.Rule.Id == "554" || c.Source.Rule.Id == "553" || c.Source.Rule.Id == "550") {
				agents[c.Source.Agent.Id] = 0x01
				fmt.Println("==========文件篡改", agents[c.Source.Agent.Id])
			}
		}
	}
	//病毒检测
	riskList, err := wazuh_server.VirusList(agents_model.ESListReq{Start: now.Add(-interval).Unix(), End: now.Unix(), Limit: 10000})
	fmt.Println("病毒检测 ： ", len(riskList.Hits))
	if err != nil {
		fmt.Println("文件完整性列表获取失败")
	} else {
		for _, c := range riskList.Hits {
			if c.Source.Agent.Id == "000" { //wazuh agent 忽略
				continue
			}
			//病毒
			if agents[c.Source.Agent.Id]&0x02 == 0 && c.Source.Rule.Id == "87105" {
				agents[c.Source.Agent.Id] = agents[c.Source.Agent.Id] | 0x02

				fmt.Println("==========病毒检测", agents[c.Source.Agent.Id])
			}
		}
	}
	//暴力破解
	attList, err := wazuh_server.ATTCKESList(agents_model.ESListReq{Start: now.Add(-interval).Unix(), End: now.Unix(), Limit: 10000})
	fmt.Println("暴力破解 ： ", len(attList.Hits))
	if err != nil {
		fmt.Println("文件完整性列表获取失败")
	} else {
		for _, c := range attList.Hits {
			if c.Source.Agent.Id == "000" { //wazuh agent 忽略
				continue
			}
			//病毒 ssh: 5700 - 5759 rdp : 9500 - 9505 9510 9551
			ruleId, _ := strconv.Atoi(c.Source.Rule.Id)
			if agents[c.Source.Agent.Id]&0x04 == 0 && (ruleId >= 5700 && ruleId <= 5759 ||
				ruleId >= 9500 && ruleId <= 9505 || ruleId == 9510 || ruleId == 9551) {
				agents[c.Source.Agent.Id] = agents[c.Source.Agent.Id] | 0x04
				fmt.Println("==========暴力破解", agents[c.Source.Agent.Id])
			}
		}
	}
	users := map[string]string{}
	admins := map[string]string{}
	fmt.Println("send email : ", agents)
	for id, flag := range agents {
		info, err := wazuh_server.AgentInfo(id)
		if err != nil {
			fmt.Printf("获取wazuh agent[%s] 信息失败：%s", id, err.Error())
		} else {
			//获取agent 所属用户信息
			if len(info.Affected_items) == 0 {
				continue
			}
			group := info.Affected_items[0].Group[0]
			var email string
			var isExist bool
			if strings.HasPrefix(group, "admin_") { //admin
				id := strings.TrimPrefix(group, "admin_")
				email, isExist = admins[id]
				if !isExist {
					em, err := getAdminEmail(id, emailId)
					if err != nil {
						fmt.Printf("获取管理员[%s] 邮箱信息失败：%s", id, err.Error())
						continue
					}
					email = em
					admins[id] = em
				}
			} else if strings.HasPrefix(group, "user_") { //admin
				id := strings.TrimPrefix(group, "user_")
				email, isExist = users[id]
				if !isExist {
					em, err := getUserEmail(id)
					if err != nil {
						fmt.Printf("获取平台用户[%s] 邮箱信息失败：%s", id, err.Error())
						continue
					}
					email = em
					users[id] = em
				}
			} else {
				continue
			}
			if flag&0x01 > 0 { //文件篡改
				_ = sendEmail(host, user, password, from, email, fmt.Sprintf("请注意！检测到您的[%s](%s)主机存在文件篡改入侵事件。如需查询详情，请前往智安云综合防御平台。",
					info.Affected_items[0].Name, info.Affected_items[0].IP), port)
			}
			if flag&0x02 > 0 { //病毒检测
				_ = sendEmail(host, user, password, from, email, fmt.Sprintf("请注意！检测到您的[%s](%s)主机存在病毒入侵事件。如需查询详情，请前往智安云综合防御平台。",
					info.Affected_items[0].Name, info.Affected_items[0].IP), port)
			}
			if flag&0x04 > 0 { //暴力破解
				_ = sendEmail(host, user, password, from, email, fmt.Sprintf("请注意！检测到您的[%s](%s)主机存在暴力破解入侵事件。如需查询详情，请前往智安云综合防御平台。",
					info.Affected_items[0].Name, info.Affected_items[0].IP), port)
			}
		}
	}
	return nil
}
func getEmailInfo() (host, user, password, from string, port int, id uint32) {
	info, _ := edge_message_media_instances.MessageMediaInstance{}.GetEmail()
	return info.Smtp, info.Username, info.Password, info.From, 465, info.Id
}

//发送邮件
func sendEmail(host, user, password, from, email, content string, port int) error {
	return util.SendEmail(host, user, password, from, "243971996@qq.com", content, port)
}

//获取用户邮箱
func getUserEmail(idstr string) (string, error) {
	id, _ := strconv.ParseUint(idstr, 10, 64)
	user, err := edge_users.GetInfoById(id)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

//获取管理员邮箱
func getAdminEmail(idstr string, emailId uint32) (string, error) {

	id, _ := strconv.ParseUint(idstr, 10, 32)
	edge_message_recipients.MessageRecipient{AdminId: uint32(id), InstanceId: emailId}.GetEmail()
	return "", nil
}
