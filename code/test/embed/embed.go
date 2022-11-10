package embed

import (
	_ "embed"
	"fmt"
)

/*
embed是在Go 1.16中新加包。它通过//go:embed指令，可以在编译阶段将静态资源文件打包进编译好的程序中，并提供访问这些文件的能力。

embed的三种数据类型及使用
在embed中，可以将静态资源文件嵌入到三种类型的变量，分别为：字符串、字节数组、embed.FS文件类型

https://zhuanlan.zhihu.com/p/351931501
*/
var (
	//go:embed vrpm.json
	vrpm string
)

func main() {
	fmt.Println(vrpm)
}
