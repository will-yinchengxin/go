package seven

import (
	"container/heap"
	"math"
)

/*
https://leetcode-cn.com/problems/find-median-from-data-stream/

数据流的中位数

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，
	[2,3,4] 的中位数是 3
	[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：
	void addNum(int num) - 从数据流中添加一个整数到数据结构中。
	double findMedian() - 返回目前所有元素的中位数。

示例：
	addNum(1)
	addNum(2)
	findMedian() -> 1.5
	addNum(3)
	findMedian() -> 2

进阶:
	如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法
	如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法
*/

/*
  解题思路：
	构建两个堆，一个 大顶堆 和一个 小顶堆
		大顶堆存放前半段数据
		小顶堆存放后半段数据
	两个堆顶的元素的组合（如果两个堆元素数量相等，就是两个堆顶 数值相加/2, 如果不等，就是元素较多的那个的堆顶元素）
*/
type priorityQueue []int
func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
// 大顶堆
type maxXHeap struct {
	priorityQueue
}
func (pq maxXHeap) Less(i, j int) bool {
	return pq.priorityQueue[i] > pq.priorityQueue[j]
}
func (pq *maxXHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return math.MaxInt32
	}
	return pq.priorityQueue[0]
}
// 小顶堆
type minXHeap struct {
	priorityQueue
}
func (pq minXHeap) Less(i, j int) bool {
	return pq.priorityQueue[i] < pq.priorityQueue[j]
}
func (pq *minXHeap) Peek() interface{} {
	if pq.Len() == 0 {
		return math.MinInt32
	}
	return pq.priorityQueue[0]
}
//----------------------------------------------------------------------------------------------------------
type MedianFinder struct {
	minQueue *minXHeap
	maxQueue *maxXHeap
}
func Constructor() MedianFinder {
	minQueue := &minXHeap{}
	maxQueue := &maxXHeap{}
	heap.Init(minQueue)
	heap.Init(maxQueue)
	return MedianFinder{
		minQueue: minQueue,
		maxQueue: maxQueue,
	}
}
func (this *MedianFinder) AddNum(num int) { // 加入元素
	if this.maxQueue.Len() == 0 || this.maxQueue.Peek().(int) >= num { 	// 向大顶堆中先加入元素
		heap.Push(this.maxQueue, num)
	} else { // 否则小顶堆中加入元素
		heap.Push(this.minQueue, num)
	}

	// 检测 大顶堆 和 小顶堆 中元素的数量
	for this.maxQueue.Len() < this.minQueue.Len() {
		elem := heap.Pop(this.minQueue)
		heap.Push(this.maxQueue, elem)
	}
	for this.minQueue.Len() < this.maxQueue.Len()-1 {
		elem := heap.Pop(this.maxQueue)
		heap.Push(this.minQueue, elem)
	}
}
func (this *MedianFinder) FindMedian() float64 { // 查找中位数
	if this.maxQueue.Len() > this.minQueue.Len() {
		return float64(this.maxQueue.Peek().(int))
	} else {
		return (float64(this.maxQueue.Peek().(int)) + float64(this.minQueue.Peek().(int)))/2
	}
}
