package leecode

import (
	"math"
	"strings"
)

/**
字符串转数字(atoi)
**/
func strToInt(str string) int {
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
	result := 0
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
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return result*sign
}

/*
数字转数组
 例: 153
	153 % 10 = 3
	(153 / 10) % 10 = 5
	if (被除数 < 10) {
		直接输出
	}
*/

/**
数字转字符串(itoa)  var arr = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} string(arr)
*/

