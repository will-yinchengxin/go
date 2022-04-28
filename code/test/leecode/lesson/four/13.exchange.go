package four

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

示例：
	输入：nums = [1,2,3,4]
	输出：[1,3,2,4]
	注：[3,1,2,4] 也是正确的答案之一。

提示：
	0 <= nums.length <= 50000
	0 <= nums[i] <= 10000
*/
func Exchange(nums []int) []int {
	lenN := len(nums)
	if lenN <= 1 {
		return nums
	}
	i, j := 0, lenN - 1
	for {
		for nums[i] % 2 == 1 {
			i++
			if i == lenN - 1 {
				break
			}
		}
		for nums[j] % 2 == 0{
			j--
			if j == 0 {
				break
			}
		}
		if i >= j  {
			break
		}
		swapExchange(nums, i, j)
	}
	return nums
}
func swapExchange(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}