package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	// post请求
	templateParams := map[string]string{
		"unionId": "5cc97b95-f9fd-4ba5-ba8a-f03ffc10bb8d",
	}
	dataJson, _ := json.Marshal(templateParams)
	code, res, err  := HttpJsonPost("http://webapi-fat.shadowcreator.com/100026/v1/dy/authenticationInfo", string(dataJson))
	fmt.Println(code, res, err)

	// get请求
	headers := map[string]string{
		"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEzMzksImV4cCI6MTYzNDY5NTk3OCwiaXNzIjoiU2hhZG93Q3JlYXRvci1Qb3J0YWxXZWJzaXRlIn0.assn4X2syfcDnVo_Bf6siHiav1jPY_-SuTGgF5Z_XLI",
	}
	code, res, err = HttpGet("http://127.0.0.1:8899/v1/account/accountInfo", WithHeaders(headers))
	fmt.Println(code, res, err)
}

func httpRequest(method string, url string, options ...*Option) (code int, content string, err error) {
	//start := time.Now()

	reqOpts := defaultRequestOptions() // 默认的请求选项
	fmt.Println(options)
	for _, opt := range options {      // 在reqOpts上应用通过options设置的选项
		opt.apply(reqOpts)
	}
	// 创建请求对象
	req, err := http.NewRequest(method, url, strings.NewReader(reqOpts.data))
	// 记录请求日志
	//defer func() {
	//	dur := int64(time.Since(start) / time.Second)
	//	if dur >= 500 {
	//		log.Println("Http"+method, url, "in", reqOpts.data, "out", content, "err", err, "dur/ms", dur)
	//	} else {
	//		log.Println("Http"+method, url, "in", reqOpts.data, "out", content, "err", err, "dur/ms", dur)
	//	}
	//}()
	defer req.Body.Close()

	// 设置请求头
	if len(reqOpts.headers) != 0 {
		for key, value := range reqOpts.headers {
			req.Header.Add(key, value)
		}
	}
	if err != nil {
		return
	}
	// 发起请求
	client := &http.Client{Timeout: reqOpts.timeout}
	resp, error := client.Do(req)
	if error != nil {
		return 0, "", error
	}
	// 解析响应
	defer resp.Body.Close()
	code = resp.StatusCode
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)

	return
}

// 发起GET请求
func HttpGet(url string, options ...*Option) (code int, content string, err error) {
	return httpRequest(http.MethodGet, url, options...)
}

// 发起POST请求，请求头指定Content-Type: application/json
func HttpJsonPost(url string, data string) (code int, content string, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	code, content, err = httpRequest(
		http.MethodPost, url, WithHeaders(headers), WithData(data))

	return
}

// 针对可选的HTTP请求配置项，模仿gRPC
// 使用的Options设计模式实现
type requestOption struct {
	timeout time.Duration
	data    string
	headers map[string]string
}

type Option struct {
	apply func(option *requestOption)
}

func defaultRequestOptions() *requestOption {
	return &requestOption{ // 默认请求选项
		timeout: 5 * time.Second,
		data:    "",
		headers: nil,
	}
}

func WithTimeout(timeout time.Duration) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.timeout = timeout
		},
	}
}

func WithHeaders(headers map[string]string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.headers = headers
		},
	}
}

func WithData(data string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.data = data
		},
	}
}
