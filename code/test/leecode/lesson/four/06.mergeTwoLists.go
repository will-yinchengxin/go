package four
/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4

限制：
	0 <= 链表长度 <= 1000
*/
// 递归实现
// 时间复杂度：O(n)
// 空间复杂度：O(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	// 结果链表
	resul := new(ListNode)
	prev := resul

	// 进行递归
	prev = mergeTwoLists_r(l1, l2, prev)

	return resul.Next
}
func mergeTwoLists_r(l1 *ListNode, l2 *ListNode, prev *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		prev.Next = l2
	}
	if l2 == nil && l1 != nil {
		prev.Next = l1
	}
	if l2 != nil && l1 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
		mergeTwoLists_r(l1, l2, prev)
	}

	return prev
}