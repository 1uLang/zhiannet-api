package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}

//攻击日志
func Test_attack_log_list(t *testing.T) {
	InitDB()
	list, err := GetAttackLogList(&AttackLogReq{
		NodeId: 6,
		Addr:   "182.150.0.37",
	})
	fmt.Println(list)
	fmt.Println(err)

}

//流量日志
func Test_traffic_log_list(t *testing.T) {

	InitDB()
	list, err := GetTrafficLogList(&TrafficLogReq{
		NodeId: 6,
		Addr:   "182.150.0.37",
		Level:  3,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//链接日志
func Test_link_log_list(t *testing.T) {

	InitDB()
	list, err := GetLinkLogList(&LinkLogReq{
		NodeId: 1,
		Addr:   "182.150.0.37",
		Level:  3,
	})
	fmt.Println(list)
	fmt.Println(err)
}

type req struct {
	Amount      int     `json:"amount"`
	Description string  `json:"description"`
	FCurrency   int     `json:"fCurrency"`
	Hotel       []hotel `json:"hotel"`
	MemberId    int     `json:"memberId"`
	NonceStr    string  `json:"nonce_str"`
	NotifyUrl   string  `json:"notify_url"`
	OrderId     string  `json:"order_id"`
	PayType     int     `json:"pay_type"`
	SignType    string  `json:"sign_type"`
	TimeExpire  string  `json:"time_expire"`
	TimeStamp   string  `json:"time_stamp"`
	ApiKey      string  `json:"api_key"`
}

type hotel struct {
	Id   string
	Name string
}

//amount=1&description=商品描述&fCurrency=0&hotel[0].id=2&hotel[0].name=济南旗舰店&memberId=202106010000005&nonceStr=hZr3uYblmsKGDcuJRHBq02Cg&notifyUrl=http://api-dev.smartgo.fun:1102/api/notify&orderId=bp21062410000011&payType=1&signType=MD5&timeExpire=6/24/2021 3:00:03 PM&timeStamp=1624517103&apiKey=39fd20a6a49ea3b6-fcdc412c6e919b35
func Test_to_map(t *testing.T) {
	ho := []hotel{
		{Id: "2", Name: "济南旗舰店"},
		{Id: "2", Name: "济南旗舰店"},
	}
	req := req{
		Amount:      1,
		Description: "商品描述",
		FCurrency:   0,
		Hotel:       ho,
		MemberId:    202106010000005,
		NonceStr:    "hZr3uYblmsKGDcuJRHBq02Cg",
		NotifyUrl:   "http://api-dev.smartgo.fun:1102/api/notify",
		OrderId:     "bp21062410000011",
		PayType:     1,
		SignType:    "MD5",
		TimeExpire:  "6/24/2021 3:00:03 PM",
		TimeStamp:   "1624517103",
		ApiKey:      "39fd20a6a49ea3b6-fcdc412c6e919b35",
	}

	queryMap := map[string]interface{}{}
	queryMap["amount"] = req.Amount
	queryMap["description"] = req.Description
	queryMap["fCurrency"] = req.FCurrency
	queryMap["memberId"] = req.MemberId
	queryMap["nonceStr"] = req.NonceStr
	queryMap["notifyUrl"] = req.NotifyUrl
	queryMap["orderId"] = req.OrderId
	queryMap["payType"] = req.PayType
	queryMap["signType"] = req.SignType
	queryMap["timeExpire"] = req.TimeExpire
	queryMap["timeStamp"] = req.TimeStamp
	queryMap["apiKey"] = req.ApiKey
	if len(req.Hotel) > 0 {
		for k, v := range req.Hotel {
			keyId := fmt.Sprintf("hotel[%v].id", k)
			keyName := fmt.Sprintf("hotel[%v].name", k)
			queryMap[keyId] = v.Id
			queryMap[keyName] = v.Name
		}
	}
	fmt.Println(queryMap)
}
