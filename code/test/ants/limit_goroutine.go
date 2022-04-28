package ants

import "sync"

/*
	// 设置最多有3个goroutine同时工作
	sem :=  ants.NewSemaphore(3)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer sem.Done()
			sem.Add(1)
			fmt.Printf("%d I am working at time %d \n", i, time.Now().Unix())
			// 为了更好的查看效果这里延迟2秒
			time.Sleep(time.Second * 2)
		}(i)
	}
	sem.Wait()
*/

// 控制同时工作的 goroutine
type semaphore struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

func NewSemaphore(max int) *semaphore {
	return &semaphore{
		ch: make(chan struct{}, max),
		wg: new(sync.WaitGroup),
	}
}

func (s *semaphore) Add(delta int)  {
	s.wg.Add(1)
	for i := 0; i < delta; i++ {
		s.ch <- struct{}{}
	}
}

func (s *semaphore) Done()  {
	<- s.ch
	s.wg.Done()
}

func (s *semaphore) Wait()  {
	s.wg.Wait()
}