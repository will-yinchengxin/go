package pool

import (
	"sync"
)

// pool 的作用生成指定数目的 worker 将 task 分发下来
type Pool struct {
	Tasks       []*Task
	concurrency int
	collector   chan *Task
	wg          sync.WaitGroup
}

func NewPool(task []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       task,
		concurrency: concurrency,
		collector:   make(chan *Task, 1000),
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		woker := NewWorker(p.collector, i)
		woker.Start(&p.wg)
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}
	close(p.collector)

	p.wg.Wait()
}
