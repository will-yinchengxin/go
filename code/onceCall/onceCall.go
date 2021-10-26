package onceCall

import (
	"sync"
)

type CacheEntry struct {
	data []byte
	err  error
	wait *sync.Once
}

type OrderServer struct {
	cache map[string]*CacheEntry
	mutex sync.Mutex
}

func (o *OrderServer) Query(key string) ([]byte, error) {
	o.mutex.Lock()
	if o.cache == nil {
		o.cache = make(map[string]*CacheEntry)
	}

	// 如果已经有相同的请求，那么直接等待返回相同结果
	// 第一次请求，返回的数据可以从db中进行获取,这里进行手动模拟
	entry, ok := o.cache[key]
	if !ok {
		entry = &CacheEntry{
			data: make([]byte, 0),
			err: nil,
			wait: new(sync.Once),
		}
		o.cache[key] = entry
	}
	o.mutex.Unlock()
	// 请求数据,这里只请求一次
	entry.wait.Do(func() {
		entry.data, entry.err = getDbData()
	})
	return entry.data, nil
}

func getDbData() ([]byte, error) {
	return []byte("this is test data"), nil
}