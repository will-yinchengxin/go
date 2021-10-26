package poolAno

import "fmt"

type Task struct {
	Err error
	Data interface{}
	f func(data interface{}) error
}

func NewTask(data interface{}, f func(data interface{}) error) *Task {
	return &Task{
		Data: data,
		f: f,
	}
}

func Process(workId int, task *Task) {
	fmt.Printf("workId %d process task %v \n", workId, task.Data)
	task.Err = task.f(task.Data)
}