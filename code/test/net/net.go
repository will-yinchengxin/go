package net

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func Conn() {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	//}
	fmt.Println("start conn")
	service := "127.0.0.1:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println(string(result))
	os.Exit(0) // go run main.go 127.0.0.1:8080
}
func Server() {
	service := "localhost:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		fmt.Println("start server")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // 不需要在意返回参数
		conn.Close()                // 结束客户端连接
	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

//  port 1 到 65535 的无符号整数
func ports() {
	// os.Exit 状态码 0 表示成功，非 0 表示出错
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network----type service \n", os.Args[0])
		os.Exit(0)
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(port) // go run main.go tcp mysql
	os.Exit(0)
}

func host() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s Ipaddr", os.Args[0])
		return
	}
	name := os.Args[1]

	//addr := net.ParseIP(name)
	//if addr == nil {
	//	fmt.Println("invalid addr")
	//	os.Exit(0)
	//}
	//mask := addr.DefaultMask()
	//network := addr.Mask(mask)
	//ones, bits := mask.Size()
	//fmt.Println("Addr is",
	//	addr.String(),
	//	"Default mask length is ", bits,
	//	"\n Leading ones count is ", ones,
	//	"\n Mask is (hex) ", mask.String(),
	//	" \n Network is ", network.String())
	// go run main.go 192.168.15.100
	/*
		Addr is 192.168.15.100 Default mask length is  32
		 Leading ones count is  24
		 Mask is (hex)  ffffff00
		 Network is  192.168.15.0
	*/

	// 查看域名对应ip单个
	addr, _ := net.ResolveIPAddr("ip", os.Args[1]) // go run main.go www.baidu.com
	fmt.Println(addr.String())                     // 182.61.200.7

	// 查看主机所有ip
	addrs, _ := net.LookupHost(name) // go run main.go www.baidu.com
	fmt.Println(addrs)               // [182.61.200.7 182.61.200.6]
}
