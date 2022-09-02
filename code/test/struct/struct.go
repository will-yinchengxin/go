package struct

import (
	"fmt"
	"time"
	"unsafe"
)

/*
	类型								占用字节大小
	byte, uint8, int8					1
	uint16, int16						2
	uint32, int32, float32				4
	uint64, int64, float64, complex64	8
	complex128							16

	type People struct {
		ID          int64     // Sizeof: 8 byte  Alignof: 8  Offsetof: 0
		Gender      int8      // Sizeof: 1 byte  Alignof: 1  Offsetof: 8
		NickName    string    // Sizeof: 16 byte Alignof: 8 Offsetof: 16
		Description string    // Sizeof: 16 byte Alignof: 8 Offsetof: 32
		IsDeleted   bool      // Sizeof: 1 byte  Alignof: 1  Offsetof: 48
		Created     time.Time // Sizeof: 24 byte Alignof: 8  Offsetof: 56
	}
*/
type People struct {
	Man
	Created     time.Time
	NickName    string
	Description string
	ID          int64
	Gender      int8
	IsDeleted   bool
}

type Man struct {
	Created     time.Time
	NickName    string
	Description string
	ID          int64
	Gender      int8
	IsDeleted   bool
}

/*
This analyzer find structs that can be rearranged to use less memory, and provides
a suggested edit with the most compact order.
*/
func main() {
	p := People{}
	fmt.Println(unsafe.Sizeof(p)) // 80

	/*
		go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
		fieldalignment -fix ./    # 使用

		WilldeMacBook-Pro:test yinchengxin$ fieldalignment -fix ./
		/Users/yinchengxin/GolandProjects/test/main.go:26:13: struct with 152 pointer bytes could be 128
		/Users/yinchengxin/GolandProjects/test/main.go:36:10: struct of size 80 could be 72

	*/
}
