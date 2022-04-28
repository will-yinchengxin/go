package three
/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/
func maxSlidingWindow(nums []int, k int) []int {
	lenN := len(nums)
	if lenN < 1 {
		return nil
	}
	i := 0
	stack := []int{}
	for k <= lenN {
		val := getMax(nums[i:k])
		stack = append(stack, val)
		i++
		k++
	}

	return stack
}

func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}