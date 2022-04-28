package two
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

输入：head = [1,2]
输出：[2,1]

输入：head = []
输出：[]
*/
func reverseList(head *ListNode) *ListNode {
	// 新的链表记录结果
	var result *ListNode
	prev := head
	/*
		1) 记录当前节点的next
		2) 将当前 对象 集成至倒叙链表
		3) 倒叙列表重新赋值(加入了新元素)
		4) 游标重新赋值 tmp, 继续遍历顺序链表
	*/
	for prev != nil {
		tmp := prev.Next
		prev.Next = result
		result = prev
		prev = tmp
	}
	return result
}