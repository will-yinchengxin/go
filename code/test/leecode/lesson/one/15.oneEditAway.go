package one

/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

示例1:
	输入:
	first = "pale"
	second = "ple"
	输出: True


示例2:
	输入:
	first = "pales"
	second = "pal"
	输出: False
*/
func OneEditAway(first string, second string) bool {
	/*
		0) 长度不等, 允许的最大差为 1, > 1 直接返回
		1) 长度相等逐个遍历对比, 允许不同字符数量 <= 1
		2) 删除 和 编辑的思路一致
	*/
	shortStr := ""
	longStr := ""
	if len(first) > len(second) {
		shortStr = second
		longStr = first
	} else {
		shortStr = first
		longStr = second
	}
	if len(longStr) - len(shortStr) > 1 {
		return false
	}
	if len(longStr) == 0 || len(shortStr) == 0 {
		return true
	}
	// 相等情况
	if len(longStr) == len(shortStr) {
		equalTag := 0
		for i := 0; i < len(longStr); i++ {
			if equalTag > 1 {
				return false
			}
			if longStr[i] != shortStr[i] {
				equalTag++
			}
		}
		return equalTag < 2
	}
	// 不相等
	// "islands" "aislands"
	// "ab" "a"
	// "ab" "acb"
	// "dskjlfj" "dskjlfja"
	logPos := 0
	shortPos := 0
	noEqualTag := 0
	for noEqualTag < 2 {
		if len(shortStr) == shortPos {
			break
		}
		if shortStr[shortPos] != longStr[logPos] {
			noEqualTag++
		} else {
			shortPos++
		}
		logPos++
	}
	return shortPos + 1 >= logPos
}
