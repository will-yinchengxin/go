package type

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//定义函数类型 op
type op func(a, b int) int

//形参指定传入参数为函数类型op
func Oper(fu op, a, b int) int {
	return fu(a, b)
}
func main() {
	//在go语言中函数名可以看做是函数类型的常量，所以我们可以直接将函数名作为参数传入的函数中。
	aa := Oper(add, 1, 2)
	fmt.Println(aa)
	bb := Oper(sub, 1, 2)
	fmt.Println(bb)
}
