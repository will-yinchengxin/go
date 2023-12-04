### 简介
Golang插件是一种可用于扩展和增强Golang应用程序功能的机制。它允许开发人员将功能封装为可重用的模块，并将其作为插件加载到应用程序中。

使用规则：
- 创建插件：开发人员可以使用Golang编写插件代码，并将其编译为共享库（.so文件）。
- 加载插件：应用程序可以使用Golang的插件包（plugin）加载插件。加载插件时，应用程序需要指定插件的路径和名称。
- 调用插件函数：一旦插件加载成功，应用程序可以通过插件对象调用插件中定义的函数。

场景：
- 动态扩展功能：插件机制允许应用程序在运行时加载和卸载插件，从而实现动态扩展功能。例如，一个文本编辑器可以加载插件来支持不同的文件格式。
- 模块化开发：插件机制可以帮助开发人员将应用程序拆分为多个模块，每个模块负责不同的功能。这样可以提高代码的可维护性和可重用性。
- 自定义功能：插件机制允许用户根据自己的需求自定义应用程序的功能。例如，一个图像处理应用程序可以提供插件接口，让用户编写自己的图像处理算法。

### 使用 Demo
创建插件代码：
```go
package main

import (
	"fmt"
)

func HelloWorld(name string) {
	fmt.Println("Hello", name)
}
func main() {}
````
执行: `go build -buildmode=plugin -o w_plugin.so main.go`

创建应用程序代码：
```go
package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("w_plugin.so")
	if err != nil {
		log.Fatal("plugin.Open", err)
	}
	helloWorldFunc, err := p.Lookup("HelloWorld")
	if err != nil {
		log.Fatal("p.Lookup", err)
	}
	helloWorldFunc.(func(string))("Will")
}
````
执行：`go run main.go`

输出结果:
```
Hello Will
```