package eight
/*
https://leetcode-cn.com/problems/sudoku-solver/

解数独

编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用'.'表示。
*/
var rows [9]map[byte]bool
var cols [9]map[byte]bool
var blocks [3][3]map[byte]bool
var solved bool
func solveSudoku(board [][]byte)  {
	/*
		1）判定为回溯问题
		2）多阶段决策
		3）可选列表（1-9）
		4）带入回溯模板

		为了快速判别，减少回溯的次数，这里可以加以标记
			rows[][]bool，如：rows[0][5] = true 表示第 0 行的 5 已经存在 （位图方式）
			clos[][]bool，如：clos[0][5] = true 表示第 0 列的 5 已经存在 （位图方式）
			blocks[][][]bool， 如：blocks[0][0][5] = true 表示第一个九宫格的 5 已经存在

		这里也可以选用 map 来记录一下结果
	*/
	solved = false // 是否找到可行解

	rows = [9]map[byte]bool{}
	cols = [9]map[byte]bool{}
	blocks = [3][3]map[byte]bool{}
	// 初始化各类map
	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
		cols[i] = map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			blocks[i/3][j/3] = map[byte]bool{'1': false, '2': false, '3': false, '4': false, '5': false, '6': false, '7': false, '8': false, '9': false}
		}
	}


	for i := 0; i < 9; i++ { // 这里给定 9， 固定为 9*9 宫格
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j]
				rows[i][num] = true
				cols[j][num] = true
				blocks[i/3][j/3][num] = true
			}
		}
	}
	backtrackSolveSudoku(0, 0, board)
}
func backtrackSolveSudoku(row, col int, board [][]byte)  {
	if row == 9 { // 递归结束条件
		solved = true
		return
	}
	if board[row][col] != '.' { // 当前单位为数值
		nextRow := row
		nextCols := col + 1
		if col == 8 {
			nextRow = row + 1
			nextCols = 0
		}

		backtrackSolveSudoku(nextRow, nextCols, board)

		if solved { // 如果已经找到可行解，提前终止条件
			return
		}
	} else { // 这里为 . 需要判别填充字符
		for num := byte('1'); num <= byte('9'); num++ {
			if !rows[row][num] && !cols[col][num] && !blocks[row/3][col/3][num] { // 当 当前数字满足存在的条件时

				board[row][col] = num // 将字符插入熟读序列

				// 做选择，更改 节点状态
				rows[row][num] = true
				cols[col][num] = true
				blocks[row/3][col/3][num] = true

				nextRow := row
				nextCols := col + 1
				if col == 8 {
					nextRow = row + 1
					nextCols = 0
				}

				backtrackSolveSudoku(nextRow, nextCols, board) // 递归

				if solved { // 如果已经找到可行解，提前终止条件
					return
				}

				// 撤销选择，恢复 节点状态
				board[row][col] = '.'
				rows[row][num] = false
				cols[col][num] = false
				blocks[row/3][col/3][num] = false
			}
		}
	}
}