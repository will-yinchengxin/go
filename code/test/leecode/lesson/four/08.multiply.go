package four

import "math"

/*
递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。

示例1:
	 输入：A = 1, B = 10
	 输出：10

示例2:
	 输入：A = 3, B = 4
	 输出：12

提示:
	保证乘法范围不会溢出
*/
func multiply(A int, B int) int {
	/*
		1) 判断符号位, 正数或者负数
		2) 求两个数的绝对值
		3) 以值较小的作为递归的次数
		4) 每次 min/2, 纸质 min/2 为 1, 返回 max (min 个 max 相加 = (min/2 个 max) + (min/2 个 max) + (max 或者 0))
		5) 判断 min%2 的值, 如果为 1 说明为奇数个, 后置位 + max, 如果为 0, 则为 偶数个, 后置位不变
	*/

	// 判断符号
	tag := false
	if (A < 0 && B > 0) || (A > 0 && B < 0) {
		tag = true
	}
	A = int(math.Abs(float64(A)))
	B = int(math.Abs(float64(B)))
	res := multiply_r(A, B)
	if tag {
		return 0 - res
	}
	return res
}

func multiply_r(A int, B int) int {
	// 小的数值作为递归的次数, 如果为 1, 直接返回结果
	min := int(math.Min(float64(A), float64(B)))
	max := int(math.Max(float64(A), float64(B)))
	if min == 1 {
		return max
	}
	// min 个 max 相加 = (min/2 个 max) + (min/2 个 max) + (max 或者 0)
	halfVal := multiply_r(min/2, max)
	if min % 2 == 1 {
		return halfVal + halfVal + max
	} else {
		return  halfVal + halfVal
	}
}