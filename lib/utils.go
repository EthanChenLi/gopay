package lib

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
	"unsafe"
)

//随机生成字符串
func GetRandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte = make([]byte, 0, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return BytesToString(result)
}

//随机生成纯字符串
func GetRandomPureString(l int) string {
	str := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte = make([]byte, 0, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return BytesToString(result)
}

//随机生成数字字符串
func GetRandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return BytesToString(result)
}

// BytesToString 0 拷贝转换 slice byte 为 string
func BytesToString(b []byte) (s string) {
	_bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	_sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	_sptr.Data = _bptr.Data
	_sptr.Len = _bptr.Len
	return s
}

func GetClientIp() string {
	return "127.0.0.1"
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Hmac256(str, apiKey string) string {
	h := hmac.New(sha256.New, []byte(apiKey))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//签名创建
//https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=4_3
//参数名ASCII码从小到大排序（字典序）；
//◆ 如果参数的值为空不参与签名；
//◆ 参数名区分大小写；
//◆ 验证调用返回或微信主动通知签名时，传送的sign参数不参与签名，将生成的签名与该sign值作校验。
//◆ 微信接口可能增加字段，验证签名时必须支持增加的扩展字段

func WechatSignCreated(params *gmap.StrAnyMap, key string) string {
	//排序sign参数
	params.Removes([]string{"sign"})
	//字典排序
	keys := params.Keys()
	sort.Strings(keys)
	var uri []string
	for _, item := range keys {
		if params.Get(item) != "" {
			uri = append(uri, fmt.Sprintf("%v=%v", item, params.Get(item)))
		}
	}
	uriStr := strings.Join(uri, "&")
	//追加上key
	uriStr += fmt.Sprintf("&key=%s", key)
	var encryptionString string
	if params.Get("sign_type") == "MD5" {
		encryptionString = Md5(uriStr)
	} else {
		encryptionString = Hmac256(uriStr, key)
	}
	return strings.ToUpper(encryptionString)
}

func HttpPost(url string, xmlString string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(xmlString)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

//map转xml
func Map2Xml(params map[string]interface{}) string {
	strBuf := bytes.Buffer{}
	strBuf.WriteString("<xml>")
	for key, item := range params {
		strBuf.WriteString(fmt.Sprintf("<%s><![CDATA[%v]]></%s>", key, item, key))
	}
	strBuf.WriteString("</xml>")
	return strBuf.String()
}

//三目运算符
func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	} else {
		return f
	}
}
