package cond

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Cond() {
	c := sync.Cond{L: &sync.Mutex{}}
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			fmt.Printf("运动员 %d 已经就绪\n", i)
			// 广播唤醒所有等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		fmt.Println("裁判员被唤醒一次\n")
	}
	c.L.Unlock()
	// 准备起跑
	fmt.Println("准备起跑\n")
}
