package test

import (
	"fmt"
	"sort"
	"testing"
)

// 函数要以 Benchmark 为前缀
// -bench 指定名称则只执行具体测试方法而不是全部  . 则是对所有的benchmark函数测试
// count 测试的次数

//  go test -bench=test -benchmem -count=2
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		s := make([]int, 0)
		for i := 0; i < 10000; i++ {
			s = append(s, i)
		}
	}
}

// go test -bench=will -benchmem
func BenchmarkWill(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

// 对一个[]string做排序的内存分配压测
func BenchmarkSortStrings(b *testing.B) {
	s := []string{"heart", "lungs", "brain", "kidneys", "pancreas"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var ss sort.StringSlice = s
		var si sort.Interface = ss
		sort.Sort(si)
	}
}

/*
goos: windows
goarch: amd64
pkg: test/test
Benchmark_will-12       18181680                64.5 ns/op             2 B/op          1 allocs/op
PASS
ok      test/test       1.856s

-12 		 表示GOMAXPROCS（线程数）的值为12
18181680 	 表示执行了 18181680次
64.5 ns/op 	 表示每次操作花费 64.5 纳秒
2 B/op  	 表示每次申请了 2B的内存
1 allocs/op	 表示每次操作申请了1次内存
*/