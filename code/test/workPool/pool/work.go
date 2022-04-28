package pool

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID       int
	taskChan chan *Task
}

// 创建worker 将任务放置在channel中
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
	}
}

// 循环从channel中读取任务进行消费
func (w *Worker) Start(wg *sync.WaitGroup) {
	fmt.Printf("Starting worker %d \n", w.ID)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.taskChan {
			Process(w.ID, task)
		}
	}()
}
