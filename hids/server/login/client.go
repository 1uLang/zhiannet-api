package login

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"time"
)

//获取请求客户端
func GetHttpClient(req *ApiKey) *resty.Client {

	Client := resty.New().SetDebug(false).SetTimeout(time.Second * 60)
	if req.IsSsl {
		Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return Client
}
