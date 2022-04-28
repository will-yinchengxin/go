package four
/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:
	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL

*/
// 递归解决问题
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// 遍历至尾节点
	newList := reverseList(head.Next)
	// 下面两个步骤的顺序不能反了
	head.Next.Next = head
	head.Next = nil

	return newList
}