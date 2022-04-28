package selice

import (
	"fmt"
	"sort"
)

/*
	切片可以看成是对数组的包装形式, 它包装的数组称为该切片的底层数组
	反之: 切片是针对其底层数组中某个连续片段的描述
	如: var ips = []string{"192.168.1.0", "192.168.1.1"}

	切片的长度是可变的, 且并不是类型的一部分, 只要元素类型相同, 两个切片的类型就是一致的

	一个切片的类型的零值总为 nil, 此 切片 的 长度 和 容量 都为 0

	切片包含三个部分: 1)指向底层数组的指针  2)切片的长度  3)切片的容量
	切片容量: 再不更换底层数组的前提下, 底层数组长度的最大值

	append 添加元素的时候, 有可能改变底层数组的长度,引发扩容,生成新的底层数组,此时 新,旧切片 也就有可能指向不同的底层数组

	make 函数初始化一个 切片(还可以初始化 channel, map), 用make函数初始化的切片值中的每一个元素值都会是其元素类型的零值，这里ips中的那100个元素的值都会是空字符串""
*/

// 声明一个二维切片
func NewSelice(lenMatrix int) {
	newMatrix := make([][]int, lenMatrix)
	for i := range newMatrix {
		newMatrix[i] = make([]int, lenMatrix)
	}
}

// 切片复制
func CopySelice() {
	// copy方式
	a := []int{123,123,4324}
	b := make([]int, len(a))  // 一次将内存申请到位
	copy(b, a)

	// append 方式, 这种方式较慢, 但是后续如果要添加更多的元素, 这种效率较快
	b = append([]int(nil), a...)
}

// 切片删除
func DeleteSelice() {
	a := []int{123,123,4324}
	copy(a[1:], a[1+1:])
	a[len(a)-1] = 0 // 或类型T的零值
	a = a[:len(a)-1]
}

func CutSelice() {
	a := []int{123,123,4324,123,54,2313,123}
	i := 2
	j := 5
	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = 0 // 或类型T的零值
	}
	a = a[:len(a)-j+i]
}

// 数组去重
func UniqueSelice() {
	in := []int{3,2,1,4,3,2,1,4,1} // 切片元素可以是任意可排序的类型
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// 需要保存原始数据时
		// in[i], in[j] = in[j], in[i]
		// 只需要保存需要的数据时
		in[j] = in[i]
	}
	result := in[:j+1]
	fmt.Println(result)
}

// 分批次处理一个切片
func BatchByBatchDealSelice() {
	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3
	batches := make([][]int, 0, (len(actions) + batchSize - 1) / batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions) // [[0 1 2] [3 4 5] [6 7 8] [9]]
}