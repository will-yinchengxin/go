package goroutine

import (
	"fmt"
	"time"
)

/*
	只要main函数不结束，子goroutine会一直执行
	main函数结束，子goroutine会立即结束
*/


type obj struct {
}
var ch = make(chan obj) //

func main() {
	//string.StringClone()
	doReq()
	select {
	case <-ch :
		fmt.Println(234432)
	}
}

// 这里会出现 父 goroutine 停止了 子 goroutine 还在执行的问题
// 原因: 这里使用的是非 buffer channel
// 解决: 解决办法是将 ch 从无缓冲的通道改为有缓冲的通道，因此子goroutine 即使在父 goroutine 退出后也始终可以发送结果
func doReq() {

	go func() {
		ch <- obj{}
		fmt.Println(ch)
	}()
	select {
	case a := <-ch:
		fmt.Println(a, "out")
		return
	case <- time.After(time.Nanosecond): // 这里需要注意的是 就是 time.After 导致的内存泄露问题，只要注意程序不是频繁执行 select 即可
		fmt.Println("time out")
	}
}
/* 结果:
	time out
	0xc00006e060
	234432
*/