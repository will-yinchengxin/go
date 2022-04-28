package Lock

import "sync"

type SliceQueue struct {
	data []interface{}
	mu sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

func (s *SliceQueue) Enqueue(v interface{}) {
	s.mu.Lock()
	s.data = append(s.data, v)
	s.mu.Unlock()
}

func (s *SliceQueue) Dequeue() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v
}