package five

/*
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：
	输入: s1 = "abc", s2 = "bca"
	输出: true

示例 2：
	输入: s1 = "abc", s2 = "bad"
	输出: false

说明：
	0 <= len(s1) <= 100
	0 <= len(s2) <= 100
*/
func CheckPermutation(s1 string, s2 string) bool {
	hashMap := make(map[byte]int)
	for i, lenS1 := 0, len(s1); i < lenS1; i++ {
		if _, ok := hashMap[s1[i]]; !ok {
			hashMap[s1[i]] = 1
		} else {
			hashMap[s1[i]]++
		}
	}
	for i, lenS2 := 0, len(s2); i < lenS2; i++ {
		if len(hashMap) <= 0 {
			return false
		}
		if _, ok := hashMap[s2[i]]; ok {
			if hashMap[s2[i]] > 1 {
				hashMap[s2[i]]--
			} else {
				delete(hashMap, s2[i])
			}
		}
	}
	if len(hashMap) > 0 {
		return false
	}
	return true
}