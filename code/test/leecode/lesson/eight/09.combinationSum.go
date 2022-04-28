package eight
/*
https://leetcode-cn.com/problems/combination-sum/

组合总和

给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为target 的不同组合数少于 150 个。



示例1：
	输入：candidates = [2,3,6,7], target = 7
	输出：[[2,2,3],[7]]
	解释：
		2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
		7 也是一个候选， 7 = 7 。
		仅有这两种组合。

示例2：	
	输入: candidates = [2,3,5], target = 8
	输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
	输入: candidates = [2], target = 1
	输出: []

提示：
	1 <= candidates.length <= 30
	1 <= candidates[i] <= 200
	candidate 中的每个元素都 互不相同
	1 <= target <= 500
*/
var resultCombinationSum [][]int
func combinationSum(candidates []int, target int) [][]int {
	resultCombinationSum = [][]int{}
	path := []int{}
	backtrackCombinationSum(candidates, 0, path, target)
	return resultCombinationSum
}
/*
candidates 可选列表
k 阶段
path 路径
left 剩余总和
*/
func backtrackCombinationSum(candidates []int, k int, path []int, left int) {
	if left == 0 { // 结束条件,加入结果集
		temp := make([]int, len(path))
		copy(temp, path)
		resultCombinationSum = append(resultCombinationSum, temp)
		return
	}
	if k == len(candidates) { // 遍历结束,没有找到符合条件的结果集
		return
	}
	for i := 0; i <= left/candidates[k]; i++ { // 决策阶段
		for j := 0; j < i; j++ {
			path = append(path, candidates[k])
		}
		backtrackCombinationSum(candidates, k+1, path, left - i*candidates[k])
		for j := 0; j < i; j++ {
			path = path[:len(path)-1]
		}
	}
}