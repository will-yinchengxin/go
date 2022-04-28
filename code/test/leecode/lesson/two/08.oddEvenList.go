package two

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。
	你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
	输入: 1->2->3->4->5->NULL
	输出: 1->3->5->2->4->NULL

示例 2:
	输入: 2->1->3->5->6->4->7->NULL
	输出: 2->3->6->7->1->5->4->NULL

说明:
	应当保持奇数节点和偶数节点的相对顺序。
	链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/
// 1
// 1 2
// 1 2 3
// 1 2 3 4
// [1,2,3,4,5]   // [1,5]
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	/*
	1) 奇的下一个为偶数, 偶数的下一个奇数
	2) 奇偶拼接为结果集
	3) 边界: 偶数 及 偶数下一个 值均不为 nil
	*/
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func oddEvenListAno(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead := new(ListNode)
	oddTail := oddHead
	evenHead := new(ListNode)
	evenTail := evenHead

	p := head
	count := 1
	for p != nil {
		tmp := p.Next
		if count % 2 == 1 { // 奇数
			p.Next = nil
			oddTail.Next = p
			oddTail = oddTail.Next
		} else { // 偶数
			p.Next = nil
			evenTail.Next = p
			evenTail = evenTail.Next
		}
		count++
		p = tmp
	}
	oddTail.Next = evenHead.Next
	return oddHead.Next
}