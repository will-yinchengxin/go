package four
/*
对链表进行插入排序。

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：
插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

示例 1：
	输入: 4->2->1->3
	输出: 1->2->3->4

示例2：
	输入: -1->5->3->4->0
	输出: -1->0->3->4->5
*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	/*
		1) 将原链表的数据插入到结果链表中
		2) 怎样将原链表的数据插入到有序的结果链表中
			例: 6 插入到 r->2->5->7 中, 怎样寻找它插入的位置
			三要素:
			 - 指针的初始值
				q = head
			 - 遍历的结束条件
				p == nil && q.Next != nil
			 - 核心逻辑...
				q.next.val > p.val, 将 p 插入到 q 的后面, 否则 q = q.next
				if q.Next != nil && q.Next.Val > p.val {
					tmp := q.Next
					q.Next = p
					pNext = p.Next
					p.Next = tmp
				} else {
					q = q.Next
				}
		3) 特殊情况:
			将 1 插入到 r->2->5->7 中
				p.Next = r
				r = p
		tip: 采用头插法:
			// 2  插入 1->3->5 -- 头插
			p.Next = q.Next
			q.Next = p
	*/
	p := head // 原始链表
	res := new(ListNode) // 结果链表
	q := res
	for p != nil {
		// 提前记录下 head.Next
		tmp := p.Next
		for q.Next != nil && q.Next.Val <= p.Val {
			q = q.Next
		}
		// 2  插入 1->3->5 -- 头插
		p.Next = q.Next
		q.Next = p

		// 更新p节点
		p = tmp
		// 让 q 复原, 再次成为 res 的头节点
		q = res
	}
	return res.Next
}
