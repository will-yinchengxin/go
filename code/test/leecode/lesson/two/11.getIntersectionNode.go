package two
/*
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
+

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 利用hashMap, 空间复杂度为 O(len(headA))
	// 这里hashMap 存储的一定是地址, 不能为 val 值

	// 双指针可以将空间复杂度降低值 O(1)
	if headA == nil || headB == nil {
		return nil
	}


	helpMap := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		helpMap[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if helpMap[tmp] {
			return tmp
		}
	}
	return nil
}

func getIntersectionNodeAno(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 求出 A, B 两个链表的长度
	lenA := 0
	prevA := headA
	for prevA != nil {
		lenA++
		prevA = prevA.Next
	}
	lenB := 0
	prevB := headB
	for prevB != nil {
		lenB++
		prevB = prevB.Next
	}
	// 让较长的链表走 long - short 步
	prevA = headA
	prevB = headB
	if lenA >= lenB {
		for i := 0; i < lenA - lenB; i++ {
			prevA = prevA.Next
		}
	} else {
		for i := 0; i < lenB - lenA; i++ {
			prevB = prevB.Next
		}
	}
	// 两个链表同时遍历
	for prevA != nil && prevB != nil && prevA != prevB {
		prevA = prevA.Next
		prevB = prevB.Next
	}
	if prevA == nil || prevB == nil {
		return nil
	}
	return prevA
}