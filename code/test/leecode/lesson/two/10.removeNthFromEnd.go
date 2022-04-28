package two
/*
给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。

示例 1：
	输入：head = [1,2,3,4,5], n = 2
	输出：[1,2,3,5]
示例 2：
	输入：head = [1], n = 1
	输出：[]
示例 3：
	输入：head = [1,2], n = 1
	输出：[1]
*/

// 特殊情况不能处理 list 长度 < n + 1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	/*
		1) 遍历一次链表, 记录长度, 从新遍历计数
		2) 双指针方式, newHead 预先遍历 n 个节点后, newHead 再与 tail(给一个虚拟头节点, 解决剔除第 n 的元素的问题) 同时遍历
		3) 这里注意, 要给 tail 添加一个虚拟头节点, 方便处理 剔除的间隔 问题
	*/
	newHead := head
	result := new(ListNode)
	result.Next = head
	tail := result
	tag := 0
	for newHead != nil {
		if tag < n {
			tag++
		} else {
			tail = tail.Next
		}
		newHead = newHead.Next
	}

	tail.Next = tail.Next.Next

	return result.Next
}


func removeNthFromEndAno(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return nil
	}
	fast := head
	slow := head

	count := 0
	for fast != nil { // 遍历第一遍使链表到达正数 k 的位置
		count ++
		if count == n {
			break
		}
		fast = fast.Next
	}
	if fast == nil { // 链表长度不足 k
		return head
	}

	// 添加查找的 pre
	pre := new(ListNode)
	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	if pre.Next == nil {
		head = head.Next
	} else {
		pre.Next = slow.Next
	}
	return head
}