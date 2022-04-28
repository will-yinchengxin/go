package one

/*
编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
target = 20
输出：false

输入：matrix = [
				[1, 4, 7,11,15],
				[2, 5, 8,12,19],
				[3, 6, 9,16,22],
				[10,13,14,17,24],
				[18,21,23,26,30]
			  ],
target = 5
输出：true
*/

func SearchMatrix(matrix [][]int, target int) bool {
	// 1) 暴力方式进行循环遍历, 查询结果

	// 2) 采用螺旋查找的方式
	// 		每一行依次递增
	// 		每一列依次递增
	lenMax := len(matrix)
	lenMaxSon := len(matrix[0])
	if lenMax == 0 || lenMaxSon == 0 {
		return false
	}

	// 下标记录
	i, j := 0, lenMaxSon-1
	for i <= lenMax - 1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		if matrix[i][j] < target {
			i++
			continue
		}
	}
	return false
}
