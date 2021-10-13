package mini

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"gopay/lib"
	"gopay/wechat/collect"
	"strconv"
	"time"
)

type UnifiedOrderParams struct {
	OutTradeNo string
	Body string
	TotalFee int
	SignType string
	Openid string
}

type responseUnifiedOrder struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
	ErrCode string `xml:"err_code"`
	MchId string `xml:"mch_id"`
	AppId string `xml:"app_id"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	PrepayId string `xml:"prepay_id"`
}

const unifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
const Success = "SUCCESS"
const Fail = "FAIL"

//统一下单
func (mini *Implement)unifiedOrderService() (*collect.ResponseWithMiniOfUnifiedOrder, error) {
	miniParams := mini.Params.(UnifiedOrderParams)
	params :=  mini.buildParamsWithUnifiedOrder(miniParams)
	body , err := lib.HttpPost(unifiedOrderUrl, lib.Map2Xml(params.Map()))
	if err != nil {
		return nil, err
	}
	var resp *responseUnifiedOrder
	if err := xml.Unmarshal(body, &resp); err !=nil {
		return nil, err
	}
	//判断响应结果
	if resp.ReturnCode == Fail{
		return nil, errors.New(resp.ReturnMsg)
	}
	if resp.ResultCode == Fail {
		return nil,  errors.New(resp.ErrCodeDes)
	}
	return mini.buildResponseParamsWithUnifiedOrder(resp, miniParams.SignType, mini.BaseConfig.AppKey), nil
}

//构建返回给微信的参数
func(mini *Implement)buildResponseParamsWithUnifiedOrder (data *responseUnifiedOrder ,signType ,AppKey string) *collect.ResponseWithMiniOfUnifiedOrder {
	var resp collect.ResponseWithMiniOfUnifiedOrder
	resp.TimeStamp = strconv.Itoa(int(time.Now().Unix()))
	resp.Package = fmt.Sprintf("prepay_id=%s",data.PrepayId)
	resp.NonceStr =  data.NonceStr
	resp.SignType = signType
	respParams := gmap.StrAnyMap{}
	respParams.Set("timeStamp" ,resp.TimeStamp)
	respParams.Set("nonceStr" ,resp.NonceStr)
	respParams.Set("package" ,resp.Package)
	respParams.Set("signType" ,resp.SignType)
	resp.PaySign = lib.WechatSignCreated(&respParams, AppKey)
	return &resp
}


//构建参数
func (mini *Implement)buildParamsWithUnifiedOrder(miniParams UnifiedOrderParams) gmap.StrAnyMap {
	params := gmap.StrAnyMap{}
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
	params.Set("openid", miniParams.Openid)
	for key, item := range mini.Exp {
		params.Set(key, item)
	}
	params.Set("sign", lib.WechatSignCreated(&params, mini.BaseConfig.AppKey))
	return params
}