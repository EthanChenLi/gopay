package wechat

//方法接口集合

type IEntity interface {
	Unifiedorder(IParams, map[string]interface{}) //统一下单接口
}

//IParams 入参参数
type IParams interface {}