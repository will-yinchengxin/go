## Go run 与 Go build
### go run
编译马上运行 go 程序, 它只接收main包下的文件作为参数, 如果不是 main 包就会提示错误 `go run: cannot run non-main package`
使用改名命令所编译的二进制文件最终回存放在一个临时目录中,可以通过 -t / -x 参数来进行查看,这两个参数是打印编译中所有执行的命令
-n 不会继续执行编译后的二进制文件, -x 会继续执行编译后的二进制文件
``[root@65935aeb351a project]# go run -n[-x] main.go``

大致流程分为:
- 创建编译依赖所需要的临时目录(临时环境变量 WORK), 可以通过设置全局的 GOTMPDIR 来进行更改
- 编译和生成编译所需要的依赖
- 创建并进入编译二进制文件所需要的临时目录, 即 exe 目录
- 生成可执行文件
- 执行可执行文件

### go build
`go build [-o output] [-i] [build flags] [packages]` 编译指定的源文件,软件包及依赖包,其不会运行编译后的二进制文件
指定 二进制 文件为其他名称, 通过 -o 调整参数 go build -o 例: `go build -o youAndme`

### go build 和 go run 的区别
go build 命令和 go run 命令执行过程差不多, 唯一不同是 go build 会编译并执行生成好的二进制文件, 并将其命名为 blog-server(当前的目录名)
并删除编译生成的临时目录
