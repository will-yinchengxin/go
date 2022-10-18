package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// 内存泄漏
func TestLeak() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d, %d \n", i, <-randStream)
	}
}

// 解决内存泄漏
func DealLeak() {
	newRandStream := func(done chan struct{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}
	done := make(chan struct{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d, %d \n", i, <-randStream)
	}
	close(done)
	time.Sleep(time.Second) // 模拟郑子啊工作，查看 defer fmt 语句内容
	/*
		3 random ints:
		0, 5577006791947779410
		1, 8674665223082153551
		2, 6129484611666145821
		newRandStream closure exited.
	*/
	// 负责创建 goroutine 也负责关掉对应的 goroutine
}
