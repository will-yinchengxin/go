package pool

import "fmt"

type Task struct {
	Err error
	Data interface{}
	f func(interface{}) error
}

// 传递data 和 待执行的函数
func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{
		Data: data,
		f: f,
	}
}

// 消费task任务的函数
// 利用 f函数 消费 data 将错误信息返回给 err
func Process(workId int, task *Task) {
	fmt.Printf("worker %d process task %v \n", workId, task.Data)
	task.Err = task.f(task.Data)
}