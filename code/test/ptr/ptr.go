package ptr

import "fmt"

func Ptr () {
	//----------------- ptr --------------------------------------------
	var test = [3]int{1, 2, 3}
	fmt.Printf("arr -> %p \n", &test[0]) // arr -> 0xc00000e380
	fmt.Printf("arr -> %p \n", &test[1]) // arr -> 0xc00000e388
	fmt.Printf("arr -> %p \n", &test)    // arr -> 0xc00000e380
	// go 可以使用 & 运算符取地址,也可以使用new创建指针
	// go 的数组名不是元素首地址
	// go 中不支持指针运算,但是可以利用 unsafe 绕过Go语言的类型系统,直接跳过内库
	//----------------- ptr --------------------------------------------
}

