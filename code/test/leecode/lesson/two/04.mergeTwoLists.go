package two
/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := new(ListNode)
	prev := result
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}
	if l1 == nil && l2 != nil {
		prev.Next = l2
	}
	if l2 == nil && l1 != nil {
		prev.Next = l1
	}
	return result.Next
}