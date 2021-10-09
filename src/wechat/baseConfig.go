package wechat

import "gopay/src/wechat/mini"

//初始化定义

// BaseConfig 微信基础配置
type BaseConfig struct {
	CoreConfig WechatCoreConfig
}

func NewConfig(baseWechatConfig WechatCoreConfig) *BaseConfig {
	return &BaseConfig{
		CoreConfig: baseWechatConfig,
	}
}

//Mini 小程序
func (bc *BaseConfig)Mini() IEntity  {
	return mini.NewImplement(bc)
}
