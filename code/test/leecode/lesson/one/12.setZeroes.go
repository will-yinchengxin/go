package one

import (
	"fmt"
)

/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

示例 1：
	输入：
	[
	  [1,1,1],
	  [1,0,1],
	  [1,1,1]
	]
	输出：
	[
	  [1,0,1],
	  [0,0,0],
	  [1,0,1]
	]

示例 2：
	输入：
	[
	  [0,1,2,0],
	  [3,4,5,2],
	  [1,3,1,5]
	]
	输出：
	[
	  [0,0,0,0],
	  [0,4,5,0],
	  [0,3,1,0]
	]
*/


func SetZeroes(matrix [][]int)  {
	/*
		1） 记录数组长度 n
		2） 记录每个子集长度 s
		3） 利用 hashMap 记录为 0 的下标， 如果存在直接跳过
		4） 最后组装输出的数组
	*/
	n := len(matrix)
	s := len(matrix[0])
	if n == 0 || s == 0{
		return
	}
	// 记录横列
	transverseMap := make(map[int]bool, n)
	// 记录纵列
	verticalMap := make(map[int]bool, s)

	for i, i2 := range matrix {
		for i3, i4 := range i2 {
			if i4 == 0 {
				if _, ok := transverseMap[i]; !ok {
					transverseMap[i] = true
				}
				verticalMap[i3] = true
			}
			if _, ok := verticalMap[i3]; ok {
				continue
			}
			if len(verticalMap) == s || len(transverseMap) == n {
				break
			}
		}
	}
	//fmt.Println(transverseMap, verticalMap)
	// 生成最终矩阵
	for i, inti := range matrix {
		for j, _ := range inti {
			_, ok1 := transverseMap[i]
			_, ok2 := verticalMap[j]
			if ok1 || ok2 {
				inti[j] = 0
			}
		}
	}
	fmt.Println(matrix)
}
