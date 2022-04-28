package seven

import (
	"container/heap"
	"math"
	"sort"
)

/*
https://leetcode-cn.com/problems/top-k-frequent-elements/

前 K 个高频元素

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
	输入: nums = [1,1,1,2,2,3], k = 2
	输出: [1,2]

示例 2:
	输入: nums = [1], k = 1
	输出: [1]

提示：
	1 <= nums.length <= 105
	k 的取值范围是 [1, 数组中不相同的元素的个数]
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
*/
type QElement struct{
	val int
	count int
}
type minQEHeap []QElement
func (pq minQEHeap) Len() int {
	return len(pq)
}
func (pq minQEHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *minQEHeap) Push(x interface{}) {
	*pq = append(*pq, x.(QElement))
}
func (pq *minQEHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func (pq minQEHeap) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}
func (m minQEHeap) Peek() interface{} {
	if m.Len() == 0 {
		return QElement{math.MinInt32, -1}
	}
	return m[0]
}
// 优先级队列，末尾为最小值，首部为最大值
func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, 0)
	for _, num := range nums { // 通过 hashMap 存储每个元素出现的频率
		count := 0
		if c, ok := counts[num]; ok {
			count = c
		}
		counts[num] = count + 1
	}

	queue := &minQEHeap{}

	for num, count := range counts {
		if queue.Len() < k {
			heap.Push(queue, QElement{val:num, count: count})
		} else {
			if queue.Peek().(QElement).count < count { // 堆顶出现频率最高的元素 频率 < 当前元素出现 频率
				heap.Pop(queue) // 弹出末尾最小元素
				heap.Push(queue, QElement{val:num, count: count})
			}
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(queue).(QElement).val
	}
	sort.Ints(result)
	return result
}
