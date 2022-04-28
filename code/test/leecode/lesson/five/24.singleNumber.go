package five

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
	输入: [2,2,1]
	输出: 1

示例2:
	输入: [4,1,2,1,2]
	输出: 4
*/
func singleNumber(nums []int) int {
	lenN := len(nums)
	hashMap := make(map[int]int)
	for i := 0; i < lenN; i++ {
		if _, ok := hashMap[nums[i]]; !ok {
			hashMap[nums[i]] = 1
		} else {
			hashMap[nums[i]]++
		}
	}
	for i, val := range hashMap {
		if val == 1 {
			return i
		}
	}
	return -1
}