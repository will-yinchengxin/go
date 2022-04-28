package Lock

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/*
 go 通过 sync/atomic 提供了对原子操作的支持,支持的类型为 int32 int64 uint32 uintptr, 使用时以实际类型替换 XXXtype也就是对应的操作方法

 那么 Mutex 和 atomic 包操作有什么区别呢
	目的: 互斥锁(Mutex)保护一段逻辑,原子操作(atomic)用于对一个变量的更新保护
	底层实现: 互斥锁(Mutex)由操作系统调度实现, 原子操作(atomic)由底层硬件指令直接提供支持,这些指令在执行过程种是不允许中断的,因此原子操作
		可以在 lock-free 的情况下保证并发安全,并且它的性能可以做到随CPU个数增长而多线性拓展

	注意一个问题:
		Mutex 中的 Unlock 会被任意goroutine释放,即便没有这个互斥锁的goroutine,也可以进行这个操作,Mutex本身没有包含
		goroutine的信息,所以Unlock不会对此进行检查,这种涉及 Mutex 一直保持至今
		所以我们再申请使用锁的时候一定遵循 谁使用谁释放 的原则
*/

// 计数器
func MutexAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpend := time.Now().Sub(start).Seconds()
	fmt.Printf("use mutex a is %d, time is %v second\n", a, timeSpend)
}

func AtomicAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	timeSpend := time.Now().Sub(start).Seconds()
	fmt.Printf("use atomic a is %d, time is %v second\n", atomic.LoadInt32(&a), timeSpend)
}

/*
	所有原子操作方法的被操作参数形式必须为指针类型,通过指针可以获取被操作数在内存中的地址,从而施加特殊的CPU指令,
	确保同一时间只有一个goroutine能够进行操作
*/

// 比较交换
// 简称CAS(compare adn swap) atomic下这一类的操作的前缀为 CompareAndSwap
// 这里类操作在进行交换前先确保被操作的值没有被改变,即仍然保存这参数old所记录的值,满足此前提条件才进行交换操作
// 类似于数据库常见的乐观锁操作
// 当有大量的goroutine进行读写操作的时候,可能导致CAS操作无法成功,可以利用for循环多次尝试

// unsafe.Pointer提供了绕过Go语言指针类型限制的方法，unsafe指的并不是说不安全，而是说官方并不保证向后兼容。

func CAS() {
	type P struct {
		x, y, z int
	}

	var pP *P

	// 定义一个执行unsafe.Pointer值的指针变量
	// fmt.Println(&pP , unsafe.Pointer(&pP), (*unsafe.Pointer)(unsafe.Pointer(&pP)))
	// 0xc000006028 0xc000006028 0xc000006028
	var unsafe1 = (*unsafe.Pointer)(unsafe.Pointer(&pP))

	// old Pointer
	var sy P

	// 为了演示效果先将unsafe1设置成old pointer
	px := atomic.SwapPointer(unsafe1, unsafe.Pointer(&sy))
	fmt.Println(unsafe1)
	// 执行CSA操作,交换成功,返回结果true
	/*
		  func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
		  伪代码:
			 if *addr == old {
				*addr = new
				return true
			  }
			  return false
	*/
	y := atomic.CompareAndSwapPointer(unsafe1, unsafe.Pointer(&sy), px)
	fmt.Println(y)
}

/*
  互斥锁Mutex中有个字段 state 字段, 表示锁状态的状态位
	type Mutex struct {
		state int32
		sema  uint32
	}
  0 代表锁空闲 1 代表加锁

  Mutex 部分实现代码:
	func (m *Mutex) Lock() {
	   // Fast path: grab unlocked mutex.
	   if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		   if race.Enabled {
			   race.Acquire(unsafe.Pointer(m))
		   }
		   return
	   }
	   // Slow path (outlined so that the fast path can be inlined)
		m.lockSlow()
	}
*/

// atomic.Value保证任意值的读写安全
// atomic 提供了一套以 Store 开头的方法,用来保证各种类型变量的并发写安全,避免其他操作读取到了修改过程中的脏数据

// func StoreInt32(addr *int32, val int32)

// 如果想要并发安全的设置一个结构体的多个字段
//	1) 可以将其转化为指针,通过StorePointer设置
//  2) 还可以使用atomic.Value,它在底层为我们完成了从具体 指针类型 到 unsafe.Pointer 之间的转换
//      它使得我们可以不依赖于不保证兼容性的unsafe.Pointer类型, 同时又能将任意数据类型的读写操作封装成原子性操作（中间状态对外不可见）

func AtomicValue() {
	type Rectangle struct {
		length int
		width  int
	}
	var rect atomic.Value

	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rectLocal := new(Rectangle)
			rectLocal.width = i
			rectLocal.length = i + 5
			rect.Store(rectLocal)
		}()
	}
	wg.Wait()
	res := rect.Load().(*Rectangle)
	fmt.Println(res.length, res.width)
}

