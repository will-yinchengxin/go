package five
/*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:
	 输入：[1, 2, 3, 3, 2, 1]
	 输出：[1, 2, 3]

示例2:
	 输入：[1, 1, 1, 1, 2]
	 输出：[1, 2]

提示：
	链表长度在[0, 20000]范围内。
	链表元素在[0, 20000]范围内。

进阶：
	如果不得使用临时缓冲区，该怎么解决？
*/

// 使用 hash 解决
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hashMap := make(map[int]struct{})
	prev := head

	res := new(ListNode)
	prevRes := res
	for prev != nil {
		if _, ok := hashMap[prev.Val]; !ok {
			hashMap[prev.Val] = struct{}{}
			prevRes.Next = prev
			prevRes = prevRes.Next
		}
		prev = prev.Next
	}
	prevRes.Next = nil
	return res.Next
}