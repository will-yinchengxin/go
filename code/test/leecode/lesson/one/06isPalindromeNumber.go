package one

import "strconv"

func IsPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(s) - 1
	for left < right {         // 优化写法 for left < right && !judgeStr(s[left]) {}
		// 找到左右第一个符合要求的字符
		for left < right {
			if judgeStr(s[left]) {
				break
			}
			left++
		}
		for left < right {     // 优化写法 for left < right && !judgeStr(s[right]) {}
			if judgeStr(s[right]) {
				break
			}
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}