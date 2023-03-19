## go build命令
```go
package main

import (
	"fmt"
	"os"
	"runtime"
)

/*
 * Auth：Will Yin
 * Date：2023/3/19 15:20

 go build -o will main.go   		// output 指定编译输出的名称，代替默认的包名

 go build -o -x will main.go   		// -x 打印出执行的命令名

 -ldflags "-X" 通过链接选项 -X 来动态传入版本信息,
 flags="-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`"
 go build -ldflags "-X main.name=will" -o will main.go
 go build -ldflags "-X 'main.name=will' -X 'main.version=$(go version)'" -o will main.go # 多个参数要以 '' 包裹分隔
 # 这里使用 main.** 正常可以替换成 pkgName.*** 如: -X 'github.com/willyin/willshark/conf.VERSION=$TRAVIS_TAG'

 -ldflags "-s -w"  # 压缩编译后的体积
 go build -ldflags "-X 'main.name=will' -X 'main.version=$(go version)' -s -w"  -o will main.go
*/

var (
	name      = ""
	version   = ""
	goVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "-n" || args[1] == "--name") {
		fmt.Println("Get Name: ", name)
		fmt.Println("Go Version: ", version)
		fmt.Println("Go Version Runtime: ", goVersion)
		return
	}
}

/*
 * Auth：Will Yin
 * Date：2023/3/19 16:16

 编译跨平台的只需要修改GOOS、GOARCH、CGO_ENABLED三个环境变量即可
 Linux 下编译 Mac 和 Windows 64位可执行程序
 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
 	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
*/
````
