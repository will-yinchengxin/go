package nine

import "sort"

/*
https://leetcode-cn.com/problems/pond-sizes-lcci/

水域大小

你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。
池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。

示例：
	输入：
		[
		  [0,2,1,0],
		  [0,1,0,1],
		  [1,1,0,1],
		  [0,1,0,1]
		]
	输出：
		[1,2,4]

提示：
	0 < len(land) <= 1000
	0 < len(land[i]) <= 1000
*/

/*
	这里可以不使用 visited 来记录已访问坐标,可以直接修改 land 数组中 0 为其他值
*/
var hPondSizes int
var wPondSizes int
var countLand int
var pondSizesArr []int
func pondSizes(land [][]int) []int {

	hPondSizes = len(land)
	wPondSizes = len(land[0])
	pondSizesArr = make([]int, 0)
	countLand = 0
	for i := 0; i < hPondSizes; i++ {
		for j := 0; j < wPondSizes; j++ {
			if land[i][j] == 0 {
				countLand = 0
				dfsPondSizes(land, i, j)
				pondSizesArr = append(pondSizesArr, countLand)
			}
		}
	}
	sort.Ints(pondSizesArr)
	return pondSizesArr
}
func dfsPondSizes(land [][]int, i, j int) {
	countLand++
	land[i][j] = 1 // !!!!!!
	dir := [][]int{{-1,0},{1,0},{0,1},{0,-1}, {-1,-1},{1,1},{-1,1},{1,-1}}
	for ii := 0; ii < 8; ii++ {
		newi := i + dir[ii][0]
		newj := j + dir[ii][1]
		if newi < 0 || newi >= hPondSizes || newj < 0 || newj >= wPondSizes || land[newi][newj] != 0 {
			continue
		}
		dfsPondSizes(land, newi, newj)
	}
}

//var visitedPondSizes [][]bool
//var hPondSizes int
//var wPondSizes int
//var countLand int
//var pondSizesArr []int
//func pondSizes(land [][]int) []int {
//	/*
//		这里可以不使用 visited 来记录已访问坐标,可以直接修改 land 数组中 0 为其他值
//	*/
//	hPondSizes = len(land)
//	wPondSizes = len(land[0])
//	visitedPondSizes = make([][]bool, hPondSizes)
//	for i := 0; i < hPondSizes; i++ {
//		visitedPondSizes[i] = make([]bool, wPondSizes)
//	}
//	pondSizesArr = make([]int, 0)
//	countLand = 0
//	for i := 0; i < hPondSizes; i++ {
//		for j := 0; j < wPondSizes; j++ {
//			if land[i][j] == 0 && !visitedPondSizes[i][j] {
//				countLand = 0
//				dfsPondSizes(land, i, j)
//				pondSizesArr = append(pondSizesArr, countLand)
//			}
//		}
//	}
//	sort.Ints(pondSizesArr)
//	return pondSizesArr
//}
//func dfsPondSizes(land [][]int, i, j int) {
//	visitedPondSizes[i][j] = true
//	countLand++
//	land[i][j] = 1
//	dir := [][]int{{-1,0},{1,0},{0,1},{0,-1}, {-1,-1},{1,1},{-1,1},{1,-1}}
//	for ii := 0; ii < 8; ii++ {
//		newi := i + dir[ii][0]
//		newj := j + dir[ii][1]
//		if newi < 0 || newi >= hPondSizes || newj < 0 || newj >= wPondSizes || visitedPondSizes[newi][newj] || land[newi][newj] != 0 {
//			continue
//		}
//		dfsPondSizes(land, newi, newj)
//	}
//}