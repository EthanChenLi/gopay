# gopay
支付SDK for golang

#微信支付
### 初始化配置
```go
client := wechat.NewClient(config.WechatCoreConfig{
		AppId:  "",
		Secret: "",
		MchId:  "",
		AppKey: "",
		NotifyUrl: "",
	})
```
## 微信小程序
### 统一下单
```go
    //params返回给小程序端支付
	params, err := client.Mini().Unifiedorder(mini.Params{
		OutTradeNo: "", //订单号
		Body:       "支付", 
		TotalFee:   1, //支付金额（单位分）
		SignType: wechat.SignTypeMd5, //支付MD5 、 HMAC-SHA256
		Openid: "", //用户的openid
	}, nil)
```
#### 参数说明：
```client.Mini().Unifiedorder()```第二个参数为```map[string]interface{}```类型，用于添加其他微信需要的参数  
参考： https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1

### 订单查询
```go
    params, err := client.Mini().OrderQuery(mini.OrderQueryParams{
		OutTradeNo:    "",//订单号， 二选一
 		TransactionId: "" //支付流水号， 二选一
	})
```

### 关闭订单
```go
params, err :=client.Mini().CloseOrder(mini.CloseOrderParams{
		OutTradeNo: "", //订单号
	})
```
### 签名校验（可用于所有返回参数校验)
```go
    //params 请求返回的结构体
	// appkey 微信的支付key
	// md5 设置的加密方式
	// @return bool 
	ok := client.VerifySign(params, client.CoreConfig.AppKey, wechat.SignTypeMd5)
```