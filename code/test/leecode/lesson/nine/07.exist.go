package nine
/*
https://leetcode-cn.com/problems/word-search/

单词搜索

给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：
	输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	输出：true

示例 2：
	输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
	输出：true

示例 3：
	输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
	输出：false

提示：
	m == board.length
	n = board[i].length
	1 <= m, n <= 6
	1 <= word.length <= 15
	board 和 word 仅由大小写英文字母组成
*/
var existed bool
var hExist int
var wExist int
func exist(board [][]byte, word string) bool {
	/*
		改题目在 dfs 上做了升级, 有些回溯的感觉
	*/
	hExist = len(board)
	wExist = len(board[0])
	existed = false
	for i := 0; i < hExist; i++ {
		for j := 0; j < wExist; j++ {
			// 这里设计 回溯 问题, 所以全局的 visited 记录会有问题, 有点类似 trie 树
			visited := make([][]bool, hExist)
			for ii := 0; ii < hExist; ii++ {
				visited[ii] = make([]bool, wExist)
			}
			dfsExist(board, word, i, j, 0, visited)
			if existed {
				return existed
			}
		}
	}
	return existed
}
// k 表示当前遍历字符串的位置
func dfsExist(board [][]byte, word string, i, j, k int, visited [][]bool) {
	// 终止条件
	if existed || word[k] != board[i][j] {
		return
	}
	visited[i][j] = true
	if k == len(word) - 1 { // 找到目标字符串
		existed = true
		return
	}

	// 进行 dfs 探测
	dir := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for tag := 0; tag < 4; tag++ {
		newi := i + dir[tag][0]
		newj := j + dir[tag][1]
		if newi >= 0 && newi < hExist && newj >= 0 && newj < wExist && !visited[newi][newj] {
			dfsExist(board, word, newi, newj, k+1, visited)
		}
	}
	visited[i][j] = false // 这里形同回溯
}