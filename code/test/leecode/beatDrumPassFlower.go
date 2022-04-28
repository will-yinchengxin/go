package leecode

import (
	"fmt"
	"time"
)
/*
击鼓传花: 四个goroutine
*/
type token struct {}

func Drum() {
	num := 4
	var chs []chan token
	// 向切片中插入数据，循环执行4次
	for i := 0; i < num; i++ {
		chs = append(chs, make(chan token))
	}
	// 1 % 4 = 1 / 2 % 4 = 2 / 3 % 4 = 3 / 4 % 4 = 0
	// 唤醒4个 goroutine
	for i := 0; i < num; i++ {
		go worker(i, chs[i], chs[(i+1)%num])
	}
	chs[0] <- struct{}{} // 删除此行： fatal error: all goroutines are asleep - deadlock!

	for {} // 或者使用 select{}
}

func worker(id int, ch chan token, next chan token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		next <- token
	}
}
