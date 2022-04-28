package four

/*
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。

进阶：
	你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？ 非递归实现 快排/归并, 时间复杂度 O(nlogn), 空间复杂度 O(1)

示例 1：
	输入：head = [4,2,1,3]
	输出：[1,2,3,4]

示例 2：
	输入：head = [-1,5,3,4,0]
	输出：[-1,0,3,4,5]

示例 3：
	输入：head = []
	输出：[]

提示：
	链表中节点的数目在范围 [0, 5 * 10^4] 内
	-10^5 <= Node.val <= 10^5
*/
// 不符合要求, 空间复杂度为 O(n), 时间复杂度因为 O(n)
//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	p := head
//	arr := []int{}
//	for p != nil {
//		arr = append(arr, p.Val)
//		p = p.Next
//	}
//	p = head
//	sort.Ints(arr)
//	for i, lenA := 0, len(arr); i < lenA; i++ {
//		p.Val = arr[i]
//		p = p.Next
//	}
//	return head
//}

// 不使用 api 方式
func sortList(head *ListNode) *ListNode {
	// 采用归并方式, 设置一个 step 长度
	// 进行一个简化, 不进行分, 直接以 step 为一组进行排序, 例如以 2 为步长
	// 将两个排序的结果, 尾插至结果链表中:
	//  	例: head -> 2 -> 4 -> 6 -> 3 -> 7 -> 8,
	// 		排序: 2->4, 3->6, 7->8,
	//		结果链表 2->4->3->6->7->8
	// 然后进行两个有序链表的合并,
	if head == nil {
		return head
	}
	// 统计长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 设置步长
	step := 1
	for step < length {
		dummyHead := &ListNode{} // 结果链表
		tail := dummyHead

		p := head
		for p != nil { // 从 head 开始以 step 为分组个数进行排序, 并插入结果链表尾部
			// 第一个分段
			q := p // 用来记录每次分段的 尾节点
			count := 1
			for q != nil && count < step {
				q = q.Next
				count++
			}

			// 判断特殊情况, 没有节点或者只有一个节点
			if q == nil || q.Next == nil {
				tail.Next = p
				break
			}

			// 第二个分段
			r := q.Next
			count = 1
			for r != nil && count < step {
				r = r.Next
				count++
			}

			// 保存下一次起点
			var tmp *ListNode
			if r != nil {
				tmp = r.Next
			}

			// 进行两段的合并排序
			twoSortPartHead, twoSortParttail := mergeNodeList(p, q, r)
			tail.Next = twoSortPartHead
			tail = twoSortParttail
			// 重置 p 排序下面两段
			p = tmp
		}
		// 重置原链表
		head = dummyHead.Next
		// 步数扩容
		step *= 2
	}

	return head
}
// 合并两个有序链表
// 返回 头节点 和 尾节点
func mergeNodeList(a, b, end *ListNode) (*ListNode, *ListNode) {
	// 重组出两个待合并链表
	l1 := a
	l2 := b.Next
	b.Next = nil
	if end != nil {
		end.Next = nil
	}

	// 结果链表
	res := new(ListNode)
	prev := res

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
		prev = end // 记录尾节点
	}
	if l2 == nil && l1 != nil {
		prev.Next = l1
		prev = b // 记录尾节点
	}
	return res.Next, prev
}