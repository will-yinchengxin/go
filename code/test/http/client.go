package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 阿里云短信发送： https://studygolang.com/articles/32385
func Client() {
	/*
		//请求参数
		templateParams := map[string]interface{}{
			"number": code,
		}
		templateJson, _ := json.Marshal(templateParams)
		//参数转换
		body := url.Values{}
		body.Add("mobile", sms.Phone)
		body.Add("templateId", templateId)
		body.Add("templateParams", string(templateJson))
		//创建请求
		req, _ := http.NewRequest("POST", core.AppConfig.SMS.Url+core.AppConfig.SMS.Send, strings.NewReader(body.Encode()))
		//设置header头
		req.Header.Set("appId", strconv.Itoa(consts.APP_ID))
		req.Header.Set("sk", core.AppConfig.SMS.Sk)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		//发起请求
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return errors.New("发送短信失败httpStatus:" + resp.Status)
		}
		resBody, _ := ioutil.ReadAll(resp.Body)
		var res map[string]interface{}
		err = json.Unmarshal(resBody, &res)
		if err != nil {
			return err
		}
		if res["code"].(float64) != 200 {
			err = errors.New(res["message"].(string))
			return
		}

		//发送成功 写入redis
		//获取redigo的常连接
		conn := s.Rds.DB.Get()
		//创建可以打点的YcRedis对象
		rdb := ycredis.New(conn).WithContext(ctx)
		defer rdb.Close()
		rdbKey := consts.SMS_REDIS_PREFIX + ":" + sms.Type + ":" + sms.Phone
		if _, err = rdb.Do("SETEX", rdbKey, core.AppConfig.SMS.Expire, code); err != nil {
			return err
		}
		_ = s.incr(sms.Phone, ctx)
		return
	*/

	c := http.DefaultClient
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// 设置请求头信息
	//request.Header.Set("","")
	if err != nil {
		return
	}
	response, err := c.Do(request)
	defer response.Body.Close()
	b, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))

	/*
	 response 转化为json
		mapData := make(map[string]interface{})
		errJson := json.Unmarshal(byteData, &mapData)
		if errJson != nil {
			return "", errors.New("http response is not valid"), ""
		}
		if _, ok := mapData["code"]; !ok {
			return "", errors.New("http response is not valid"), ""
		}
		if _, ok := mapData["data"]; !ok {
			return "", errors.New("http response is not valid"), ""
		}
		if code, ok := mapData["code"].(float64); !ok || code != 200 {
			return "", errors.New("http response code not 200"), ""
		}
		if _, ok := mapData["data"].(map[string]interface{}); !ok {
			return "", errors.New("http response data empty"), ""
		}
	*/
}

//累计发送次数
//func incr(phone string, ctx context.Context) (err error) {
//	defer func() {
//		if err != nil {
//			_ = core.Log.Error(logs.TraceFormatter{
//				Trace: logrus.Fields{
//					"error": err.Error(),
//				},
//			})
//		}
//	}()
//	conn := s.Rds.DB.Get()
//	rdb := ycredis.New(conn).WithContext(ctx)
//	defer rdb.Close()
//	rdbKey := consts.SMS_REST_TIMES_REDIS_PREFIX + ":" + phone
//	//校验key是否存在
//	exist, err := redis.Bool(rdb.Do("EXISTS", rdbKey))
//	if err != nil {
//		return
//	}
//	if !exist {
//		if _, err = rdb.Do("SETEX", rdbKey, utils.GetSecondsNowToNextDay(), 0); err != nil {
//			return
//		}
//	}
//	if _, err = rdb.Do("INCR", rdbKey); err != nil {
//		return
//	}
//	return
//}


//func (m *Client) Send(send Send) error {
//	b, _ := json.Marshal(send)
//	req, _ := http.NewRequest("POST", m.mailConfig.Url+"/mail/send", bytes.NewBuffer(b))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header["sk"] = []string{m.mailConfig.Sk}
//	req.Header["appId"] = []string{m.mailConfig.AppId}
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	var res response
//	err = json.Unmarshal(body, &res)
//	if err != nil {
//		return err
//	}
//	if res.Code != 200 {
//		return errors.New(res.Message)
//	}
//	return nil
//}