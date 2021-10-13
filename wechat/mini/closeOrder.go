package mini

import (
	"encoding/xml"
	"github.com/gogf/gf/container/gmap"
	"gopay/lib"
	"gopay/wechat/collect"
)

//关闭订单
//应用场景
//以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
//注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。

const closeOrderUrl = "https://api.mch.weixin.qq.com/pay/closeorder"

type CloseOrderParams struct {
	OutTradeNo string //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
}


func (mini *Implement)closeOrderService() (*collect.ResponseWithMiniOfCloseOrder, error) {
	params := mini.buildParamsWithCloseOrder(mini.Params.(CloseOrderParams))
	resp, err := lib.HttpPost(closeOrderUrl, lib.Map2Xml(params.Map()))
	if err != nil {
		return nil, err
	}
	var data *collect.ResponseWithMiniOfCloseOrder
	if err = xml.Unmarshal(resp, &data);err != nil {
		return nil, err
	}
	return data, nil
}

func (mini *Implement)buildParamsWithCloseOrder(data CloseOrderParams) gmap.StrAnyMap {
	params := gmap.StrAnyMap{}
	params.Set("appid", mini.BaseConfig.AppId)
	params.Set("mch_id", mini.BaseConfig.MchId)
	params.Set("nonce_str", lib.GetRandomString(32))
	params.Set("sign_type", collect.WechatSignTypeMd5)
	params.Set("out_trade_no", data.OutTradeNo)
	params.Set("sign", lib.WechatSignCreated(&params, mini.BaseConfig.AppKey))
	return params
}