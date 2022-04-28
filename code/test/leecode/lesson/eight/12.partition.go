package eight
/*
https://leetcode-cn.com/problems/palindrome-partitioning/

分割回文串

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

示例 1：
	输入：s = "aab"
	输出：[["a","a","b"],["aa","b"]]

示例 2：
	输入：s = "a"
	输出：[["a"]]

提示：
	1 <= s.length <= 16
	s 仅由小写英文字母组成
*/
var resultPartition [][]string
func partition(s string) [][]string {
	resultPartition = [][]string{}
	path := []string{}
	backtrackPartition(s, 0, path)
	return resultPartition
}
/*
originalStr 原始字串
k 阶段
path 路径
*/
func backtrackPartition(originalStr string, k int, path []string) {
	if k == len(originalStr) { // 结束条件
		temp := make([]string, len(path))
		copy(temp, path)
		resultPartition = append(resultPartition, temp)
		return
	}
	// 决策阶段
	for i := k; i < len(originalStr); i++ {
		if judgePartition(originalStr, k, i) { // 判断是否为回文串
			path = append(path, originalStr[k:i+1]) // 这里注意字符串长度范围问题
			backtrackPartition(originalStr, i+1, path)
			path = path[:len(path)-1]
		}
	}
}
// 判别回文也可以借助栈来实现
func judgePartition(originalStr string, start, end int) bool {
	i := start
	j := end
	for i <= j {
		if (originalStr[i] != originalStr[j]) {
			return false
		}
		i++
		j--
	}
	return true
}