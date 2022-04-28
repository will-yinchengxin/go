package four
/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
	输入：head = [1,3,2]
	输出：[2,3,1]

限制：
	0 <= 链表长度 <= 10000
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 非递归实现
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	arr := reverseArr(head)
	lenArr := len(arr)
	res := make([]int, lenArr)
	for i := 0; i < lenArr/2; i++ {
		res[i], res[lenArr - 1 - i] = res[lenArr - 1 - i], res[i]
	}
	// 0 1 2 3 4
	if lenArr % 2 > 0 {
		res[lenArr / 2] = arr[lenArr / 2]
	}
	return res
}
func reverseArr(head *ListNode) []int {
	arr := []int{}
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	return arr
}

// 递归实现
var result = []int{}
func reversePrintAno(head *ListNode) []int {
	if head == nil {
		return nil
	}
	reverseArrAno(head)
	return result
}
func reverseArrAno(head *ListNode) {
	// 终止条件
	if head == nil {
		return
	}
	reverseArrAno(head.Next)
	result = append(result, head.Val)
}