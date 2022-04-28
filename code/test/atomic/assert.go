package atomic

import "fmt"

func Assert()  {
	var x interface{}
	x = 10
	value := x.(int)
	fmt.Print(value)

	//声明的类型要是interface
	//ok 表示断言是否正确
	var test interface{}
	test = "this is test"
	valueOne, ok := test.(string)
	if ok {
		fmt.Println(valueOne)
	}
}
