package three
/*
三合一。描述如何只用一个数组来实现三个栈。(将一个数组分为三个栈)

你应该实现 push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum) 方法。stackNum 表示栈下标，value 表示压入的值。

构造函数会传入一个stackSize参数，代表每个栈的大小。
*/
type TripleInOne struct {
	size   int
	lens   [3]int
	stacks []int
}


func ConstructorTripleInOne(stackSize int) TripleInOne {
	return TripleInOne{size: stackSize, stacks: make([]int, stackSize*3)}
}


func (this *TripleInOne) Push(stackNum int, value int)  {
	if this.lens[stackNum] == this.size {
		return
	}
	this.stacks[this.size*stackNum+this.lens[stackNum]] = value
	this.lens[stackNum]++
}


func (this *TripleInOne) Pop(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+this.lens[stackNum]-1]
	this.lens[stackNum]--
	return val
}


func (this *TripleInOne) Peek(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+this.lens[stackNum]-1]
	return val
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.lens[stackNum] == 0
}