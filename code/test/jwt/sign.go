package jwt

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"
)

//func (u *ClientRouter) Initialize(app *gin.Engine) {
//	prefix := app.Group("v1")
//	{
//		apps := prefix.Group("client").
//			Use(middlewares.Sign())
//		{
//			// 这里放置路由
//		}
//	}
//}
//func Sign() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var req request.Sign
//		validate := validator.NewValidator()
//		if errMsg := validate.ParseHeader(c, &req); errMsg != "" {
//			utils.MessageError(c, errMsg)
//			c.Abort()
//			return
//		}
//		//获取数据
//		rawDataBody, _ := c.Get("rawData")
//		rawData := rawDataBody.([]byte)
//		err := utils.CheckSign(rawData, core.AppConfig.BasicConfig.SignKey, req.Timestamp, req.Sign)
//		if err != nil {
//			utils.Error(c, &utils.CodeType{Code: utils.SignError.Code, Msg: err.Error()})
//			c.Abort()
//			return
//		}
//	}
//}

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
