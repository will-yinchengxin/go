package goroutine

import (
	"runtime"
	"sync"
	"time"
)

type Cache struct { // 对外使用的 Cache
	*cache // cache 实例
}
type cache struct {
	defaultExpiration time.Duration             // 默认的过期时间，添加一个键值对时如果设置默认的过期时间（即代码里的 DefaultExpiration）则会使用到该值
	items             map[string]Item           // 存放的键值对
	mu                sync.RWMutex              // 读写锁
	onEvicted         func(string, interface{}) // 删除key时的回调函数
	janitor           *janitor                  // 定期清理器 定期检查过期的 Item
}
type Item struct { // 键值对
	Object     interface{} // 存放 K-V 的值，可以存放任何类型的值
	Expiration int64       // 键值对的过期时间(绝对时间)
}
type janitor struct { // 清理器结构体
	Interval time.Duration // 清理时间间隔
	stop     chan bool     // 是否停止
}

func newCache(de time.Duration, m map[string]Item) *cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	return c
}

// 运行清理器
func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{ // 实例化清理器
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.Run(c) // 开协程，定时清理器跑起来！！
}
func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.Interval) // 开定时器
	for {
		select {
		case <-ticker.C:
			// 定时调用 DeleteExpired() 执行过期删除操作, DeleteExpired() 暂时省略。。。。
		case <-j.stop: // 接收到停止清理器的信号，下面便停止定时器并返回，退出协程
			ticker.Stop()
			return
		}
	}
}
func newCacheWithJanitor(de time.Duration, ci time.Duration, m map[string]Item) *Cache {
	c := newCache(de, m) // 实例化cache
	C := &Cache{c}
	if ci > 0 {
		runJanitor(c, ci) // 开启清理器
		/*
			这里退出后台 goroutine 使用的是下面这个函数，当 GC 准备释放对象时，会调用该函数指定的方法。
			runtime.SetFinalizer(obj,func(obj *typeObj))

			当我们取消对 C 对象的引用时，如果不退出 runJanitor() 开启的 goroutine 就会造成内存泄漏。当 gc 准备释放 C 时，会调用指定函数 stopJanitor()，
			Run() 方法便能收到信号，退出协程，gc 也会将 c 释放掉
		*/
		runtime.SetFinalizer(C, stopJanitor) // 指定调用函数停止后台 goroutine
	}
	return C
}
func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}
