package two

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。

输入：head = [1,1,2]
输出：[1,2]

输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newNode := new(ListNode)
	newNode.Next = head
	prev := newNode.Next
	for prev.Next != nil {
		if prev.Val != prev.Next.Val {
			prev = prev.Next
		} else {
			prev.Next = prev.Next.Next
		}
	}
	return newNode.Next
}