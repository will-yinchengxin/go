package nine
/*
https://leetcode-cn.com/problems/color-fill-lcci/

颜色填充

编写函数，实现许多图片编辑软件都支持的「颜色填充」功能。
待填充的图像用二维数组 image 表示，元素为初始颜色值。初始坐标点的行坐标为 sr 列坐标为 sc。需要填充的新颜色为 newColor 。
「周围区域」是指颜色相同且在上、下、左、右四个方向上存在相连情况的若干元素。
请用新颜色填充初始坐标点的周围区域，并返回填充后的图像。

示例：
	输入：
		image = [[1,1,1],[1,1,0],[1,0,1]]
		sr = 1, sc = 1, newColor = 2
		输出：[[2,2,2],[2,2,0],[2,0,1]]
	解释:
		初始坐标点位于图像的正中间，坐标 (sr,sc)=(1,1) 。
		初始坐标点周围区域上所有符合条件的像素点的颜色都被更改成 2 。
		注意，右下角的像素没有更改为 2 ，因为它不属于初始坐标点的周围区域。

提示：
	image 和image[0]的长度均在范围[1, 50] 内。
	初始坐标点 (sr,sc) 满足0 <= sr < image.length 和0 <= sc < image[0].length 。
	image[i][j] 和newColor表示的颜色值在范围[0, 65535] 内。
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	m := len(image[0])
	dfsFloodFill(image, n, m, sr, sc, image[sr][sc], newColor)
	return image
}
func dfsFloodFill(image[][]int, n, m, sr, sc, color, newColor int) {
	image[sr][sc] = newColor
	dirs := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
	for i := 0; i < 4; i++ {
		newr := sr + dirs[i][0]
		newc := sc + dirs[i][1]
		if newr < 0 || newr >= n || newc < 0 || newc >= m || image[newr][newc] != color || image[newr][newc] == newColor {
			continue
		}
		dfsFloodFill(image, n, m, newr, newc, color, newColor)
	}
}