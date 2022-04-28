package one

import "strings"

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串
*/
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
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

func judgeStr(str byte) bool {
	return (str >= 'a' && str <= 'z') || (str >= '0' && str <= '9')
}
