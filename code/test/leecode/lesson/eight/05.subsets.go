package eight

/*
https://leetcode-cn.com/problems/subsets/

子集

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
	输入：nums = [1,2,3]
	输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
	输入：nums = [0]
	输出：[[],[0]]

提示：
	1 <= nums.length <= 10
	-10 <= nums[i] <= 10
	nums 中的所有元素 互不相同
*/
var resultSubsets [][]int
func subsets(nums []int) [][]int {
	resultSubsets = [][]int{}
	path := []int{}
	backtrackSubsets(nums, 0, path)
	return resultSubsets
}
/*
nums[k] 可选列表, 选择 或者 不选
k 当前阶段
path 路径
*/
func backtrackSubsets(nums []int, k int, path []int) {
	if k == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		resultSubsets = append(resultSubsets, temp)
		return
	}
	backtrackSubsets(nums, k+1, path) // 不放置

	path = append(path, nums[k]) // 放置
	backtrackSubsets(nums, k+1, path)
	path = path[:len(path)-1]
}