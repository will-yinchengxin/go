package leecode

/*
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

// 贪心算法
func maxSubArray(nums []int) int {
	lenNum := len(nums)
	if lenNum == 0 {
		return 0
	}
	m := nums[0]
	tmp := 0
	for i := 0; i < lenNum; i++ {
		tmp += nums[i]
		if tmp > m {
			m = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return m
}