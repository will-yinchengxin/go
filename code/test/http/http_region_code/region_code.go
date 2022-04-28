package http_region_code

import (
	"fmt"
	"net/http"
)
/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
*/


// HandlerFunc实现了接口Handler。HandlerFunc类型只是为了方便注册函数类型的处理器。
// 我们当然可以直接定义一个实现Handler接口的类型，然后注册该类型的实例：
type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g)
}

func StartServer() {
	http.Handle("/greeting", greeting("Welcome, test"))
	// ListenAndServe 创建了一个Server 实例
	http.ListenAndServe(":8080", nil)
}