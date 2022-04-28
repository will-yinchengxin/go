package four

import "sort"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

示例1:
	输入: s = "anagram", t = "nagaram"
	输出: true

示例 2:
	输入: s = "rat", t = "car"
	输出: false


提示:
	1 <= s.length, t.length <= 5 * 104
	s 和 t仅包含小写字母
*/
// 对两个字母进行排序
func IsAnagram(s string, t string) bool {
	// 字符串转数组
	s1, s2 := []byte(s), []byte(t)

	lenS1 := len(s1)
	lenS2 := len(s2)
	if lenS1 != lenS2 {
		return false
	}

	//quickSortByte(s1, len(s1))
	//quickSortByte(s2, len(s2))
	//fmt.Println(s1, s2)

	// 对数组进行排序
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	return string(s1) == string(s2)
}

func quickSortByte(a []byte, n int) {
	quickSortByte_r(a, 0, n-1)
}
func quickSortByte_r(a []byte, low, high int) {
	if low >= high {
		return
	}
	point := partitionByte(a, low, high)
	quickSortByte_r(a, low, point-1)
	quickSortByte_r(a, point+1, high)
}
func partitionByte(a []byte, low, high int) int {
	// 这里取最后一个节点作为分区点
	i, j :=  low, high -1
	for {
		for a[i] < a[high] {
			i++
			if i == high {
				break
			}
		}
		for a[j] > a[high] {
			j--
			if j == low {
				break
			}
		}
		if j <= i {
			break
		}
		swapByte(a, i, j)
	}
	swapByte(a, i, high)
	return i
}
func swapByte(a []byte, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}