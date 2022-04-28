package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var StrChan = make(chan string, 3)

func main() {
	fmt.Println(1e9)
	return

	Ticker()
	return

	Timer()
	return

	noBufferChannel()
	return

	test()
	return

	chanOne := make(chan struct{}, 1)
	chanTwo := make(chan struct{}, 2)

	go func() {
		<-chanOne
		fmt.Println("receive Sync signal")
		time.Sleep(time.Second)
		for {
			if val, ok := <-StrChan; ok {
				fmt.Println("receive param: ", val)
			} else {
				break
			}
		}
		chanTwo <- struct{}{}
	}()

	go func() {
		for _, i2 := range []string{"a", "b", "c", "d"} {
			StrChan <- i2
			if i2 == "c" {
				chanOne <- struct{}{}
				fmt.Println("send Sync signal")
			}
		}
		time.Sleep(time.Second * 2)
		close(StrChan)
		chanTwo <- struct{}{}
	}()
	<-chanTwo // 控制主 goroutine 不过早的执行结束
	<-chanTwo
	close(chanOne)
	close(chanTwo)
}

/*
因此，当接收方从通道接收到一个值类型的值时，对该值的修改就不会影响到发送
方持有的那个源值。但对于引用类型的值来说，这种修改会同时影响收发双方持有的值。

Map 引用类型，修改会同时影响收发双方持有的值
*/
type Count struct {
	count int
}

var MapChan = make(chan map[string]int, 1)

func test() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if val, ok := <-MapChan; ok {
				val["count"]++
				//elem := val
				//elem["count"]++
			} else {
				break
			}
		}
		syncChan <- struct{}{}
	}()
	go func() {

		for i := 0; i < 5; i++ {
			countMap := map[string]int{"count": i}
			MapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Println(countMap)
		}
		close(MapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

/*
已知，单向通道通常由双向通道转换而来。那么，单向通道是否可以转换回双向通
道呢？
答案是否定的。请记住，通道允许的数据传递方向是其类型的一部分。对于两个
通道而言，数据传递方向的不同就意味着它们米利的不同


已知．从一个还未被初始化的通道中接收元素值会导致当前 goroutine 的永久阻塞，
*/

/*
当有一个 case 被选中时，运行时系统就会执行该 case 及其包含的语句，而其他 case
会被忽略。如果同时有多个 case 满足条件，那么运行时系统会通过一个伪随机的算法选
中一个 case

另一方面，如果 select 语句中的所有 case 都不满足选择条件，并且没有 default case，
那么当前 goroutine 就会一直被阻塞于此，直到至少有一个 case 中的发送或接收操作可以
立即进行为止

其中的 break 语句，它的作用是立即结束当前 select 语句的执行。
*/

func noBufferChannel() {
	sendInterval := time.Second
	reciveInterval := time.Second * 2

	//intChan := make(chan int, 5) // 这里使用 buffer chan 和 非 buffer chan 的结果不是一致的
	intChan := make(chan int) // 这里使用 buffer chan 和 非 buffer chan 的结果不是一致的

	go func() {
		var ts0, ts1 int64
		for i := 0; i < 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("send: ", i)
			} else {
				fmt.Printf("\n send i %d, interval: %d \n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendInterval)
		}
		close(intChan)
	}()
	var ts2, ts3 int64
LOOP:
	for {
		select {
		case val, ok := <-intChan:
			if !ok {
				break LOOP
			}
			ts3 = time.Now().Unix()
			if ts2 == 0 {
				fmt.Println("\n receive val:", val)
			} else {
				fmt.Printf("receive val: %d, intercal: %d \n", val, ts3-ts2)
			}
		}
		ts2 = time.Now().Unix()
		time.Sleep(reciveInterval)
	}
	fmt.Println("end!")
}

/*
定时器

如果你在定时器到期之前停止了它，那么该定时器的字段 C 也就没有机会缓冲任何元素值了。
更具体地讲，若调用定时器的 Stop 方法的结果值为 true，那么在这之后再去试图从它的 C 字段中接收元素是不会有任何结果的。
更重要的是，这样做还会使当前 goroutine 永远阻塞！因此，在重置定时器之前一定不要再次对它的C字段执行接收操作。

另一方面，如果定时器到期了，但由于某种原因你未能及时地从它的 C 字段中接收元素值，那么该字段就会一直缓冲着那个元素值，
即使在该定时器重置之后也会是如此,由于C(也就是通知通道）的容量总是1，因此就会影响重置后的定时器再次发送到期通
知。虽然这不会造成阻塞，但是后面的通知会被直接丢掉。因此，如果你想要复用定时器，就应该确保旧的通知已被接收。

*/
func Timer() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan) // 由 producer 发起关闭
	}()
	timeOut := time.Second
	var timer *time.Timer
	//Loop:
	for {
		if timer == nil {
			timer = time.NewTimer(timeOut)
		} else {
			timer.Reset(timeOut)
		}
		select {
		case data, ok := <-intChan:
			if !ok {
				fmt.Println("end")
				//break Loop
				timer.Stop()
				return
			}
			fmt.Println("receive data: ", data)
		case <-timer.C:
			fmt.Println("time out, we will reset it for one second")
		}
	}
}

/*
断续器
*/
func Ticker() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Microsecond)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("add end")
	}()
	var sum int
	for i2 := range intChan {
		fmt.Println("receive data: ", i2)
		sum += i2
		if sum > 10 {
			fmt.Println("already ten, sum is:", sum)
			ticker.Stop()
			close(intChan)
			break
		}
	}
	fmt.Println("receive end")
}

/*
channel：
--------------------------------------------------------------------------------------------------------
造成 channel panic 的流程
- 关闭一个 nil 值 channel 会引发 panic ``func main() {  var ch chan struct{}  close(ch)}``
- 关闭一个已关闭的 channel 会引发 panic ``func main() {  ch := make(chan struct{})  close(ch)  close(ch)}``
- 向一个已关闭的 channel 发送数据 ``func main() {  ch := make(chan struct{})  close(ch)  ch <- struct{}{}}``
*/
