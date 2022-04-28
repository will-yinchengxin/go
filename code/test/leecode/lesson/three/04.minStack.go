package three

import "math"

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop()—— 删除栈顶的元素。
top()—— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
*/
type MinStack struct {
	queue1, queue2 []int
}


func ConstructorMinStack() MinStack {
	return MinStack{
		queue1: []int{},
		queue2: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(val int)  {
	if len(this.queue2) == 0 {
		this.queue2 = append(this.queue2, val)
	} else {
		if val < this.queue2[len(this.queue2) - 1] {
			this.queue2 = append(this.queue2, val)
		} else {
			this.queue2 = append(this.queue2, this.queue2[len(this.queue2) - 1])
		}
	}
	this.queue1 = append(this.queue1, val)
}


func (this *MinStack) Pop()  {
	this.queue1 = this.queue1[:len(this.queue1)-1]
	this.queue2 = this.queue2[:len(this.queue2)-1]
}


func (this *MinStack) Top() int {
	return this.queue1[len(this.queue2)-1]
}


func (this *MinStack) GetMin() int {
	return this.queue2[len(this.queue2) - 1]
}