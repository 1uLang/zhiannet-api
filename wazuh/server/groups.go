package server

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/groups"
	"github.com/1uLang/zhiannet-api/wazuh/request"
)

func CreateGroup(name string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	return groups.Create(req, name)
}
