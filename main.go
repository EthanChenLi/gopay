package main

import (
	"gopay/lib/config"
	"gopay/wechat"
	"gopay/wechat/mini"
)

func main() {
	client := wechat.NewClient(config.WechatCoreConfig{
		AppId:     "wx0fbf65a1ba507cf0",
		Secret:    "b8f1c66819e729a66306ab3690419e76",
		MchId:     "1614403327",
		AppKey:    "64711ac8dd5e73106383811b519dd9a6",
		NotifyUrl: "https://baidu.com",
	})
	//统一下单
	/*params, _ := client.Mini().UnifiedOrder(mini.UnifiedOrderParams{
		OutTradeNo: "123123123123123",
		Body:       "支付",
		TotalFee:   1,
		SignType: wechat.SignTypeMd5,
		Openid: "oKgMU5Fc3EN0Z76fPIdPt0XScnoA",
	}, nil)*/
	//查询订单
	params, _ := client.Mini().OrderQuery(mini.OrderQueryParams{
		OutTradeNo: "202110111507087483",
	})
	//参数校验
	client.VerifySign(params, client.CoreConfig.AppKey, wechat.SignTypeMd5)

	//关闭订单
	//params, _ :=client.Mini().CloseOrder(mini.CloseOrderParams{OutTradeNo: "202110111507087483"})

}
