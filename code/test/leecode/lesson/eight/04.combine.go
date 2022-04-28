package eight
/*
https://leetcode-cn.com/problems/combinations/

组合

给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

示例 1：
	输入：n = 4, k = 2
	输出：
	[
	  [2,4],
	  [3,4],
	  [2,3],
	  [1,2],
	  [1,3],
	  [1,4],
	]

示例 2：
	输入：n = 1, k = 1
	输出：[[1]]

提示：
	1 <= n <= 20
	1 <= k <= n
*/
var resultCombine [][]int
func combine(n int, k int) [][]int {
	resultCombine = [][]int{}
	path := []int{}
	backtrackCombine(n, 1, path, k)
	return resultCombine
}
/*
	n 可选列表的截止数据
	start 当前阶段
	path 路径
	k 标记量
*/
func backtrackCombine(n, start int, path []int, k int) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		resultCombine = append(resultCombine, tmp)
		return
	}
	if start == n+1 { // 提前退出
		return
	}

	// 因为这里只有两种情况, 可以放置, 可以不放置
	// 可以使用 for 循环, 也是可以不使用

	// 不适用 for 循环
	backtrackCombine(n, start+1, path, k) // 不放置

	path = append(path, start) // 放置
	backtrackCombine(n, start+1, path, k)
	path = path[:len(path)-1] // 放置
	// 使用 for 循环写法
	//for i := start; i <= n; i++ {
	//	path = append(path, i)
	//	backtrackCombine(n, i+1, path, k)
	//	path = path[:len(path)-1]
	//}
}