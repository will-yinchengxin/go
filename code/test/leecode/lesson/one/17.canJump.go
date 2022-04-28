package one

/*
给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
*/
/*
示例1：
	输入：nums = [2,3,1,1,4]
	输出：true
	解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例2：
	输入：nums = [3,2,1,0,4]
	输出：false
	解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：
	1 <= nums.length <= 3 * 104
	0 <= nums[i] <= 105
*/
func CanJump(nums []int) bool {
	/*
	0) 如果没有 0 一定可以跳过去
	1) 有0可能会跳过去, 有可能跳不回去

	2) 统计每个元素最大的可达距离,
	存在一个位置 x, 它本身可到达, 它的跳跃长度为 x + nums[x], 只要 x + nums[x] >= y (lenNum - 1) 则找到题解
	*/
	//lenNum := len(nums)
	//if lenNum == 0 {
	//	return false
	//}
	//rightMost := 0
	//for i := 0; i < lenNum; i++ {
	//	if i <= rightMost {
	//		if rightMost < i + nums[i] {
	//			rightMost = i + nums[i]
	//		}
	//		if rightMost >= lenNum - 1 {
	//			return true
	//		}
	//	}
	//}
	//return false
	lenNum := len(nums)
	if lenNum <= 0 {
		return false
	}
	if lenNum == 1 {
		return true
	}
	// 0 1 0 1 4
	reachMax := 0
	for i := 0; i < lenNum; i++ {
		if reachMax >= i {
			if reachMax < i + nums[i] {
				reachMax = i + nums[i]
			}
			if reachMax >= lenNum - 1 {
				return true
			}
		}
	}
	return false
}

/*
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
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
		// 点睛之笔
		if tmp < 0 {
			tmp = 0
		}
	}
	return m
}