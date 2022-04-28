package eight
/*
https://leetcode-cn.com/problems/subsets-ii/

子集 II

给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
	输入：nums = [1,2,2]
	输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

示例 2：
	输入：nums = [0]
	输出：[[],[0]]

提示：
	1 <= nums.length <= 10
	-10 <= nums[i] <= 10
*/
var resultSubsetsWithDup [][]int
func subsetsWithDup(nums []int) [][]int {
	/*
		可以得到结果集后,对比每一个数组是否重复,进行剃重
		但是这种 容器 与 容器 间的判重 kind of silly

		我们可以 就每个数字出现的频率做一个统计
	*/
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count := 1
		if _, ok := hashMap[nums[i]]; ok {
			count += hashMap[nums[i]]
		}
		hashMap[nums[i]] = count
	}
	lenHM := len(hashMap)
	uniqueNums := make([]int, 0, lenHM) // 记录数字
	counts := make([]int, 0, lenHM) // 记录数字出现的频率
	for i := 0; i < len(nums); i++ {
		if val, ok := hashMap[nums[i]]; ok {
			uniqueNums = append(uniqueNums, nums[i])
			counts = append(counts, val)
			delete(hashMap, nums[i])
		}
	}
	// 进行回溯过程
	resultSubsetsWithDup = [][]int{}
	path := []int{}
	backtrackSubsetsWithDup(uniqueNums, counts, 0, path)
	return resultSubsetsWithDup
}
/*
uniqueNums 包含的 num 数值唯一
counts 每个数值出现的频率
通过 uniqueNums 和 counts 的组合找到可选列表

k 当前遍历到 counts 的第几个元素
path 最终路径
*/
func backtrackSubsetsWithDup(uniqueNums, counts []int, k int, path []int) {
	if k == len(uniqueNums) { // 终止条件
		temp := make([]int, len(path))
		copy(temp, path)
		resultSubsetsWithDup = append(resultSubsetsWithDup, temp)
		return
	}

	/*
		利用 count 将重复元素的所有组合列举出来, 如 [2,2,2], 双循环得到 [2], [2,2], [2,2,2]

		利用循环怎样 append 进去, 再利用循环以同样方式 remove 掉
	*/
	for count := 0; count <= counts[k]; count++ { 	// 做决策
		for i := 0; i < count; i++ {
			path = append(path, uniqueNums[k])
		}
		backtrackSubsetsWithDup(uniqueNums, counts, k+1, path)
		for i := 0; i < count; i++ {
			path = path[:len(path)-1]
		}
	}
}