package main

import (
	"fmt"
	"goTest/calc"
)

func main() {
	ret := calc.OperationFactory("+").Exec(1, 2)
	fmt.Println("the result is", ret)
}