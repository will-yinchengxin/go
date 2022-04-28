package one

/*
给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。// 空间复杂度为 0

输入：[
		[1,2,3],
		[4,5,6],
		[7,8,9]
     ]
输出：[
		[7,4,1],
		[8,5,2],
		[9,6,3]
	 ]

输入：[
		[5,1,9,11],
		[2,4,8,10],
		[13,3,6,7],
		[15,14,12,16]
	]
输出：[
		[15,13,2,5],
		[14,3,4,1],
		[12,6,8,9],
		[16,7,10,11]
	]

输入：[[1]]
输出：[[1]]

输入：[
		[1,2],
		[3,4]
	]
输出：[
		[3,1],
		[4,2]
	]
*/

// 1) 借助辅助数组
func Rotate(matrix [][]int)  {
	// 借助辅助数组
	lenMatrix := len(matrix)
	if lenMatrix == 0 {
		return
	}
	newMatrix := make([][]int, lenMatrix)
	// 声明一个二维切片
	for i := range newMatrix {
		newMatrix[i] = make([]int, lenMatrix)
	}

	for i := 0; i < lenMatrix; i++ {
		for j := 0; j < lenMatrix; j++ {
			// 正旋转 和 逆旋转 关键
			newMatrix[j][lenMatrix-1-i] = matrix[i][j]
		}
	}
	for i := 0; i < lenMatrix; i++ {
		for j := 0; j < lenMatrix; j++ {
			matrix[i][j] = newMatrix[i][j]
		}
	}
}

// 2) 翻转代替旋转
func RotateOne(matrix [][]int)  {
	n := len(matrix)
	//先上下翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			// 原始的 行/列 号 i j
			// 目标的 行/列 号 n-i-1 j
			swap(matrix, i, j, n-i-1, j)
		}
	}
	//再对角翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// 解释 横纵坐标的交换
			// 原始的 行/列 号 i j
			// 目标的 行/列 号 j i
			swap(matrix, i, j, j, i)
		}
	}
}
func swap (matrix [][]int, i, j, p, q int) {
	matrix[i][j], matrix[p][q] = matrix[p][q], matrix[i][j]
}

// 3) 原地旋转
func RotateTwo(matrix [][]int)  {
	n := len(matrix)
	s1_i := 0
	s1_j := 0
	for n > 1 {
		s2_i := s1_i
		s2_j := s1_j + n-1
		s3_i := s1_i + n-1
		s3_j := s1_j + n-1
		s4_i := s1_i + n-1
		s4_j := s1_j
		for move := 0; move<=n-2; move++ {
			p1_i := s1_i + 0
			p1_j := s1_j + move
			p2_i := s2_i + move
			p2_j := s2_j + 0
			p3_i := s3_i + 0
			p3_j := s3_j - move
			p4_i := s4_i - move
			p4_j := s4_j + 0
			swap4(matrix, p1_i, p1_j, p2_i, p2_j, p3_i, p3_j, p4_i, p4_j)
		}
		s1_i++
		s1_j++
		n-=2
	}

}
func swap4(a [][]int, i1, j1, i2, j2, i3, j3, i4, j4 int) {
	tmp := a[i1][j1]
	a[i1][j1] = a[i4][j4]
	a[i4][j4] = a[i3][j3]
	a[i3][j3] = a[i2][j2]
	a[i2][j2] = tmp
}
/*
变形 n*m 顺时针旋转 90°
*/





















