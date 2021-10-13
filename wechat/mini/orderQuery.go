package mini

import (
	"encoding/xml"
	"github.com/gogf/gf/container/gmap"
	"gopay/lib"
	"gopay/wechat/collect"
)

//订单查询

//应用场景
//该接口提供所有微信支付订单的查询，商户可以通过查询订单接口主动查询订单状态，完成下一步的业务逻辑。
//
//需要调用查询接口的情况：
//
//◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知（查单实现可参考：支付回调和查单实现指引）；
//◆ 调用支付接口后，返回系统错误或未知交易状态情况；
//◆ 调用刷卡支付API，返回USERPAYING的状态；
//◆ 调用关单或撤销接口API之前，需确认支付状态；
//https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2

const orderQuery = "https://api.mch.weixin.qq.com/pay/orderquery"

//OrderQueryParams 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
type OrderQueryParams struct {
	OutTradeNo string //二选一
	TransactionId string //二选一
}


func (mini *Implement)orderQueryService() (*collect.ResponseWithMiniOfOrderQuery, error){
	params := mini.buildParamsWithQueryService(mini.Params.(OrderQueryParams))
	body , err := lib.HttpPost(orderQuery, lib.Map2Xml(params.Map()))
	if err != nil {
		return nil, err
	}
	var resp *collect.ResponseWithMiniOfOrderQuery
	if err := xml.Unmarshal(body, &resp); err !=nil {
		return nil, err
	}
	return resp, nil
}

func (mini *Implement)buildParamsWithQueryService(data OrderQueryParams) gmap.StrAnyMap {
	params := gmap.StrAnyMap{}
	params.Set("appid", mini.BaseConfig.AppId)
	params.Set("mch_id", mini.BaseConfig.MchId)
	params.Set("nonce_str", lib.GetRandomString(32))
	params.Set("sign_type", collect.WechatSignTypeMd5)
	if data.OutTradeNo != "" {
		params.Set("out_trade_no" , data.OutTradeNo)
	}else{
		params.Set("transaction_id", data.TransactionId)
	}
	params.Set("sign", lib.WechatSignCreated(&params, mini.BaseConfig.AppKey))
	return params
}
