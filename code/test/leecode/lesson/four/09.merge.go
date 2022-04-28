package four

import (
	"sort"
)

/*
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 m 和 n。

示例:
	输入:
		A = [1,2,3,0,0,0], m = 3
		B = [2,5,6],       n = 3
	输出:
		[1,2,2,3,5,6]

说明:
	A.length == n + m
*/
func mergeA(A []int, m int, B []int, n int)  {
	copy(A[m:], B)
	sort.Ints(A)
}
// 双指针
// 时间复杂度 O(m+n)
// 空间复杂度O(m+n)
func MergeB(A []int, m int, B []int, n int) {
	lenA := len(A)
	sortArr := make([]int, 0, lenA)
	i, j := 0, 0
	for i < m && j < n {
		if A[i] < B[j] {
			sortArr = append(sortArr, A[i])
			i++
		} else {
			sortArr = append(sortArr, B[j])
			j++
		}
	}
	sortArr = append(sortArr, A[i:m]...)
	sortArr = append(sortArr, B[j:n]...)
	copy(A[0:], sortArr)
}

// 双指针, 从末尾向头部开始遍历
// [1,2,3,0,0,0] 3 -> 2 -> 1, [2,5,6] 6 -> 5 -> 2
// 时间复杂度 O(m+n)
// 空间复杂度O(1)
func MergeC(A []int, m int, B []int, n int) {
	lenA := len(A)
	i, j, end := m-1, n-1, lenA-1
	for i >= 0 && j >= 0 && end >= 0 {
		if A[i] > B[j] {
			A[end] = A[i]
			i--
		} else {
			A[end] = B[j]
			j--
		}
		end--
	}
	for i >= 0 && end >= 0 {
		A[end] = A[i]
		i--
		end--
	}
	for j >= 0 && end >= 0 {
		A[end] = B[j]
		j--
		end--
	}
}