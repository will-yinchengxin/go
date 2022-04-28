package _func

import (
	"errors"
	"fmt"
)

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (byteNum int, err error) {
	return fmt.Println(contents)
}

func GetCon() {
	var p Printer
	p = printToStd
	p("something")
}

/*
高阶函数定义：
	1）接受其他函数作为参数传入
	2）把其他的函数作为结果返回

tips：
	函数类型属于引用类型，它的值可以为nil，而这种类型的零值恰恰就是nil
*/
type operate func(x, y int) int

var op = func(x, y int) int {
	return x + y
}
//func Calculate(x int, y int, op operate) (int, error) {
//	if op == nil {
//		return 0, errors.New("invalid operation")
//	}
//	return op(x, y), nil
//}
// --------------------------------------------------------
type calculateFunc func(x, y int) (int, error)
// 定义一个匿名函数(闭包)作为返回值
/*
闭包存在的意义：
	1. 提高某个功能的灵活性，可以让使用方提供一部分功能的实现。但却可以控制这一部分的大小。
	2. 提供动态替换某个功能的部分实现的可能性。这里的关键在于动态。
	3. 使得代码动态替换的粒度缩小到函数级别。相比之下，模版类型的动态替换粒度是实例级别的。
*/
func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}