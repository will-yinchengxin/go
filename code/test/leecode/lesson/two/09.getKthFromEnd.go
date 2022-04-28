package two
/*
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

示例：
	给定一个链表: 1->2->3->4->5, 和 k = 2.
	返回链表 4->5.
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	/*
		1) 双指针的方式, 定义一个头指针 head 和尾指针 tail, 之间的距离为 k
		2)连个指针通知遍历, 当头指针 head 的下一个值为nil, 返回尾指针 tail 即可
	*/
	tail := head
	newHead := head
	tag := 1
	for newHead != nil && newHead.Next != nil {
		if tag < k {
			tag++
		} else {
			tail = tail.Next
		}
		newHead = newHead.Next
	}
	return tail
}

func getKthFromEndAno(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	fast := head
	slow := head

	count := 0
	for fast != nil { // 遍历第一遍使链表到达正数 k 的位置
		count ++
		if count == k {
			break
		}
		fast = fast.Next
	}
	if fast == nil { // 链表长度不足 k
		return nil
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}