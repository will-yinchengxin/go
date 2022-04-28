package two

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
	输入：head = [1,2,2,1]
	输出：true

示例 2：
	输入：head = [1,2]
	输出：false

提示：
	链表中节点数目在范围[1, 105] 内
	0 <= Node.val <= 9
*/

// 遍历, 将数值插入至数组, 遍历数组, 判断是否回文
func isPalindrome(head *ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)
	for i, v := range nums[:n/2] {
		if v != nums[n-1-i] {
			return false
		}
	}
	return true
}

/*
	1) 寻找中间节点, 注意单双数的处理
	2) 反转后半段(前半段), 与后半段(前半段)进行对比
*/
func isPalindromeAno(head *ListNode) bool {
	if head == nil {
		return true
	}
	prevEven := middleNodeAno(head)
	revPrevEven := reverseListAno(prevEven)
	for revPrevEven != nil  {
		if head.Val != revPrevEven.Val {
			return false
		}
		head = head.Next
		revPrevEven = revPrevEven.Next
	}
	return true
}

func middleNodeAno(head *ListNode) (*ListNode) {
	// 单步数
	prevSingle := head
	// 双步数
	prevEven := head

	for prevEven.Next != nil && prevEven.Next.Next != nil {
		prevEven = prevEven.Next.Next
		prevSingle = prevSingle.Next
	}

	return prevSingle.Next
}

func reverseListAno(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
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