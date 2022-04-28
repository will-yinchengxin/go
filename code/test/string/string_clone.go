package string

import (
	"fmt"
	"reflect"
	"unsafe"
)

// pkg: https://pkg.go.dev/strings@master#Clone

// src\cmd\vendor\golang.org\x\sys\internal\unsafeheader\unsafeheader.go
func StringClone() {
	/*
		string 的底层结构:

		type String struct {
			Data unsafe.Pointer
			Len  int
		}

		reflect.StringHeader 结构是对字符串底层结构的反射表示
	*/
	s := "adsflajdsfads"
	s1 := s[:4]
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	s1Header := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println(sHeader.Len == s1Header.Len) // false len 不相等很好理解
	fmt.Println(sHeader.Data == s1Header.Data) // true data 相等?
	// s 很大, 而我们需要它的某个短字串, 这会导致内存浪费, 字串和原字串指向相同的内存, 因此整个字符串并不会被GC回收

	// string clone 就是为了解决这个问题的
	//s2 := strings.Clone(s[:4]) // 这里是 go 1.15.8, 再新版本的 1.18 中加入了此功能
	s2 := Clone(s[:4])
	s2Header := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Println(sHeader.Len == s2Header.Len) // false
	fmt.Println(sHeader.Data == s2Header.Data) // false
}

// clone 的内部实现 这里就是clone的内部实现
func Clone(s string) string {
	if len(s) == 0 {
		return ""
	}
	b := make([]byte, len(s))
	copy(b, s) // copies elements from a source slice into a destination slice
	           // (As a special case, it also will copy bytes from a string to a slice of bytes.)
	return *(*string)(unsafe.Pointer(&b)) // 实现了 []byte 到 string 的零内存拷贝转换
}