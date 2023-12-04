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

// 执行: go build -buildmode=plugin -o w_plugin.so main.go
//import (
//	"fmt"
//)
//
//func HelloWorld(name string) {
//	fmt.Println("Hello", name)
//}
//func main() {}
