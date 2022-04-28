package two

// 1) 创建一个 新链表 来存储 result
// 2) 区分开 头指针 和 头节点
// 3) "=" 左边是指针, "=" 右边是节点, "=" 则用来描述指向


/*
解题三要素:
	1) 指针的初始值 p = head 或者 新增虚拟头节点 或者 其他...
	2) 遍历的结束条件, p == nil 还是 p.Next == nil
	3) 核心逻辑..

创建新的链表后, 插入方式有两种:
	1) 头插
		将 1 插入到 r->2->5->7 中
			p.Next = r
			r = p
	2) 尾插

	3) 反转列表
		1->2->3->4->5->NULL
		head.Next.Next = head
		head.Next = nil

统计链表的长度:
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 查找
func FindList(val int, head *ListNode) *ListNode {
	p := head
	for p != nil {
		if p.Val == val {
			return p
		}
		p = p.Next
	}
	return nil
}

// 链表头部插入
func InsertAtHead(val int, head *ListNode) *ListNode {
	newNode := ListNode{
		Val: val,
	}
	newNode.Next = head
	return &newNode
}

// 添加 虚拟头节点, 避免空的头节点
func AddFictitious(val int, head *ListNode) *ListNode {
	if head == nil {
		return  nil
	}
	// 添加虚拟头
	newHead := new(ListNode)
	newHead.Next = head
	prev := newHead
	for prev.Next != nil {
		if prev.Val != val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return newHead.Next
}

/*
	改变链表的万能写法
*/
func RemoveElements(head  *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	tail := newHead
	p := head
	for p != nil {
		tmp := p.Next
		if (p.Val != val) {
			p.Next = nil
			tail.Next = p
			tail = p
		}
		p = tmp
	}
	return newHead.Next
}
