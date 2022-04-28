package two
/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 2：
	输入：head = [], val = 1
	输出：[]

示例 3：
	输入：head = [7,7,7,7], val = 7
	输出：[]
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 新增虚拟头节点
	newHead := new(ListNode)
	newHead.Next = head
	prev := newHead
	for prev.Next != nil {
		if prev.Next.Val == val{
			prev.Next = newHead.Next.Next // 左边为赋值, 右边是对象
		} else {
			prev = prev.Next
		}
	}
	return newHead.Next
}
