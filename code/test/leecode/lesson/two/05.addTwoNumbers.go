package two
/*
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
	输入：l1 = [0], l2 = [0]
	输出：[0]

示例 3：
	输入：l1 = [9,9,9,9,9,9,9],
         l2 = [9,9,9,9]
	输出：     [8,9,9,9,0,0,0,1]
		      [8,9,9,0,0,0]
*/
func addTwoNumbersAno(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	prevA := l1
	prevB := l2
	result := new(ListNode)
	tail := result
	tag := 0
	for prevA != nil || prevB != nil {
		sum := 0
		if prevA != nil {
			sum += prevA.Val
			prevA = prevA.Next
		}
		if prevB != nil {
			sum += prevB.Val
			prevB = prevB.Next
		}
		if tag > 0 {
			sum += tag
			tag = 0
		}
		// 计算和的部分
		val := 0
		if sum >= 10 {
			val = sum - 10
			tag = 1
		} else {
			val = sum
		}
		tail.Next = new(ListNode)
		tail.Next.Val = val
		tail = tail.Next
	}
	if tag > 0 {
		tail.Next = new(ListNode)
		tail.Next.Val = tag
	}
	return result.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 处理 空 情况
	if l1 == nil && l2 == nil {
		return nil
	}
	result := new(ListNode)
	prev := result
	Base_System := 0
	for l1 != nil && l2 != nil {
		addNum := l1.Val + l2.Val + Base_System
		if addNum  >= 10 {
			Base_System = 1
			addNum = addNum - 10
			prev.Val = addNum
		} else {
			Base_System = 0
			prev.Val = addNum
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil {
			if Base_System > 0 {
				prev.Next = new(ListNode)
				prev.Next.Val = Base_System
			}
			return result
		}
		prev.Next = new(ListNode)
		prev = prev.Next
	}

	// 带着 Base_System 循环执行
	if l1 == nil && l2 != nil {
		for l2 != nil {
			addNum := l2.Val + Base_System
			if addNum  >= 10 {
				Base_System = 1
				addNum = addNum - 10
				prev.Val = addNum
			} else {
				Base_System = 0
				prev.Val = addNum
			}
			l2 = l2.Next
			if l2 != nil {
				prev.Next = new(ListNode)
				prev = prev.Next
			}
		}
	}
	if l2 == nil && l1 != nil {
		for l1 != nil {
			addNum := l1.Val + Base_System
			if addNum  >= 10 {
				Base_System = 1
				addNum = addNum - 10
				prev.Val = addNum
			} else {
				Base_System = 0
				prev.Val = addNum
			}
			l1 = l1.Next
			if l1 != nil {
				prev.Next = new(ListNode)
				prev = prev.Next
			}
		}
	}
	if Base_System > 0 {
		prev.Next = new(ListNode)
		prev.Next.Val = Base_System
	}
	return result
}