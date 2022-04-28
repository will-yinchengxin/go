package errgroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func ErrContext() {
	g, ctx := errgroup.WithContext(context.Background())
	dataChan := make(chan int, 20)

	// 数据生产段任务 子goroutine
	g.Go(func() error {
		defer close(dataChan)
		for i := 0; ; i++ {
			if i == 10 {
				return fmt.Errorf("data alreay fetch 10")
			}
			dataChan <- i
			fmt.Println(fmt.Sprintf("sending %d", i))
		}
	})

	// 数据消费端任务 子goroutine（3个）
	for i := 0; i < 3; i++ {
		g.Go(func() error {
			for i := 1; ; i++ {
				select {
				case <- ctx.Done():
					return ctx.Err()
				case number := <-dataChan:
					fmt.Println(fmt.Sprintf("receiving data %d", number))
				}
			}
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("main goroutine done!")

}
