package calc

type OperationInterface interface {
	Exec() int
}

type Operation struct {
	num1, num2 int
}

func (o *Operation) SetNum1(num int) {
	o.num1 = num
}

func (o *Operation) SetNum2(num int) {
	o.num2 = num
}

func (o *Operation) GetNum1() int {
	return o.num1
}

func (o *Operation) GetNum2() int {
	return o.num2
}

type OperationAdd struct {
	oper *Operation
}

func NewAdd(num1, num2 int) *OperationAdd {
	return &OperationAdd{
		oper: &Operation{
			num1: num1,
			num2: num2,
		},
	}
}

type OperationSub struct {
	oper *Operation
}

func NewSub(num1, num2 int) *OperationSub {
	return &OperationSub{
		oper: &Operation{
			num1: num1,
			num2: num2,
		},
	}
}

type OperationMul struct {
	oper *Operation
}

func NewMul(num1, num2 int) *OperationMul {
	return &OperationMul{
		oper: &Operation{
			num1: num1,
			num2: num2,
		},
	}
}

type OperationDiv struct {
	oper *Operation
}

func NewDiv(num1, num2 int) *OperationDiv {
	return &OperationDiv{
		oper: &Operation{
			num1: num1,
			num2: num2,
		},
	}
}

func (o *OperationAdd) Exec() int {
	return o.oper.num1 + o.oper.num2
}

func (o *OperationSub) Exec() int {
	return o.oper.num1 - o.oper.num2
}

func (o *OperationMul) Exec() int {
	return o.oper.num1 * o.oper.num2
}

func (o *OperationDiv) Exec() int {
	return o.oper.num1 / o.oper.num2
}