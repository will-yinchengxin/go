## trace 方式
```go
package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	trace.Stop() 
}

// 执行命令(获取可达连接):
// go tool trace build main.go
````
## GODEBUG 方式
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("helle GMP")
	}
  
	/*
		执行: go build main.go
		命令: GODEBUG=schedtrace=1000 ./main

		192:test root$ GODEBUG=schedtrace=1000 ./main
		SCHED 0ms: gomaxprocs=8 idleprocs=6 threads=3 spinningthreads=1 idlethreads=0 runqueue=0 [2 0 0 0 0 0 0 0]
		helle GMP
		SCHED 1009ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
		helle GMP
		SCHED 2013ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
		helle GMP
		SCHED 3018ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
		helle GMP
		SCHED 4025ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
		helle GMP

		SCHED: 调试的信息
		0ms：从程序启动到输出所经历的时间
		gomaxprocs：P的数量，一般默认和cpu的默认核心数是一致的
		idleprocs：处于idle状态的P的数量，gomaxprocs-idleprocs = 目前正在执行的P的数量
		threads： 线程数量（包括 M0，包括DEBUG调试的线程）
		spinningthreads： 处于自旋状态的threads的数量
		idlethreads：处于idle状态的thread状态的数量
		runqueue： 全局对了中G的数量
		[0 0 0 0 0 0 0 0]： 本地P中G的数量
	*/
}
````
