package _interface

/*
接口的优势:
	1)依赖反转,这是接口在大多数语言中对软件项目产生的影响,
	2)由编译器帮助我们检测类似于"未完全实现接口"这样的错误
*/
type OrderCreate interface {
	ValidateUser()
	CreateOrder()
}

type BookOrderCreator struct {}

func (b *BookOrderCreator) ValidateUser() {
	println("this is validateUser")
}

func createOrder(oc OrderCreate) {
	oc.ValidateUser()
	oc.CreateOrder()
}

func TestInter() {
	// 这里直接调用存在问题,
	//createOrder(BookOrderCreator{})
}

