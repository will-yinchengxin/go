package string

import "testing"
/*
执行命令: go test -bench=BenchmarkEqual -benchmem
		go test -bench=BenchmarkEqual -benchmem | grep 'ns/op'

ns/op:		每次操作花费的时间
B/op: 		申请的内存
allocs/op: 	每次操作申请了多少次内存
*/

// 2.71 ns/op (每次操作时间)
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String("This is a sample string", "this is a sample string")
	}
}

// 5.51 ns/op (每次操作时间)
func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("This is a sample string", "this is a sample string")
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualFold("This is a sample string", "this is a sample string")
	}
}