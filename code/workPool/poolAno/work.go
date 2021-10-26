package poolAno

import (
	"fmt"
	"sync"
)

type Work struct {
	Id int
	taskChan chan *Task
}

func NewWork(Id int, taskChan chan *Task) *Work {
	return &Work{
		Id: Id,
		taskChan: taskChan,
	}
}

func (w *Work) Start(wg *sync.WaitGroup) {
	fmt.Printf("Start working %d \n", w.Id)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.taskChan {
			Process(w.Id, task)
		}
	}()
}