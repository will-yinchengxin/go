package http

import (
	"fmt"
	"net"
)

func Net() {
	// parseIp 方法用于判断IP是否合法
	str := net.ParseIP("127.0.0.1")
	fmt.Println(str)
}
