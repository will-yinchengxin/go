package five

import "strings"

/*
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:
	 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
	 输出：-1
	 说明: 不存在返回-1。

示例2:
	 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
	 输出：4

提示:
	words的长度在[1, 1000000]之间
*/
// 这里比较字符串的时候要使用 strings.Compare 函数
func findString(words []string, s string) int {
	lenN := len(words)
	low := 0
	high := lenN - 1

	target := -1
	for low <= high {
		mid := low + (high - low)/2
		if strings.Compare(words[mid], s) == 0 {
			return mid
		} else if words[mid] == "" {
			// 让 low 或者 high 自增
			if strings.Compare(words[low], s) == 0{
				return low
			}
			if strings.Compare(words[high], s) == 0{
				return high
			}
			low++
			high--
		} else if strings.Compare(words[mid], s) == 1 {
			high = mid - 1
		} else if strings.Compare(words[mid], s) == -1 {
			low = mid + 1
		}
	}
	return target
}
