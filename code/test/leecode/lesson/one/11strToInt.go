package one

import (
	"math"
	"strings"
)

/*
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。

示例1:

输入: "42"
输出: 42
示例2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
    我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字
*/
func StrToInt(str string) int {

	/*
			字符转数字： “此数字的 ASCII 码” 与 “ 0 的 ASCII 码” 相减即可；
			数字拼接： 若从左向右遍历数字，设当前位字符为 c ，当前位数字为 x ，数字结果为 res ，则数字拼接公式为：
			x=ascii(c)−ascii(′0′)
			res=10×res+x

			大小范围：
				< 2147483647 math.MinInt32
		        > - 2147483648
	*/
	str = strings.TrimSpace(str)
	var result  int = 0
	sign := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if i == 0 && v == '+' {
			sign = 1
		} else if i == 0 && v == '-' {
			sign = -1
		} else {
			break
		}
		// int   类型大小为 8 字节  '
		// int64 类型大小为 8 字节
		// int64:  -9223372036854775808 ~ 9223372036854775807
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return result*sign
}