package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

/*
TLS: 用于在两个通信应用程序之间提供保密性和数据完整性
*/

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,

		/*
			这里设置的是连接超时,不是请求超时,所以不要在 http.Transport种设置超时,否则会莫名出现 io timeout 错误
		*/
		Dial: func(network, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(network, addr, time.Second*2) // 设置连接超时时间
			if err != nil {
				return nil, err
			}
			err = conn.SetDeadline(time.Now().Add(time.Second * 3)) // 设置发送接受数据超时
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func TestFunc() {
	for {
		res, err := Get("http://baidu.com/")
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res)
		time.Sleep(time.Second*1)
	}
}

func Get(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, _ := http.NewRequest("Get", url, body)
	req.Header.Add("content-type", "application/json")
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

