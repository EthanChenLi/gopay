package collect

type IEntity interface {
	Unifiedorder(IParams, map[string]interface{})  error //统一下单接口
}

//IParams 入参参数
type IParams interface {}


type WechatResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	ErrCodeDes string `xml:"err_code_des"`
	ErrCode string `xml:"err_code"`
	MchId string `xml:"mch_id"`
	AppId string `xml:"app_id"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
}