package five

import "sort"

/*
https://leetcode-cn.com/problems/valid-anagram/
*/

// 使用 hash 解决
func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}
	hashMap := make(map[byte]int)
	for i := 0;  i < lenS; i++ {
		if _, ok := hashMap[s[i]]; !ok {
			hashMap[s[i]] = 1
		} else {
			hashMap[s[i]]++
		}
	}

	for i := 0; i < lenT; i++ {
		if len(hashMap) <= 0 {
			return false
		}
		if _, ok := hashMap[t[i]]; ok {
			if hashMap[t[i]] > 1 {
				hashMap[t[i]]--
			} else {
				delete(hashMap, t[i])
			}
		}
	}
	if len(hashMap) > 0 {
		return false
	}
	return true
}

// 对字符排序
func isAnagramAno(s string, t string) bool {
	// 字符串转数组
	s1, s2 := []byte(s), []byte(t)
	lenS1 := len(s1)
	lenS2 := len(s2)
	if lenS1 != lenS2 {
		return false
	}

	// 对数组进行排序
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}