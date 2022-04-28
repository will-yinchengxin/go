package nine

/*
https://leetcode-cn.com/problems/route-between-nodes-lcci/

节点间通路

节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。

示例1:
	 输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
	 输出：true

示例2:
	 输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
	 输出 true

提示：
	节点数量n在[0, 1e5]范围内。
	节点编号大于等于 0 小于 n。
	图中可能存在自环和平行边。
*/
var visitedF []bool
var adj []map[int]bool
var found bool
// 使用 dfs 解决
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	visitedF = make([]bool, n)
	adj = make([]map[int]bool, n)
	found = false
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool, 0)
	}
	for i := 0; i < len(graph); i++ { // 构造成临接矩阵
		if _, ok := adj[graph[i][0]][graph[i][1]]; !ok {
		//if !adj[graph[i][0]][graph[i][1]] {
			adj[graph[i][0]][graph[i][1]] = true
		}
	}
	dfsF(start, target)
	return found
}
func dfsF(cur int, target int) {
	if found {return}
	if cur == target {
		found = true 
		return
	}
	visitedF[cur] = true
	for next, _ := range adj[cur] {
		if !visitedF[next] {
			dfsF(next, target)
		}
	}
}