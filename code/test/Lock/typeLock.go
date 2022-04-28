package Lock

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制mutex定义的常量
const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaterShift = iota
)

type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 如果成功则抢占到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	// 如果处于唤醒, 加锁状态或者饥饿状态,这次请求不参与竞争,返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	// & 两个位都为1时，结果才为1
	// | 两个位都为0时，结果才为0
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}
	// 尝试竞争状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func TestTryLock() {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}
	fmt.Println("can not get the lock")
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaterShift // 获取等待者的数量
	v = v + (v & mutexLocked)
	return int(v)
}

// 是否被锁
func (m *Mutex) IsLocked() bool {
	// & 两个位都为1时，结果才为1
	// | 两个位都为0时，结果才为0
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 是否处于被唤醒模式
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 是否饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}