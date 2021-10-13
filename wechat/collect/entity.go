package collect
const WechatSignTypeMd5 = "MD5" //签名类型 md5
const WechatSignTypeSha256 = "HMAC-SHA256" //签名类型SHA256

type IEntity interface {
	UnifiedOrder(IParams, map[string]interface{})  (*ResponseWithMiniOfUnifiedOrder, error) //统一下单接口
	OrderQuery(params IParams) (*ResponseWithMiniOfOrderQuery, error) //订单查询
	CloseOrder(params IParams) (*ResponseWithMiniOfCloseOrder, error) //关闭订单
}

//IParams 入参参数
type IParams interface {}




//-----小程序------
// ResponseWithMini 统一下单
type ResponseWithMiniOfUnifiedOrder struct {
	TimeStamp string `json:"timeStamp"`
	NonceStr string `json:"nonceStr"`
	Package string `json:"package"`
	SignType string `json:"signType"`
	PaySign string `json:"paySign"`
}

//ResponseWithMiniOfOrderQuery 查询订单
type ResponseWithMiniOfOrderQuery struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	MchId string `xml:"mch_id"`
	AppId string `xml:"app_id"`
	Openid string `xml:"openid"`
	IsSubscribe string `xml:"is_subscribe"`
	TradeType string `xml:"trade_type"`
	TradeState string `xml:"trade_state"`
	BankType string `xml:"bank_type"`
	TotalFee int `xml:"total_fee"`
	FeeType string `xml:"fee_type"`
	CashFee int `xml:"cash_fee"`
	CashFeeType string `xml:"cash_fee_type"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo string `xml:"out_trade_no"`
	Attach string `xml:"attach"`
	TimeEnd string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
}

//关闭订单
type ResponseWithMiniOfCloseOrder struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	Appid string `xml:"appid"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ResultMsg string `xml:"result_msg"`
}