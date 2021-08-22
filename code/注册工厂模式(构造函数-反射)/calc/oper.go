package calc

import (
	"reflect"
)

var opers map[string]interface{}

func init() {
	// interface{} 可以存放任意的类型
	// 函数方法是一个类型
	opers = make(map[string]interface{}, 0)
	opers["-"] = NewSub
	opers["+"] = NewAdd
	opers["*"] = NewMul
	opers["/"] = NewDiv
}

// 利用反射获取对象实例
func OperationFactory(num1, num2 int, flag string) OperationInterface {
	oper := opers[flag]
	// 根据类型得到反射的结构体对象Value
	valueOper := reflect.ValueOf(oper) // reflect.value

	// 设置调用函数的参数
	args := []reflect.Value{
		reflect.ValueOf(num1), reflect.ValueOf(num2),
	}
	// 根据类型调用函数或方法
	arrValueOper := valueOper.Call(args)[0]

	// 利用接口interface{} 断言转变为原有类型 ； interface{}=> 具体类型
	opertionin := arrValueOper.Interface().(OperationInterface) // 返回一个 interface{} 类型

	return opertionin
}