package reflectSelect

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

//type SelectDir int
//
//const (
//	_             SelectDir = iota
//	SelectSend    int       = iota // case Chan <- Send
//	SelectRecv                     // case <-Chan:
//	SelectDefault                  // default
//)

func reflectSelect() {
	var wg sync.WaitGroup
	wg.Add(2)
	var rchs []chan int
	for i := 0; i < 10; i++ {
		rchs = append(rchs, make(chan int))
	}

	// 创建SelectCase
	var cases = createRecvCases(rchs, true)

	// 消费者goroutine
	go func() {
		defer wg.Done()
		for {
			chosen, recv, ok := reflect.Select(cases)
			if cases[chosen].Dir == reflect.SelectDefault {
				fmt.Println("choose the default")
				continue
			}
			if ok {
				fmt.Printf("recv from channel [%d], val=%v\n", chosen, recv)
				continue
			}
			// one of the channels is closed, exit the goroutine
			fmt.Printf("channel [%d] closed, select goroutine exit\n", chosen)
			return
		}
	}()

	// 生产者goroutine
	go func() {
		defer wg.Done()
		var n int
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		for i := 0; i < 10; i++ {
			n = r.Intn(10)
			rchs[n] <- n
		}
		close(rchs[n])
	}()

	wg.Wait()
}

func createRecvCases(rchs []chan int, withDefault bool) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range rchs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	if withDefault {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectDefault,
			Chan: reflect.Value{},
			Send: reflect.Value{},
		})
	}

	return cases
}

// ------------------------------------------------
// 利用 reflect 包向 channel 中发送数据

func AnoReflectSelect() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch0, ch1, ch2 := make(chan int), make(chan int), make(chan int)
	var schs = []chan int{ch0, ch1, ch2}

	// 创建SelectCase
	var cases = createCases(schs)

	// 生产者goroutine
	go func() {
		defer wg.Done()
		for range cases {
			chosen, _, _ := reflect.Select(cases)
			fmt.Printf("send to channel [%d], val=%v\n", chosen, cases[chosen].Send)
			cases[chosen].Chan = reflect.Value{}
		}
		fmt.Println("select goroutine exit")
		return
	}()

	// 消费者goroutine
	go func() {
		defer wg.Done()
		for range schs {
			var v int
			select {
			case v = <-ch0:
				fmt.Printf("recv %d from ch0\n", v)
			case v = <-ch1:
				fmt.Printf("recv %d from ch1\n", v)
			case v = <-ch2:
				fmt.Printf("recv %d from ch2\n", v)
			}
		}
	}()

	wg.Wait()
}

func createCases(schs []chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建send case
	for i, ch := range schs {
		n := i + 100
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(n),
		})
	}

	return cases
}

// 测试性能
func BenchmarkReflectSelect(b *testing.B) {
	var c1 = make(chan int)
	var c2 = make(chan int)
	var c3 = make(chan int)

	go func() {
		for {
			c1 <- 1
		}
	}()
	go func() {
		for {
			c2 <- 2
		}
	}()
	go func() {
		for {
			c3 <- 3
		}
	}()

	chs := createCases([]chan int{c1, c2, c3})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = reflect.Select(chs)
	}
	// go test -bench .
	// reflect.Select的执行效率相对于select还是要差的，并且在其执行过程中还要做额外的内存分配
	// 大多数情况下，我们是不需要使用reflect.Select，常规select语法足以满足我们的要求。并且reflect.Select有对cases数量的约束，最大支持65536个cases，虽然这个约束对于大多数场合而言足够用了
}
