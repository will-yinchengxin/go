package seven

import (
	"container/heap"
	"math"
)

/*
https://leetcode-cn.com/problems/k-closest-points-to-origin/

最接近原点的 K 个点

我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
（这里，平面上两点之间的距离是欧几里德距离。）
你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

示例 1：
	输入：points = [[1,3],[-2,2]], K = 1
	输出：[[-2,2]]
	解释：
		(1, 3) 和原点之间的距离为 sqrt(10)，
		(-2, 2) 和原点之间的距离为 sqrt(8)，
		由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
		我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。

示例 2：
	输入：points = [[3,3],[5,-1],[-2,4]], K = 2
	输出：[[3,3],[-2,4]]
	（答案 [[-2,4],[3,3]] 也会被接受。）

提示：
	1 <= K <= points.length <= 10000
	-10000 < points[i][0] < 10000
	-10000 < points[i][1] < 10000
*/
type maxKCHeap [][]int
func (m maxKCHeap) Len() int {
	return len(m)
}
func (m maxKCHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxKCHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}
func (m *maxKCHeap) Pop() interface{} {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[0 : n-1]
	return item
}
func (m maxKCHeap) Less(i, j int) bool { // 确定大顶堆的关键
	return m[i][0] > m[j][0]
}
func (m maxKCHeap) Peek() interface{} {
	if m.Len() == 0 {
		return []int{math.MaxInt32, -1}
	}
	return m[0]
}
// ------------------------------------------------------------
func kClosest(points [][]int, k int) [][]int {
	/*
		解题思路:
			- 选择大顶堆合适
			- 先向大顶堆中添加 K 组数据
			- 后根据比较结果选择性向堆顶添加元素
			- 根据 K 的数目将结果传递至 result 中即可
	*/
	LenP := len(points)
	if LenP == 0 {
		return [][]int{}
	}
	maxKCH := &maxKCHeap{}
	heap.Init(maxKCH)
	for i := 0; i < k; i++ { // 将前 K 个元素计算结果插入 大顶堆中
		heap.Push(maxKCH, []int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i})
	}
	for i := k; i < LenP; i++ {
		tmp := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if tmp < maxKCH.Peek().([]int)[0] {
			heap.Pop(maxKCH)
			heap.Push(maxKCH, []int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i})
		}
	}
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = points[heap.Pop(maxKCH).([]int)[1]]
	}
	return ans
}
