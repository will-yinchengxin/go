package nine
/*
https://leetcode-cn.com/problems/number-of-islands/

岛屿数量

给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
	输入：
		grid = [
		  ["1","1","1","1","0"],
		  ["1","1","0","1","0"],
		  ["1","1","0","0","0"],
		  ["0","0","0","0","0"]
		]
	输出：1

示例 2：
	输入：
		grid = [
		  ["1","1","0","0","0"],
		  ["1","1","0","0","0"],
		  ["0","0","1","0","0"],
		  ["0","0","0","1","1"]
		]
	输出：3

提示：
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 300
	grid[i][j] 的值为 '0' 或 '1'
*/
var visitedIslands [][]bool
var h int
var w int
var countNumIslands int
func numIslands(grid [][]byte) int {
	h = len(grid)
	w = len(grid[0])
	visitedIslands = make([][]bool, h)
	for i := 0; i < h; i++ {
		visitedIslands[i] = make([]bool, w)
	}
	countNumIslands = 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '1' && !visitedIslands[i][j] {
				countNumIslands++
				dfsNumIslands(grid, i, j)
			}
		}
	}
	return countNumIslands
}
func dfsNumIslands(grid [][]byte, i, j int) {
	visitedIslands[i][j] = true
	dir := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
	for ii := 0; ii < 4; ii++ {
		newi := i + dir[ii][0]
		newj := j + dir[ii][1]
		if newi < 0 || newi >= h || newj < 0 || newj >= w || visitedIslands[newi][newj] || grid[newi][newj] == '0'{
			continue
		}
		dfsNumIslands(grid, newi, newj)
	}
}