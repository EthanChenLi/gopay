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

// UnifiedOrder 统一下单接口
// params :
//	OutTradeNo <string> 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一
//	Body <string> 商品简单描述，该字段请按照规范传递
//	TotalFee <int> 订单总金额，单位为分
// Openid <string> openid
// exp 补充参数
func (mini *Implement)UnifiedOrder(params collect.IParams, exp map[string]interface{}) (*collect.ResponseWithMiniOfUnifiedOrder, error) {
	mini.Params = params
	mini.Exp = exp
	param , err := mini.unifiedOrderService()
	if err !=nil {
		return nil,err
	}
	return param, nil
}

// OrderQuery 订单查询
// 该接口提供所有微信支付订单的查询，商户可以通过查询订单接口主动查询订单状态，完成下一步的业务逻辑。
func (mini *Implement)OrderQuery(params collect.IParams) (*collect.ResponseWithMiniOfOrderQuery, error){
	mini.Params = params
	return mini.orderQueryService()
}

//CloseOrder 关闭订单
//out_trade_no 订单编号
func (mini *Implement)CloseOrder(params collect.IParams) (*collect.ResponseWithMiniOfCloseOrder, error)  {
	mini.Params = params
	return mini.closeOrderService()
}