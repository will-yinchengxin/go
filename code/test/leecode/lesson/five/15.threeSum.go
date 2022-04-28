package five

import "sort"

/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]

示例 2：
	输入：nums = []
	输出：[]

示例 3：
	输入：nums = [0]
	输出：[]

提示：
	0 <= nums.length <= 3000
	-10^5 <= nums[i] <= 10^5
*/
func threeSum(nums []int) [][]int {
	/*
		1) 先对切片进行一个排序
		2) 将切片的值作为键, 键作为值(重复覆盖保留的是最后一个元素)
	*/
	sort.Ints(nums)

	lenN := len(nums)
	hashMap := make(map[int]int)
	for i := 0; i < lenN; i++ {
		hashMap[nums[i]] = i
	}

	var res [][]int // 存放结果

	// 双层循环
	for i := 0; i < lenN; i++ {
		if i != 0 && nums[i] == nums[i-1] { // 答案中不可以包含重复的三元组。避免 元素重复
			continue
		}
		for j := i+1; j < lenN; j++ {
			if j != i+1 && nums[j] == nums[j-1] { // 答案中不可以包含重复的三元组。避免 元素重复
				continue
			}
			// 这里处理的为 0 的情况, 其他情况依然使用
			target := -1*(nums[i] + nums[j])
			val, ok := hashMap[target]
			if !ok {
				continue
			}
			if val > j {
				var childRes []int
				childRes = append(childRes, nums[i])
				childRes = append(childRes, nums[j])
				childRes = append(childRes, nums[val])
				res = append(res, childRes)
			}
		}
	}
	return res
}