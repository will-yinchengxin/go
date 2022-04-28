package five
/*
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false
进阶：不要 使用任何内置的库函数，如sqrt

示例 1：
	输入：num = 16
	输出：true

示例 2：
	输入：num = 14
	输出：false

提示：
	1 <= num <= 2^31 - 1
*/
func isPerfectSquare(num int) bool {
	/*
		1) 考察 [1, num] 是否存在一个数字, 使得 X^2 = num
	*/
	low := 0
	high := num
	for low <= high {
		mid := low + (high- low)/2
		if mid * mid == num {
			return true
		} else if (mid * mid) > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}