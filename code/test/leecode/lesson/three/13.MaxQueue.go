package three

import "container/list"

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value需要返回 -1

示例 1：
	输入:
	["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
	[[],[1],[2],[],[],[]]
	输出:[null,null,null,2,1,2]

示例 2：
	输入:
	["MaxQueue","pop_front","max_value"]
	[[],[],[]]
	输出:[null,-1,-1]
*/
type MaxQueue struct {
	list *list.List
	stack []int
}

func ConstructorMaxQueue() MaxQueue {
	return MaxQueue{
		list: list.New(),
	}
}

func (this *MaxQueue) Push_back(value int)  {
	// 先进行顺序队列数值的存储
	this.list.PushBack(value)
	// 单调栈存储 递减的值
	for len(this.stack) > 0 && this.stack[len(this.stack) - 1] < value {
		this.stack = this.stack[:len(this.stack) - 1]
	}
	this.stack = append(this.stack, value)
}

func (this *MaxQueue) Max_value() int {
	if this.list.Len() <= 0 {
		return -1
	}
	return this.stack[0]
}

func (this *MaxQueue) Pop_front() int {
	if this.list.Len() <= 0 {
		return -1
	}
	val := this.list.Front().Value.(int)
	this.list.Remove(this.list.Front())
	if len(this.stack) > 0 {
		if val == this.stack[0] {
			this.stack = this.stack[1:]
		}
	}
	return val
}
