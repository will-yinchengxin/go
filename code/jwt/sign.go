package jwt

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
	"time"
	"errors"
)

//验签
func CheckSign(rawData []byte, singKey string, timestamp int64, sign string) error {
	if sign == "" {
		return errors.New("sign 不能为空")
	}
	// 过期时间可作为常量设置到配置文件中
	if time.Now().Unix()-timestamp > 30 {
		return errors.New("timestamp超过一定时长,请重新生成")
	}

	realSign, err := MakeSign(rawData, singKey, timestamp)
	if err != nil {
		return err
	}
	if sign != realSign {
		return errors.New("签名认证失败")
	}
	return nil
}
func MakeSign(rawdata []byte, singKey string, timestamp int64) (sign string, err error) {
	reqS := make(map[string]interface{})
	err = json.Unmarshal(rawdata, &reqS)
	if err != nil {
		return
	}

	//1.key进行按字母自然排序
	var keys []string
	for k := range reqS {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//2.key val组成字符串
	keyS := ""
	for _, key := range keys {
		if key == "sign" {
			continue
		}
		reqV := fmt.Sprintf("%v", reqS[key])
		if keyS == "" {
			keyS += key + "=" + "" + reqV
		} else {
			keyS += "&" + key + "=" + reqV
		}
	}
	keyS += "&time=" + fmt.Sprintf("%v", timestamp)
	//3.两次md5
	sign = fmt.Sprintf("%x", md5.Sum([]byte(keyS)))         //将[]byte转成16进制
	sign = fmt.Sprintf("%x", md5.Sum([]byte(sign+singKey))) //将[]byte转成16进制
	return
}
