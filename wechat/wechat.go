package wechat

import (
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"gopay/lib"
	"gopay/lib/config"
	"gopay/wechat/collect"
	"gopay/wechat/mini"
	"reflect"
	"sort"
	"strings"
)

const SignTypeMd5 = "MD5"            //签名类型 md5
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
func (bc *Client) Mini() collect.IEntity {
	return mini.NewImplement(bc.CoreConfig)
}

//签名校验
func (bc *Client) VerifySign(params interface{}, appKey string, signType string) bool {
	valMap := gmap.StrAnyMap{}
	typeOf := reflect.TypeOf(params).Elem()
	valueOf := reflect.ValueOf(params).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		key := typeOf.Field(i).Tag.Get("xml")
		val := valueOf.Field(i).Interface()
		valMap.Set(key, val)
	}
	originalSign := valMap.Get("sign")
	valMap.Remove("sign")
	//字典排序
	keys := valMap.Keys()
	sort.Strings(keys)
	var uri []string
	for _, item := range keys {
		if valMap.Get(item) != "" {
			uri = append(uri, fmt.Sprintf("%v=%v", item, valMap.Get(item)))
		}
	}
	uriStr := strings.Join(uri, "&")
	//追加上key
	uriStr += fmt.Sprintf("&key=%s", appKey)
	var newString string
	if signType == SignTypeMd5 {
		newString = lib.Md5(uriStr)
	} else {
		newString = lib.Hmac256(uriStr, appKey)
	}
	return strings.ToUpper(newString) == originalSign
}
