package eight
/*
https://leetcode-cn.com/problems/combination-sum-ii/

组合总和 II

给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合。

示例1:
	输入: candidates =[10,1,2,7,6,1,5], target =8,
	输出:
		[
		[1,1,6],
		[1,2,5],
		[1,7],
		[2,6]
		]

示例2:
	输入: candidates =[2,5,2,1,2], target =5,
	输出:
		[
		[1,2,2],
		[5]
		]

提示:
	1 <=candidates.length <= 100
	1 <=candidates[i] <= 50
	1 <= target <= 30
*/
var resultCombinationSum2 [][]int
func combinationSum2(candidates []int, target int) [][]int {
	/*
		因为要求每个数字只能使用一次,我们这里还需要统计一下每个 数字 和其 频率
		uniqueNums 和 counts
	*/
	hashMap := make(map[int]int)
	for i := 0; i < len(candidates); i++ {
		count := 1
		if _, ok := hashMap[candidates[i]]; ok {
			count += hashMap[candidates[i]]
		}
		hashMap[candidates[i]] = count
	}
	lenHM := len(hashMap)
	uniqueNums := make([]int, 0, lenHM)
	counts := make([]int, 0, lenHM)
	for i := 0; i < len(candidates); i++ {
		if val, ok := hashMap[candidates[i]]; ok {
			uniqueNums = append(uniqueNums, candidates[i])
			counts = append(counts, val)
			delete(hashMap, candidates[i])
		}
	}

	resultCombinationSum2 = [][]int{}
	path := []int{}
	backtrackCombinationSum2(uniqueNums, counts, 0, path, target)
	return resultCombinationSum2
}
/*
	uniqueNums, counts 通过循环推到出当前的可选列表
	k 阶段
	path 路径
	left 升序和 即 path 内元素和
*/
func backtrackCombinationSum2(uniqueNums, counts []int, k int, path []int, left int) {
	if left == 0 { // 结束条件, 加入结果集
		temp := make([]int, len(path))
		copy(temp, path)
		resultCombinationSum2 = append(resultCombinationSum2, temp)
		return
	}
	if k == len(uniqueNums) || left < 0 { // 遍历完成所有元素,未查询到正确结果集
		return
	}
	for count := 0; count <= counts[k]; count++ { // 决策阶段
		for i := 0; i < count; i++ {
			path = append(path, uniqueNums[k])
		}
		backtrackCombinationSum2(uniqueNums, counts, k+1, path, left-count*uniqueNums[k])
		for i := 0; i < count; i++ {
			path = path[:len(path)-1]
		}
	}
}