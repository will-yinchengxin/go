package race

import (
	"fmt"
)

/*
==================
WARNING: DATA RACE
Write at 0x00c00013e008 by goroutine 7:
  main.getNumber.func1()
      /Users/yinchengxin/GolandProjects/test/main.go:15 +0x38

Previous read at 0x00c00013e008 by main goroutine:
  main.getNumber()
      /Users/yinchengxin/GolandProjects/test/main.go:18 +0x88
  main.main()
      /Users/yinchengxin/GolandProjects/test/main.go:9 +0x33

Goroutine 7 (running) created at:
  main.getNumber()
      /Users/yinchengxin/GolandProjects/test/main.go:14 +0x7a
  main.main()
      /Users/yinchengxin/GolandProjects/test/main.go:9 +0x33
==================
*/
func Race() {
	fmt.Println(getNumber())
}
func getNumber() int {
	var i int
	go func() {
		i = 5
	}()
	return i
}
