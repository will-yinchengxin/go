package eight
/*
https://leetcode-cn.com/problems/combination-sum-iii/

组合总和 III

找出所有相加之和为n的k个数的组合。组合中只允许含有 1-9 的正整数，并且每种组合中不存在重复的数字。

说明：
	所有数字都是正整数。
	解集不能包含重复的组合。

示例 1:
	输入: k = 3, n = 7
	输出: [[1,2,4]]

示例 2:
	输入: k = 3, n = 9
	输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
var resultCombinationSum3 [][]int
func combinationSum3(k int, n int) [][]int {
	resultCombinationSum3 = [][]int{}
	path := []int{}
	backtrackCombinationSum3(k, n, path, 1)
	return resultCombinationSum3
}
func backtrackCombinationSum3(k, left int, path []int, start int) {
	if left == 0 && len(path) == k { // 结束条件
		temp := make([]int, len(path))
		copy(temp, path)
		resultCombinationSum3 = append(resultCombinationSum3, temp)
		return
	}
	if len(path) > k || left < 0 || start > 9 {
		return
	}
	// 决策阶段(0, 1背包问题)
	backtrackCombinationSum3(k, left, path, start+1) // 不添加元素
	path = append(path, start) // 添加元素
	backtrackCombinationSum3(k, left-start, path, start+1)
	path = path[:len(path)-1]
}