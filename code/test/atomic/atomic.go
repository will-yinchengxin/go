package atomic

import (
	"fmt"
	"sync/atomic"
)

func Atomic() {
	// 载入
	var i64 int64 = 1
	num := atomic.LoadInt64(&i64)
	fmt.Println(num) // 1

	// 增减操作
	atomic.AddInt64(&i64, 12)
	fmt.Println(i64) // 13

	// 存储
	atomic.StoreInt64(&i64, 78)
	fmt.Println(i64) // 78
}

func Value() {
	var countVal atomic.Value
	// store 接收一个 interface 类型的数据
	countVal.Store([]int{1,2,2,3})
	// load 返回 atomic.Value 中的值
	list := countVal.Load().([]int)
	fmt.Println(list) // [1 2 2 3]
}