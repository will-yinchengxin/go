package nine
/*
https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

机器人的运动范围:

地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外）
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，
因为3+5+3+8=19。请问该机器人能够到达多少个格子？

示例 1：
	输入：m = 2, n = 3, k = 1
	输出：3

示例 2：
	输入：m = 3, n = 1, k = 0
	输出：1

提示：
	1 <= n,m <= 100
	0 <= k <= 20
*/
var visited [][]bool
var count int
func movingCount(m int, n int, k int) int {
	visited = make([][]bool, m)
	count = 0
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dfs(0, 0, m, n, k)
	return count
}
func dfs(i, j, m, n, k int) {
	visited[i][j] = true
	count++
	coordinate := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for ii := 0; ii < 4; ii++ { // 对山下左右的坐标进行探测 (2, 3, 1)
		newi := i + coordinate[ii][0]
		newj := j + coordinate[ii][1]
		if newi >= m || newi < 0 || newj >= n || newj < 0 || visited[newi][newj] || !check(newi, newj, k) {
			continue
		}
		dfs(newi, newj, m, n, k)
	}
}
func check(i, j, k int) bool {
	sum := 0
	for i !=0 {
		sum += (i%10)
		i /= 10
	}
	for j != 0 {
		sum += (j%10)
		j /= 10
	}
	return sum <= k
}