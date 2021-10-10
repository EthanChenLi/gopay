package mini

import (
	"gopay/lib/config"
	"gopay/wechat/collect"
)

//小程序支付方法集合
const tradeType = "JSAPI" //小程序取值如下：JSAPI

// Implement  小程序支付
type Implement struct {
	BaseConfig config.WechatCoreConfig
	Params     collect.IParams
	Exp        map[string]interface{}
}

func NewImplement(baseConfig config.WechatCoreConfig) *Implement {
	return &Implement{
		BaseConfig: baseConfig,
	}
}

// Unifiedorder 统一下单接口
// params :
//	OutTradeNo <string> 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一
//	Body <string> 商品简单描述，该字段请按照规范传递
//	TotalFee <int> 订单总金额，单位为分
// exp 补充参数
func (mini *Implement)Unifiedorder(params collect.IParams, exp map[string]interface{}) error {
	mini.Params = params
	mini.Exp = exp
	err := mini.unifiedorderService()
	if err !=nil {
		return err
	}


	return nil
}
