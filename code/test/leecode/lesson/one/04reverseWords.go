package one

import (
	"strings"
)

func reverseWords(s string) string {
	res := strings.Fields(s)

	length := len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return strings.Join(res, " ")
}

// 双指针:
// 索引从右边向左搜索首个空格
// 添加单词 s[i+1:j+1]
// 索引 i 跳过两个单词间的所有空格
// 执行 j=i 此时 j 指向单词的尾字符
func reverseWordsAno(s string) string {
	s = strings.Trim(s, " ")
	lenS := len(s)
	if lenS <= 1 {
		return s
	}
	str := ""
	for k, j := lenS-1, lenS-1; k >= 0; k-- {
		if s[k] == ' ' && k != j {
			str = str + s[k+1:j+1] + " "
			j = k-1
			continue
		}
		if s[k] == ' ' && s[j] == ' ' {
			j = k-1
			continue
		}
		if k == 0 {
			str = str + s[k:j+1]
		}
	}
	return str
}