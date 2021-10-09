package wechat
//微信基础配置结构体定义

const SignTypeMd5 = "MD5" //签名类型 md5
const SignTypeSha256 = "HMAC-SHA256" //签名类型SHA256

//核心配置
type WechatCoreConfig struct {
	AppId string  //微信分配的ID
	Secret string //秘钥
	MchId string //商户号
	AppKey string //key
	NotifyUrl string //异步回调地址
}