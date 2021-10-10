package wechat

import (
	"gopay/lib/config"
	"gopay/wechat/collect"
	"gopay/wechat/mini"
)

const SignTypeMd5 = "MD5" //签名类型 md5
const SignTypeSha256 = "HMAC-SHA256" //签名类型SHA256




// Client 微信基础配置
type Client struct {
	CoreConfig config.WechatCoreConfig
}


// NewClient  初始化定义
func NewClient(coreConfig config.WechatCoreConfig) *Client {
	return &Client{
		CoreConfig: coreConfig,
	}
}


//Mini 小程序
func (bc *Client)Mini() collect.IEntity{
	return mini.NewImplement(bc.CoreConfig)
}