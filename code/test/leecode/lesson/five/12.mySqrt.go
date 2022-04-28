package five
/*
给你一个非负整数 x ，计算并返回x的算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

示例 1：
	输入：x = 4
	输出：2

示例 2：
	输入：x = 8
	输出：2
	解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

提示：
	0 <= x <= 231 - 1
*/
func mySqrt(x int) int {
	/*
		1) 在 [0, x] 区间中寻找最后一个小于等于 x 的平方数
	*/
	low := 0
	high := x

	for low <= high {
		mid := low + (high-low)/2
		if mid * mid <= x {
			if mid == x || (mid+1) * (mid+1) > x {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return - 1
}