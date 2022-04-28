package three

import "container/list"

/*
	// 用栈实现队列(两个栈来实现) (字节)
	// 方式一: 入队,直接入栈, 出队,倒腾两个栈(适合插入场景较多)
	// 方式二: 入队,倒腾两个栈, 出队, 直接出栈(适合弹出场景较多)
*/
type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

// -------------------------------------- 方式一 ----------------------------------------------
// 方式一: 入队,直接入栈, 出队,倒腾两个栈(适合插入场景较多)

// 入队,直接入栈
func (this *CQueue) StackToQueueOneEnqueue(value int) {
	this.stack1.PushBack(value)
}
// 出队,倒腾两个栈(适合插入场景较多)
func (this *CQueue) StackToQueueOneDequeue() int {
	if this.stack1 == nil {
		return -1
	}
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			// Back 返回最后一个元素
			// Remove 删除指定元素, 并返回该元素的值
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		val := this.stack2.Back()
		this.stack2.Remove(val)
		return val.Value.(int)
	}
	return -1
}
// -------------------------------------- 方式二 ----------------------------------------------
// 方式二: 入队,倒腾两个栈, 出队, 直接出栈(适合弹出场景较多)

// 入队 两个栈倒腾
func (this *CQueue) StackToQueueTwoEnqueue(value int) {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	this.stack1.PushBack(value)
	for this.stack2.Len() > 0 {
		this.stack1.PushBack(this.stack2.Remove(this.stack2.Back()))
	}
}
// 出队 直接出栈
func (this *CQueue) StackToQueueTwoDequeue() int {
	if this.stack1.Len() != 0 {
		val := this.stack1.Back()
		this.stack1.Remove(val)
		return val.Value.(int)
	}
	return -1
}