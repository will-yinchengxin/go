package flag

import (
	"flag"
	"fmt"
)

func TestFlag() {
	var module string
	flag.StringVar(&module, "module", "", "assign run module")    // 不包含默认值 go run main.go --module will
	flag.StringVar(&module, "module", "yin", "assign run module") // 包含默认值： go run main.go
	flag.Parse()
	fmt.Println(fmt.Sprintf("start run %s module", module))
	fmt.Println(fmt.Sprintf("run %s module done!", module))
}
