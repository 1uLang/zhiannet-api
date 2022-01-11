package audit

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var (
	LogSubmitAddr = []string{"123.129.208.233:5045", "123.129.208.235:5045", "123.129.208.236:5045"}
	ApiDbPath     = "./build/configs/api_db.yaml"
)

type (
	AuditLogSubmitAddr struct {
		Audit struct {
			Addr string `yaml:"addr"`
		} `yaml:"audit"`
	}
)

func InitLogAddr() {
	var yamlFile []byte
	var err error
	conf := new(AuditLogSubmitAddr)
	yamlFile, err = ioutil.ReadFile(ApiDbPath)
	//yamlFile, err = ioutil.ReadFile("/Users/dp/zhian/zhiannet-edge-line/EdgeAdmin/build/configs/api_db.yaml")

	if err != nil {
		fmt.Println("err1=", err)
		return
		//panic(fmt.Errorf("zhiannet package redis link yamlFile.Get err #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err == nil {
		if addr := strings.Split(conf.Audit.Addr, ","); len(addr) > 1 && addr[0] != "" {
			LogSubmitAddr = addr
		}
	}
	fmt.Println("err2=", err)
}
