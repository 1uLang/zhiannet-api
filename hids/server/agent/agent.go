package agent

import (
	"github.com/1uLang/zhiannet-api/hids/model/agent"
)

func Create(args *agent.CreateReq) (err error) {
	return agent.Create(args)
}

func Download(osType string) (string, error) {
	return agent.Download(osType)
}
func Install(osType string) (string, error) {
	return agent.Install(osType)
}
func List(req *agent.SearchReq) (agent.SearchResp, error) {
	return agent.List(req)
}
func ListAll() (agent.SearchResp, error) {
	return agent.ListAll(&agent.SearchReq{})
}
func Disport(macCode, opt string) error {
	return agent.Disport(macCode, opt)
}
func Delete(req *agent.DeleteReq) error {
	return agent.Delete(req)
}
func GetUserListByAgentIP(ip string)  (users,admins []uint64, err error) {
	list,err := agent.GetUserListByAgentIP(ip)
	if err != nil {
		return nil,nil, err
	}
	for _,v := range list{
		if v.AdminUserId != 0 {
			admins = append(admins, v.AdminUserId)
		}
		if v.UserId != 0 {
			users = append(users, v.UserId)
		}
	}
	return users,admins,nil
}