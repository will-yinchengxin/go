package three

import (
	"container/list"
)

/*
https://leetcode-cn.com/problems/animal-shelter-lcci/

动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”
（由其进入收容所的时间长短而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。
请创建适用于这个系统的数据结构，实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。

enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。

dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。

示例1:
	 输入：
	["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
	[[], [[0, 0]], [[1, 0]], [], [], []]
	 输出：
	[null,null,null,[0,0],[-1,-1],[1,0]]
*/
type AnimalShelf struct {
	dogList *list.List
	catList *list.List
	timeline *list.List
}


func ConstructorAnimalShelf() AnimalShelf {
	return AnimalShelf{
		list.New(),
		list.New(),
		list.New(),
	}
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	var l *list.List
	if animal[1] == 0 {
		l = this.catList
	} else {
		l = this.dogList
	}
	ele := this.timeline.PushBack(animal)
	l.PushBack(ele)
}


func (this *AnimalShelf) DequeueAny() []int {
	if this.timeline.Len() == 0 {
		return []int{-1, -1}
	}
	animal := this.timeline.Remove(this.timeline.Front()).([]int)

	var l *list.List
	if animal[1] == 0 {
		l = this.catList
	} else {
		l = this.dogList
	}
	l.Remove(l.Front()) // element.Value should equals timeline.Front()

	return animal
}


func (this *AnimalShelf) DequeueDog() []int {
	if this.dogList.Len() == 0 {
		return []int{-1, -1}
	}
	ele := this.dogList.Remove(this.dogList.Front()).(*list.Element)
	return this.timeline.Remove(ele).([]int)
}


func (this *AnimalShelf) DequeueCat() []int {
	if this.catList.Len() == 0 {
		return []int{-1, -1}
	}
	ele := this.catList.Remove(this.catList.Front()).(*list.Element)
	return this.timeline.Remove(ele).([]int)
}

