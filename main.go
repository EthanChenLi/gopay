package main

import (
	"gopay/src/wechat"
	"log"
)

func main() {
	log.Println(wechat.SignTypeMd5)
	/* _ = wechat.NewConfig(wechat.WechatCoreConfig{
		AppId:  "wx0fbf65a1ba507cf0",
		Secret: "b8f1c66819e729a66306ab3690419e76",
		MchId:  "1614403327",
		AppKey: "64711ac8dd5e73106383811b519dd9a6",
		NotifyUrl: "",
	})*/
	//client.Mini().Unifiedorder(mini.Params{
	//	OutTradeNo: "12312",
	//	Body:       "支付",
	//	TotalFee:   1,
	//	SignType: wechat.SignTypeMd5,
	//}, nil)
}
