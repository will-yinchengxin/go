package three
/*
栈(先进后出):
	操作受限的线性表 push(插入数据) pop(弹出栈顶元素) peak(查看栈顶的元素)
	基于 数组/链表 实现

	基于链表: 在头部插入数据(操作较方便)

队列(先进先出):
	基于 数组/链表 实现

	基于数组: 通常用来实现循环队列
		入队: ( tail = (tail + 1) % n)
		出队: ( head = (head + 1) % n )
		(tail + 1) % n = head // 浪费一个空间, 判断满了
*/

//---------------------------- 基于数组实现栈 ----------------------------
var arrayStack_stackArr []int
var arrayStack_count int // 栈中元素个数
var arrayStack_n int // 栈的大小
func arrayStack(n int) { // 创建一个大小为 n 的数组
	arrayStack_stackArr = make([]int, n)
	arrayStack_count = 0
	arrayStack_n = n
}
// 栈中插入元素
func arrayStackPush(item int) bool {
	if arrayStack_count == arrayStack_n {
		return false // 栈空间不足
	}
	arrayStack_stackArr[arrayStack_count] = item
	arrayStack_count++
	return true
}
// 栈中弹出元素
func arrayStackPop() int {
	if arrayStack_count == 0 {
		return - 1
	}
	val := arrayStack_stackArr[arrayStack_count-1]
	arrayStack_count--
	return val
}
// 栈返回指定元素
func arrayStackPeak() int {
	if arrayStack_count == 0 {
		return - 1
	}
	return arrayStack_stackArr[arrayStack_count - 1]
}

//---------------------------- 基于链表实现栈 ----------------------------
type ListNode struct {
	Val  int
	Next *ListNode
}

var data int
var next ListNode

func nodeStack(n int, node ListNode) { // 创建一个大小为 n 的数组
	data = n
	next = node
}
// 栈中插入元素 在头部插入数据(操作较方便)
var head *ListNode
func nodeStackPush(item int) {
	newNode := new(ListNode)
	newNode.Val = item
	newNode.Next = head
	head = newNode
}
// 栈中弹出元素 在头部插入数据(操作较方便)
func nodeStackPop() int {
	if head == nil {
		return -1
	}
	val := head.Val
	head = head.Next
	return val
}
// 栈返回指定元素
func nodeStackPeak() int {
	if head == nil {
		return -1
	}
	return head.Val
}

//---------------------------- 基于数组实现队列 ----------------------------
/*
数组实现循环队列, 两种标记和空的方法:
	1) 用 count 记录 存储在队列中的数据个数, count == 0 空, count == n 满
	2) 不适用count, head == tail 空, (tail+1) % n == head 满
*/
var arrayQueue []int
var nArrayQueue int
var headArrayQueue int
var tailArrayQueue int
func initArrayQueue(n int) {
	arrayQueue = make([]int, n)
	nArrayQueue = n
	headArrayQueue, tailArrayQueue = 0, 0
}
// 队列里插入元素 head == tail 空, (tail+1) % n == head 满
func enArrayQueue(item int) bool {
	if (tailArrayQueue + 1) % nArrayQueue == headArrayQueue {
		return false
	}
	arrayQueue[tailArrayQueue] = item
	tailArrayQueue = (tailArrayQueue + 1) % nArrayQueue // 在指定范围内循环的技巧
	return true
}
func deArrayQueue() int {
	if headArrayQueue == tailArrayQueue {
		return -1
	}
	val := arrayQueue[headArrayQueue]
	headArrayQueue = (headArrayQueue + 1) % nArrayQueue
	return val
}
//---------------------------- 基于链表实现队列 ----------------------------
var queueData int
var queueNext *ListNode

var listHead *ListNode
var listTail *ListNode

func initQueue(data int, node *ListNode) {
	queueData = data
	queueNext = node

	listHead = new(ListNode)
	listTail = new(ListNode)
}

func enListQueue(val int) {
	newNode := new(ListNode)
	newNode.Val = val
	if listTail.Next == nil {
		listTail = newNode
		listHead = newNode
	} else {
		listTail.Next = newNode
		listTail = listTail.Next
	}
}

func deListQueue() int {
	if listHead == nil {
		return -1
	}
	val := listHead.Val
	listHead = listHead.Next
	if listHead == nil {
		listTail = nil
	}
	return val
}