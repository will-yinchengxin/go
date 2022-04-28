package one

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

输入：[
		[1,2,3,4],
		[5,6,7,8],
		[9,10,11,12]
	]
输出：[
		1,2,3,4,
		8,12,11,10,
		9,5,6,7
	]
*/

func spiralOrder(matrix [][]int) []int {
	// 从 外维圈 到 内维圈 一圈一圈的遍历

	// 初始化参数
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, 0)

	// 记录每一圈四个边角的位置信息
	left := 0
	right := n - 1
	top := 0
	bottom := m - 1

	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		// top + 1 避免同一位置重复元素
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if top != bottom {
			for j := right - 1; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
		}
		if left != right {
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
