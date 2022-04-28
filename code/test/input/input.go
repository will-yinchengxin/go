package input

import (
	"bufio"
	"fmt"
	"os"
)

/*
从输入中读取信息
*/
func Input() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your content")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("find err %s \n", err)
	} else {
		input = input[:len(input)-1]
		fmt.Printf("Hello, %s \n", input)
	}
}
