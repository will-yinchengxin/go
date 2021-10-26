package workPool

import (
	"fmt"
	//"test/workPool/pool"
	"test/workPool/poolAno"
	"time"
)

func TestPool() {
	var allTask []*poolAno.Task

	// 循环塞入 100 个任务
	for i := 0; i < 100; i++ {
		task := poolAno.NewTask(i, func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("task %d processed \n", taskID)
			return nil
		})
		allTask = append(allTask, task)
	}

	// 调用链接池，开启五个worker 执行任务
	pool := poolAno.NewPool(allTask, 5)
	pool.Run()
}
