package mini

import (
	"gopay/src/wechat"
	"log"
)

type Params struct {
	OutTradeNo string
	Body string
	TotalFee int
	SignType string
}

//统一下单
func (mini *Implement)unifiedorderService()  {
	mini.buildParams()
}

func (mini *Implement)buildParams () {

	params := map[string]interface{}{}
	miniParams := mini.Params.(Params)
	params["appid"] = mini.BaseConfig.CoreConfig.AppId
	params["mch_id"] = mini.BaseConfig.CoreConfig.MchId
	params["nonce_str"] = wechat.GetRandomString(32)
	params["sign_type"] = miniParams.SignType
	params["body"] = miniParams.Body
	params["out_trade_no"] = miniParams.OutTradeNo
	params["total_fee"] = miniParams.TotalFee
	params["spbill_create_ip"] = wechat.GetClientIp()
	params["notify_url"] = mini.BaseConfig.CoreConfig.NotifyUrl
	params["trade_type"] = tradeType
	for key, item := range mini.Exp {
		params[key] = item
	}
	log.Println(params)
}