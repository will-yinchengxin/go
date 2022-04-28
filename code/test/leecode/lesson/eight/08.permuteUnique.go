package eight

/*
https://leetcode-cn.com/problems/permutations-ii/

全排列 II

给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
	输入：nums = [1,1,2]
	输出：
	[[1,1,2],
	 [1,2,1],
	 [2,1,1]]

示例 2：
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：
	1 <= nums.length <= 8
	-10 <= nums[i] <= 10
*/
var resultPermuteUnique [][]int
func permuteUnique(nums []int) [][]int {
	/*
		先构建出 nums 和 counts 两个切片来记录数量
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
	uniqueNums := make([]int, 0 , lenHM)
	counts := make([]int, 0 , lenHM)
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if val, ok := hashMap[nums[i]]; ok {
			uniqueNums = append(uniqueNums, nums[i])
			counts = append(counts, val)
			delete(hashMap, nums[i])
		}
	}
	resultPermuteUnique = [][]int{}
	path := []int{}

	backtrackPermuteUnique(uniqueNums, counts, 0, path, numsLen)
	return resultPermuteUnique
}
func backtrackPermuteUnique(uniqueNums, counts []int, k int, path []int, numsLen int) {
	// 终止条件
	if k == numsLen {
		temp := make([]int, len(path))
		copy(temp, path)
		resultPermuteUnique = append(resultPermuteUnique, temp)
		return
	}

	// 决策阶段
	for i := 0; i < len(uniqueNums); i++ {
		if counts[i] == 0 {
			continue
		}
		path = append(path, uniqueNums[i]) // 添加选择
		counts[i]--
		backtrackPermuteUnique(uniqueNums, counts, k+1, path, numsLen)
		path = path[:len(path)-1] // 撤销选择
		counts[i]++
	}
}