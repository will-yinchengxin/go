package five
/*
https://leetcode-cn.com/problems/linked-list-cycle/

使用 hash 解决
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	hashMap := make(map[*ListNode]struct{})
	prev := head
	for prev != nil {
		if _, ok := hashMap[prev]; ok {
			return true
		}
		hashMap[prev] = struct{}{}
		prev = prev.Next
	}
	return false
}