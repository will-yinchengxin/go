package eight

/*
https://leetcode-cn.com/problems/eight-queens-lcci/

八皇后

设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展

示例:

	输入：4
	输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	解释: 4 皇后问题存在如下两个不同的解法。
	[
		[".Q..", // 解法 1
		 "...Q",
		 "Q...",
		 "..Q."],

		["..Q.", // 解法 2
		 "Q...",
		 "...Q",
		 ".Q.."]
	]
*/
var resultNQueens [][]string

func solveNQueens(n int) [][]string {
	var board [][]string // 初始化棋盘（空切片）
	for i := 0; i < n; i++ {
		var tmp []string
		for j:= 0; j < n; j++ {
			tmp = append(tmp, ".")
		}
		board = append(board, tmp)
	}
	backtrackSolveNQueens(0, board, n)
	return resultNQueens
}

// row 阶段(行)
// board 路径， 记录已经做出的决策
// 可选列表，通过 board 推倒出来 没有显示记录
func backtrackSolveNQueens(row int, board [][]string, n int) {
	// 结束条件，得到可行解
	if row == n {
		res := make([]string, 0, n)
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				str += board[i][j]
			}
			res = append(res, str)
		}
		resultNQueens = append(resultNQueens, res)
		return
	}
	for col := 0; col < n; col++ {
		if isOK(board, row, col, n) {
			board[row][col] = "Q"  //做选择，第row⾏的棋⼦放到了col列
			backtrackSolveNQueens(row+1, board, n) //考察下⼀⾏
			board[row][col] = "." //恢复选择
		}
	}
}
func isOK(board [][]string, row, col int, n int) bool {
	//检查列是否冲突
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	//检查右上对角线是否冲突
	i := row - 1
	j := col + 1
	for i >= 0 && j < n {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	//检查左上对角线是否冲突
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	return true
}
