package three

/*
https://leetcode-cn.com/problems/trapping-rain-water/

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
	解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
	输入：height = [4,2,0,3,2,5]
	输出：9
*/
// 单调栈方法:
func Trap(height []int) int {
	/*
		通过观察发现有规律可循,可以使用单调栈解决问题(单调栈指栈呈一种单调趋势存储, 要么递增, 要么递减)
		0) 创建临时栈, 只存储柱子高度的下标, 将第一个柱子预先存储, stack[] = 0, 遍历从下标 1 开始
		1) 栈内首个元素为根元素,也就是柱子的最左端, 插入临时栈中即可
		2) 后续插入的元素要与栈顶元素进行比较, 如果小于栈顶元素, 则继续插入
		3) 后续插入的元素要与栈顶元素进行比较, 如果等于栈顶元素则进行一个替换
		4) 后续元素大于栈顶元素, 说明有了凹槽, 可计算其'面积'进行累加至 sum 中
		5) 计算面积的公式为: 
			mid := stack[len(stack) - 1] 
			stack = stack[:len(stack) - 1] 
			高: min(height[i], height(stack[len(stack) - 1])) - height[mid]
			宽: i - stack[len(stack) - 1]) - 1 // 因为 i 从 1 开始计数, 只求中间宽度即可
			面积(sum): 宽 * 高 
	*/
	lenH := len(height)
	if lenH < 3 {
		return 0
	} 
	stack := []int{}
	sum := 0
	stack = append(stack, 0)
	for i := 1; i < lenH; i++ {
		if height[i] < height[stack[len(stack) - 1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack) - 1]] {
			stack = stack[:len(stack) - 1]
			stack = append(stack, i)
		} else {
			for len(stack) > 1 && height[i] > height[stack[len(stack) - 1]] {
				mid := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				h := judgeMax(height[i], height[stack[len(stack) - 1]]) - height[mid]
				w := i - stack[len(stack) - 1] - 1
				sum += h * w
			}
			// 保证最左端为可用端, 处理 0 1 0 2 特殊情况
			if height[i] > height[stack[len(stack) - 1]] {
				stack = stack[:len(stack) - 1]
			}
			stack = append(stack, i)
		}
	}
	return sum
}

func judgeMax(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// 前缀/后缀统计解法 (空间换时间, 但容易理解)
// 初始化一个下标 tag, 用来记录遍历过程中的最大值 maxHeight
// 准备两个数组, 长度为 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 从左边开始扫描, 一次记录每个位置左边的最大值, 第一个位置默认为 0, (max 值不包含自己)
//						 [0,   1,  0,  2,  1,  0,  1,  3,  2,  1,  2,  1]
// 例:             lmax: |_0_|_0_|_1_|_1_|_2_|_2_|_2_|_2_|_3_|_3_|_3_|_3_|
// 右边同理, 	   rmax: |_3_|_3_|_3_|_3_|_3_|_3_|_3_|_2_|_2_|_2_|_1_|_0_|
//  for i = 0; i++ ; i < len(height) (
//		S总: = min(lmax[i], rmax[i]) - h[i] // 可能为负数, 需要判断
//	) 为负数, 直接置零

// 优化, 为了需要每次判断是否为负数, 每次去 max 值包含自己
//						 [0,   1,  0,  2,  1,  0,  1,  3,  2,  1,  2,  1]
// 例:             lmax: |_0_|_1_|_1_|_2_|_2_|_2_|_2_|_3_|_3_|_3_|_3_|_3_|
// 右边同理, 	   rmax: |_3_|_3_|_3_|_3_|_3_|_3_|_3_|_3_|_2_|_2_|_2_|_1_|
// 这样就不需要再进行 负数的判断
