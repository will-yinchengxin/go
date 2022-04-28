package ants

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
)

var num int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&num, n)
	println("run with ", n)
}

func Ants() {
	var wg sync.WaitGroup
	poolFunc, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer poolFunc.Release()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		_ = poolFunc.Invoke(int32(1))
	}
	wg.Wait()
	fmt.Printf("running goroutine %d \n", poolFunc.Running())
	fmt.Printf("finish all task, result is %d \n", num)
	/*
		running goroutine 10
		finish all task, result is 1000
	*/
}