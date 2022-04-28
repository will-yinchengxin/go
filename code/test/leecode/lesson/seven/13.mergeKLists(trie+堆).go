package seven

import "container/heap"

/*
https://leetcode-cn.com/problems/merge-k-sorted-lists/

合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
	输入：lists = [[1,4,5],[1,3,4],[2,6]]
	输出：[1,1,2,3,4,4,5,6]
解释：
	链表数组如下：
		[
		  1->4->5,
		  1->3->4,
		  2->6
		]
	将它们合并到一个有序链表中得到。
		1->1->2->3->4->4->5->6

示例 2：
	输入：lists = []
	输出：[]

示例 3：
	输入：lists = [[]]
	输出：[]

提示：
	k == lists.length
	0 <= k <= 10^4
	0 <= lists[i].length <= 500
	-10^4 <= lists[i][j] <= 10^4
	lists[i] 按 升序 排列
	lists[i].length 的总和不超过 10^4
*/
type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
// 如果不用指针：minHeap 是 slice,append 会返回新的 pq，但是新的 pq 只是 push 方法的局部变量。所以 pq 并不会改变
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}
// pop 也是因为如果不用指针改变的只是局部变量
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*m = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	// go 中的优先级队列， 可以借助 heap 来实现， 标准库地址 src\container\heap, example_pq_test.go 中有完整实现

	if len(lists) == 0 || lists == nil { // 空判断
		return nil
	}
	lenList := len(lists)

	minQ := &minHeap{}
	heap.Init(minQ)

	for i := 0; i < lenList; i++ { // 将每个元素加入队列中
		if lists[i] != nil {
			heap.Push(minQ, lists[i])
		}
	}

	head := &ListNode{}
	tail := head
	for minQ.Len() > 0 {
		cur := heap.Pop(minQ).(*ListNode)
		tail.Next = cur
		tail = tail.Next
		// 如果链表不为空，将链表从新塞回优先级队列
		if cur.Next != nil {
			heap.Push(minQ, cur.Next)
		}
	}
	return head.Next
}