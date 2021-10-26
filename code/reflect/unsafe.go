package reflect

import (
	"unsafe"
	"fmt"
)

func unSafe() {
	// unsafe 会绕过go类型的安全和内存安全检查的Go代码
	var value int64 = 5
	var p1 = &value
	fmt.Println(p1, &p1, &value, unsafe.Pointer(p1)) // 0xc0000100a0 0xc000006028 0xc0000100a0 0xc0000100a0
	var p2 = (*int32)(unsafe.Pointer(p1))
	fmt.Println(p2) // 0xc0000100a0
	// 任何go指针都可以转化为 unsafe.Pointer 指针。
	// unsafe.Pointer 类型的指针可以覆盖掉go的系统类型
	fmt.Println(*p1) // 5
	fmt.Println(*p2) // 5
	*p1 = 31243
	fmt.Println(*p1, *p2, p2) // 31243 31243 0xc000006028 0xc0000100a0
}
