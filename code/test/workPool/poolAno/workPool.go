package poolAno

import "sync"

type Pool struct {
	TaskList []*Task
	TaskChan chan *Task
	WorkNum int
	wg sync.WaitGroup
}

func NewPool(taskList []*Task, workNum int) *Pool {
	return &Pool{
		TaskList: taskList,
		WorkNum: workNum,
		TaskChan: make(chan *Task, 100),
	}
}

func (p *Pool) Run() {
	// 唤起指定数据的 work 进程
	// 此时TaskChan 为 buffer channel， 如果channel中内容为空，消费段会阻塞等待任务
	for i := 0; i < p.WorkNum; i++ {
		worker := NewWork(i, p.TaskChan)
		worker.Start(&p.wg)
	}

	for _, task := range p.TaskList {
		p.TaskChan <- task
	}
	// 消费完成关闭channel 避免内存泄露
	close(p.TaskChan)
	p.wg.Wait()
}