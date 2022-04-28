package eight

/*
https://leetcode-cn.com/problems/permutations/submissions/

46. 全排列

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
	输入：nums = [0,1]
	输出：[[0,1],[1,0]]

示例 3：
	输入：nums = [1]
	输出：[[1]]


提示：
	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	nums 中的所有整数 互不相同。
*/
/*
	全排列 给定 n 个不重复的数，求这组数的所有排列组合
		该问题就是一个标准的穷举问题，为了不重复不遗漏，我们一般这样做
		先固定第一位是啥，再固定第二位是啥....实际上这就是多阶段的决策问题，可以通过决策树来解决问题

		决策树 和 递归树 长相一致，沿着一条路径一直向下走，直到无路可走后返回上一个岔路口，从新选择新的岔路口继续向下走

		这里实质就是 len(nums) 的阶层 1*2*3
*/
var result [][]int
var used []bool

func permute(nums []int) [][]int {
	result = [][]int{}
	path := make([]int, 0)
	used = make([]bool, len(nums))
	backtrack(nums, 0, path) // 可选列表，决策阶段，路径
	return result
}

// 可选列表 nums
// 决策阶段 k
// 路径记录再 path 中
func backtrack(nums []int, k int, path []int) {
	if k == len(nums) { // 结束条件
		temp := make([]int, len(path))
		/*
			为什么加入解集时，要将数组内容拷贝到一个新的数组里，再加入解集？
				因为该 path 变量存的是地址引用，结束当前递归时，将它加入 res 后，该算法还要进入别的递归分支继续搜索，
				还要继续将这个 path 传给别的递归调用，它所指向的内存空间还要继续被操作，所以 res 中的 path 的内容会被改变，这就不对。
				所以要弄一份当前的拷贝，放入 res，这样后续对 path 的操作，就不会影响已经放入 res 的内容。
		*/
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, k+1, path)
		path = path[:len(path)-1]
		used[i] = false
	}
	//for i := 0; i < len(nums); i++ {
	//	if used[i] {
	//		continue
	//	}
	//	path = append(path, nums[i])
	//	used[i] = true
	//	backtrack(nums, k+1, path)
	//	path = path[:len(path)-1]
	//	used[i] = false
	//}
}
