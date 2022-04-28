package waitGroup

/*
// 使用 unsafe.SizeOf(变量) // 查看对应类型占用的字节数
var arr [2]uint32
fmt.Println("arr的字节数为", unsafe.Sizeof(&arr)) // 8个字节, 64bit

// 对齐系数
func main()  {
 fmt.Printf("string alignof is %d\\n", unsafe.Alignof(string("a")))
 fmt.Printf("complex128 alignof is %d\\n", unsafe.Alignof(complex128(0)))
 fmt.Printf("int alignof is %d\\n", unsafe.Alignof(int(0)))
}

// 运行结果
// string alignof is 8
// complex128 alignof is 8
// int alignof is

注意：不同硬件平台占用的大小和对齐值都可能是不一样的
*/

/*
type WaitGroup struct {
	// 辅助字段,辅助vet工具检测是否通过copy赋值这个WaitGroup实例
	noCopy noCopy
	// state1,一个复合意义的字段,包含 WaitGroup 计数/阻塞在检查点的waiter数/信号量
	state1 [3]uint32
}

// 得到state的地址和信号量的地址
func (wg *WaitGroup) state() (statep *uint64, semap *uint32) {
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
		// 如果地址是64bit对齐的,数据前两个元素做state, 最后一个做信号量
		return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
	} else {
		// 如果地址是32bit对齐的,数组后两个做state, 第一个做信号量
		return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
	}
}
*/
