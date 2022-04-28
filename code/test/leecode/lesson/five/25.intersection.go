package five
/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
	输入：nums1 = [1,2,2,1], nums2 = [2,2]
	输出：[2]

示例 2：
	输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	输出：[9,4]

说明：
	输出结果中的每个元素一定是唯一的。
	我们可以不考虑输出结果的顺序。
*/
func intersection(nums1 []int, nums2 []int) []int {
	lenN1 := len(nums1)
	lenN2 := len(nums2)

	hashMap := make(map[int]struct{})
	if lenN1 > lenN2 {
		for i := 0; i < lenN2; i++ {
			if _, ok := hashMap[nums2[i]]; !ok {
				hashMap[nums2[i]] = struct{}{}
			}
		}
	} else {
		for i := 0; i < lenN1; i++ {
			if _, ok := hashMap[nums1[i]]; !ok {
				hashMap[nums1[i]] = struct{}{}
			}
		}
	}

	resMap := make(map[int]struct{})
	if lenN1 > lenN2 {
		for i := 0; i < lenN1; i++ {
			if _, ok := hashMap[nums1[i]]; ok {
				resMap[nums1[i]] = struct{}{}
			}
		}
	} else {
		for i := 0; i < lenN2; i++ {
			if _, ok := hashMap[nums2[i]]; ok {
				resMap[nums2[i]] = struct{}{}
			}
		}
	}
	res := make([]int, 0, len(resMap))
	for i, _ := range resMap {
		res = append(res, i)
	}
	return res
}