package three
/*
栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。
该栈支持如下操作：push、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。
*/
type SortedStack struct {
	stack1, stack2 []int
}

func ConstructorSortedStack() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int)  {
	for len(this.stack1) > 0 && val > this.Peek() {
		tmp := this.stack1[len(this.stack1) - 1]
		this.stack2 = append(this.stack2, tmp)
		this.Pop()
	}
	this.stack1 = append(this.stack1, val)
	for len(this.stack2) > 0 {
		tmp := this.stack2[len(this.stack2) - 1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		this.stack1 = append(this.stack1, tmp)
	}
}

func (this *SortedStack) Pop() {
	if len(this.stack1) > 0 {
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
}


func (this *SortedStack) Peek() int {
	if len(this.stack1) > 0 {
		return this.stack1[len(this.stack1) - 1]
	}
	return 0
}


func (this *SortedStack) IsEmpty() bool {
	return len(this.stack1) == 0
}