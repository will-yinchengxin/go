package atomic
/*
atomic.Value 再没有加锁的情况下, 保证了线程安全

原因:
	1) atomic.Value 被设计用来存储任意类型的数据, 所以它的类型是一个 interface, 其内部定义了一个 ifaceWords 类型, 其实是 interface 的内部表示
		, 作用位将 interface 分解, 得到其原始类型 和 真正的值
		type ifaceWords struct {
			typ  unsafe.Pointer
			data unsafe.Pointer
		}
	2) 那就先看看 unsafe.Pointer, go 中不支持直接操作内存, 但 unsafe.Pointer 可以让程序灵活的错做内存, 它绕过了类型系统的检查,
		与任意的指针类型互相转换。也就是说，如果两种类型具有相同的内存结构（layout），我们可以将unsafe.Pointer当做桥梁，

		!!!! 让这两种类型的指针相互转换，从而实现同一份内存拥有两种不同的解读方式。

		比如说，[]byte和string其实内部的存储结构都是一样的，他们在运行时类型分别表示为 reflect.SliceHeader 和 reflect.StringHeader

		例如:
		bytes := []byte{104, 101, 108, 108, 111}
		p := unsafe.Pointer(&bytes) //将 *[]byte 指针强制转换成unsafe.Pointer
		str := *(*string)(p) //将 unsafe.Pointer再转换成string类型的指针，再将这个指针的值当做string类型取出来
		fmt.Println(str)
	3) atomic.Store 的大致逻辑:
		- 通过unsafe.Pointer将现有的和要写入的值分别转成ifaceWords类型，这样我们下一步就可以得到这两个interface{}的原始类型（typ）和真正的值（data)
		- 开始就是一个无限 for 循环。配合CompareAndSwap使用，可以达到乐观锁的效果
		- 通过LoadPointer这个原子操作拿到当前Value中存储的类型。下面根据这个类型的不同，分3种情况处理
			- 第一次写入 - 一个atomic.Value实例被初始化后，它的typ字段会被设置为指针的零值 nil，所以先判断如果typ是nil 那就证明这个Value实例还未被写入过数据。那之后就是一段初始写入的操作
				- runtime_procPin()这是runtime中的一段函数，一方面它禁止了调度器对当前 goroutine 的抢占（preemption），使得它在执行当前逻辑的时候不被打断，以便可以尽快地完成工作，因为别人一直在等待它。另一方面，在禁止抢占期间，GC 线程也无法被启用，这样可以防止 GC 线程看到一个莫名其妙的指向^uintptr(0)的类型（这是赋值过程中的中间状态）。
				- 使用CAS操作，先尝试将typ设置为^uintptr(0)这个中间状态。如果失败，则证明已经有别的线程抢先完成了赋值操作，那它就解除抢占锁，然后重新回到 for 循环第一步
				- 如果设置成功，那证明当前线程抢到了这个"乐观锁”，它可以安全的把v设为传入的新值了。注意，这里是先写data字段，然后再写typ字段。因为我们是以typ字段的值作为写入完成与否的判断依据的。
			- 第一次写入还未完成- 如果看到typ字段还是^uintptr(0)这个中间类型，证明刚刚的第一次写入还没有完成，所以它会继续循环，一直等到第一次写入完成
			- 第一次写入已完成 - 首先检查上一次写入的类型与这一次要写入的类型是否一致，如果不一致则抛出异常。反之，则直接把这一次要写入的值写入到data字段

		这个逻辑的主要思想就是，为了完成多个字段的原子性写入，我们可以抓住其中的一个字段，以它的状态来标志整个原子写入的状态

	4) atomic.Load 的大致逻辑:
		- 如果当前的typ是 nil 或者^uintptr(0)，那就证明第一次写入还没有开始，或者还没完成，那就直接返回 nil （不对外暴露中间状态）
		- 否则，根据当前看到的typ和data构造出一个新的interface{}返回出去
*/
import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

// atomic value
var AtomicValue atomic.Value

type Rectangle struct {
	width  int
	length int
}

type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func update(width, length int) {
	re := new(Rectangle)
	re.width = width
	re.length = length
	// ifaceWords 是 interface 的内部表现,
	//fmt.Println((*ifaceWords)(unsafe.Pointer(&re)).data, (*ifaceWords)(unsafe.Pointer(&re)).typ)
	AtomicValue.Store(re)
}

func TestAtomicValue() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			update(i, i+5)
		}()
	}
	wg.Wait()
	_r := AtomicValue.Load().(*Rectangle)
	fmt.Println(_r) // &{10 15}
}

func UnsafePointer() {
	bytes := []byte{104, 101, 108, 108, 111}

	p := unsafe.Pointer(&bytes) //将 *[]byte 指针强制转换成unsafe.Pointer
	str := *(*string)(p) //将 unsafe.Pointer再转换成string类型的指针，再将这个指针的值当做string类型取出来
	fmt.Println(str)
}