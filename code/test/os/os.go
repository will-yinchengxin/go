package os

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bs, err := ioutil.ReadAll(io.LimitReader(os.Stdin, int64(5))) // 最大字节数
	_, _ = os.Stdout.WriteString(string(bs) + "finish")

	// os.Stdout.WriteString("will——test，input：")
	//var a string
	//fmt.Scan(&a)
	//fmt.Println(a)
}
