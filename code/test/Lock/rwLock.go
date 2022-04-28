package Lock

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRwLock() {
	var count Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				count.Count()
				fmt.Println("read count:", count.count)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for {
		fmt.Println("Incr count:", count.count)
		count.Incr()
		time.Sleep(time.Second)
	}
}