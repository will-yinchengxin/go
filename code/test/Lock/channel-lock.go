package Lock

import "time"

type Mutex struct {
	lock chan struct{}
}

func NewMutex() *Mutex {
	channel := make(chan struct{}, 1)
	channel <- struct{}{}
	return &Mutex{
		lock: channel,
	}
}

func (m *Mutex) Lock() {
	select {
	case <-m.lock:
	}
}

func (m *Mutex) Unlock() {
	select {
	case m.lock <- struct{}{}:
	default:
		panic("unlock false")
	}
}

func (m *Mutex) tryLock() {
	timer := time.NewTicker(time.Second)
	select {
	case <-m.lock:
	case <-timer.C:
	}
}
