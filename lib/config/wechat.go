package config

//WechatCoreConfig 核心配置
type WechatCoreConfig struct {
	AppId string  //微信分配的ID
	Secret string //秘钥
	MchId string //商户号
	AppKey string //key
	NotifyUrl string //异步回调地址
}
