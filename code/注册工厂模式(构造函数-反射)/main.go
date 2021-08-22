package main

import (
	"fmt"
	"goTest/calc"
)

func main() {
	ret := calc.OperationFactory(1, 2,"+").Exec()
	fmt.Println("the result is", ret)
}