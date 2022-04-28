package five
// https://leetcode-cn.com/problems/binary-search/

// 无重复的数组中查找 第一个等于给定值的元素
func search(nums []int, target int) int {
	lenN := len(nums)
	low := 0
	high := lenN - 1
	for low <= high {
		mid := low + (high - low) / 2 // 避免数组越界
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return - 1
}