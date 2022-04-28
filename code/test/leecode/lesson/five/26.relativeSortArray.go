package five

import "sort"

/*
给你两个数组，arr1 和arr2，

arr2中的元素各不相同
arr2 中的每个元素都出现在arr1中
对 arr1中的元素进行排序，使 arr1 中项的相对顺序和arr2中的相对顺序相同。未在arr2中出现过的元素需要按照升序放在arr1的末尾。

示例：
	输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
	输出：[2,2,2,1,4,3,3,9,6,7,19]

提示：
	1 <= arr1.length, arr2.length <= 1000
	0 <= arr1[i], arr2[i] <= 1000
	arr2中的元素arr2[i]各不相同
	arr2 中的每个元素arr2[i]都出现在arr1中
*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	lenArr2 := len(arr2)
	lenArr1 := len(arr1)

	hashMap := make(map[int]int, lenArr1)
	for i := 0; i < lenArr1; i++ {
		if _, ok := hashMap[arr1[i]]; !ok {
			hashMap[arr1[i]] = 1
		} else {
			hashMap[arr1[i]]++
		}
	}

	res := make([]int, 0, lenArr1)
	for i := 0; i < lenArr2; i++ {
		if val, ok := hashMap[arr2[i]]; ok {
			for j := 0; j < val; j++ {
				res = append(res, arr2[i])
			}
			delete(hashMap, arr2[i])
		}
	}

	if len(hashMap) > 0 {
		leftArr := []int{}
		for i, val := range hashMap {
			for j := 0; j < val; j++ {
				leftArr = append(leftArr, i)
			}
		}
		sort.Ints(leftArr)
		res = append(res, leftArr...)
	}
	return res
}