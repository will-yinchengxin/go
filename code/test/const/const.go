package _const

/*
	每当const出现时, 都会使iota初始化为0.
	const中每新增一行常量声明将使iota计数一次
*/
const (
	mutexLocked1 =  iota // mutex is locked
	mutexWoken1
	mutexStarving1
	mutexWaiterShift1
) // 0 1 2 3
const (
	mutexLocked2 =  iota + 1// mutex is locked
	mutexWoken2
	mutexStarving2
	mutexWaiterShift2
) // 1 2 3 4

const (
	a1 = iota   // a1 = 0   // 又一个const出现, iota初始化为0
	a2 = iota   // a1 = 1   // const新增一行, iota 加1
	a3 = 6      // a3 = 6   // 自定义一个常量
	a4          // a4 = 6   // 不赋值就和上一行相同
	a5 = iota   // a5 = 4   // const已经新增了4行, 所以这里是4
)

// 二进制的1不断进位
const (
	mutexLocked = 1 << iota // 0001 移动 0位
	mutexWoken				// 0001 移动 1 位 0010 也就是 2
	mutexStarving			// 0010 移动 1 位 0100 也就是 4
 	mutexWaiterShift
	mutexWaiterShift22 = iota
)

// 	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, mutexWaiterShift22)