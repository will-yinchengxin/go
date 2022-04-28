package one
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度

示例1:
	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
	输入: s = "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
	输入: s = "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
	请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

示例 4:
	输入: s = ""
	输出: 0
*/
func lengthOfLongestSubstring(s string) int {
	lenS := len(s)
	left := -1
	res := 0 // 记录最长字符串长度
	window := make(map[byte]int, 0) // 存储最终的结果集

	for right := 0; right < lenS; right++ {
		ch := s[right]
		window[ch]++

		for window[ch] > 1{
			left++
			window[s[left]]--
		}

		res = max(res, right - left)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



// 最长区间刷题模板
// func module(A []int) int {
// 	n := len(A)
// 	left, ans := -1, 0
// 	for right := 0; right < n; right++ {
// 		//  1. 直接将A[right]加到区间中，形成(left, right]
// 		//  2. 将A[right]加入后，惰性原则
// 		for check((left, right]){		//  TODO，检查区间是否满足条件
// 			left++		//  不满足条件，移动左指针
// 			//  TODO 修改区间状态
// 		}
//
// 		//  assert 此时(left, right]必然满足条件
// 		ans = max(ans, i - left)
// 	}
// }

