package mini

import (
	"errors"
	"github.com/gogf/gf/container/gmap"
	"gopay/lib"
)

type Params struct {
	OutTradeNo string
	Body string
	TotalFee int
	SignType string
}

const unifiedorderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
const Success = "SUCCESS"
const Fail = "FAIL"

//统一下单
func (mini *Implement)unifiedorderService() error {
	params :=  mini.buildParams()
	resp , err := lib.HttpPost(unifiedorderUrl, lib.Map2Xml(params.Map()))
	if err != nil {
		return err
	}
	//判断响应结果
	if resp.ReturnCode == Fail{
		return errors.New(resp.ReturnMsg)
	}
	if resp.ResultCode == Fail {
		return errors.New(resp.ErrCodeDes)
	}
	//todo



	return nil
}

//构建参数
func (mini *Implement)buildParams() gmap.StrAnyMap {
	params := gmap.StrAnyMap{}
	miniParams := mini.Params.(Params)
	params.Set("appid", mini.BaseConfig.AppId)
	params.Set("mch_id", mini.BaseConfig.MchId)
	params.Set("nonce_str", lib.GetRandomString(32))
	params.Set("sign_type", miniParams.SignType)
	params.Set("body", miniParams.Body)
	params.Set("out_trade_no", miniParams.OutTradeNo)
	params.Set("total_fee", miniParams.TotalFee)
	params.Set("spbill_create_ip", lib.GetClientIp())
	params.Set("notify_url", mini.BaseConfig.NotifyUrl)
	params.Set("trade_type", tradeType)
	for key, item := range mini.Exp {
		params.Set(key, item)
	}
	params.Set("sign", lib.WechatSignCreated(&params, mini.BaseConfig.AppKey))
	return params
}