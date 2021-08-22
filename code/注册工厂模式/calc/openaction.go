package calc

type OperationInterface interface {
	Exec(a, b int) int
}

type OperationAdd struct {}
type OperationSub struct {}
type OperationMul struct {}
type OperationDiv struct {}

func (this *OperationAdd) Exec(a, b int) int {
	return a + b
}

func (this *OperationSub) Exec(a, b int) int {
	return a - b
}

func (this *OperationMul) Exec(a, b int) int {
	return a * b
}

func (this *OperationDiv) Exec(a, b int) int {
	return a / b
}

