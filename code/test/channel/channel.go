// 停止goroutine的几种方法
package channel

import (
	"context"
	"fmt"
	"time"
)

// 方法一: 借助 关闭channel 的方式关闭 goroutine
func CloseChan() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <- ch
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Println(v)
		}
	}()
	ch <- "add some message one"
	ch <- "add some message two"
	close(ch)
	time.Sleep(time.Second)
}
/*
go 中 channel 接受数据有两种方式
	1) msg := <- ch
	2) msg, ok := <- ch
*/

// 方法二: 定期轮询channel
func ChanClose() {
	ch := make(chan string, 6)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- "test chan":
			case <- done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("get value: ", i)
	}

	fmt.Println("done")
}

// 方法三: 使用 context
func WithContext() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("go on")
			}

			time.Sleep(time.Millisecond * 500)
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	<- ch
	fmt.Println("done")
}

// 我想在 goroutineA 里去停止 goroutineB，有办法吗？
/*
不可以,因为go语言中,goroutine只能自己推出,一般通过channel来控制,不能被外界其他的goroutine关闭或干掉,

如果一个goroutine被强制停止,它所拥有的资源会发生什么？堆栈被解开了吗？defer 是否被执行？
	1) 如果执行 defer，该 goroutine 可能可以继续无限期地生存下去。
	2) 如果不执行 defer，该 goroutine 原本的应用程序系统设计逻辑将会被破坏，这肯定不合理。

Go 语言中每一个 goroutine 都需要自己承担自己的任何责任
*/