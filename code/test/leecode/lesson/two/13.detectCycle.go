package two
/*
给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

示例 1：
	输入：head = [3,2,0,-4], pos = 1
	输出：返回索引为 1 的链表节点
	解释：链表中有一个环，其尾部连接到第二个节点。

示例2：
	输入：head = [1,2], pos = 0
	输出：返回索引为 0 的链表节点
	解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
	输入：head = [1], pos = -1
	输出：返回 null
	解释：链表中没有环。
*/

/*
借助hashMap

空间复杂度 O(n)
时间复杂度 O(n)
*/
func detectCycle(head *ListNode) *ListNode {
	// 使用 hashMap 来计数
	helpMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := helpMap[head]; ok {
			return head
		}
		helpMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

/*
采用纯链表解决

空间复杂度 O(1)
时间复杂度 O(n)

a + (n + 1)b + nc = 2(a + b)
a = c + (n-1)(b+c)

也是就是两个焦点相交后 与 新节点 a 相交的位置为入口节点
*/
func detectCycleAno(head *ListNode) *ListNode {
	/*
	采用纯链表解决

	空间复杂度 O(1)
	时间复杂度 O(n)

	a + (n + 1)b + nc = 2(a + b)
	a = c + (n-1)(b+c)

	也是就是两个焦点相交后 与 新节点 a 相交的位置为入口节点
	*/
	if  head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}		
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
