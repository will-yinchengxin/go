package os

import (
	"fmt"
	"os"
)

func OSGetENV() {
	env := os.Getenv("GOPATH") // 获取 go env 配置信息
	if env == "" {
		env = "dev"
	}
	fmt.Println(env) // /Users/will/go
}
