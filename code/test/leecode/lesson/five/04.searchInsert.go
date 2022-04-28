package five

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
	输入: nums = [1,3,5,6], target = 5
	输出: 2

示例2:
	输入: nums = [1,3,5,6], target = 2
	输出: 1

示例 3:
	输入: nums = [1,3,5,6], target = 7
	输出: 4

示例 4:
	输入: nums = [1,3,5,6], target = 0
	输出: 0

示例 5:
	输入: nums = [1], target = 0
	输出: 0

提示:
	1 <= nums.length <= 10^4
	-104 <= nums[i] <= 10^4
	nums 为无重复元素的升序排列数组 !!!
	-10^4 <= target <= 10^4

[1,3,5]
4
[1,2,3]
3 形同插入最前面
*/
func searchInsert(nums []int, target int) int {
	lenN := len(nums)
	low := 0
	high := lenN - 1
	for low <= high {
		mid := low + (high - low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target  {
			if mid != low && nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			if mid != high && nums[mid+1] > target {
				return mid + 1
			}
			low = mid + 1
		}
	}
	if target > nums[lenN - 1] {
		return lenN
	}
	return 0
}
