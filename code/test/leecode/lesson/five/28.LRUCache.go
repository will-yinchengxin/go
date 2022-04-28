package five
/*
https://leetcode-cn.com/problems/lru-cache/

实现 LRU 算法
*/
/*
缓存包含的操作
	1) 再缓存中查找一个数据 get
	2) 再缓存中删除一个数据 remove
	3) 再缓存中添加一个数据 put

基于 hash + 双向有序链表的实现方案
	1) 借助 hash, 快速得到 要查找 要删除 的节点
	2) 借助双向链表, 维护数据的有序性(按照时间访问)

双向有序链表(放置方式):
	1) 头放置最新数据, 尾放置最老数据
	2) 头放置最老数据, 尾放置最新数据

容量限制:
	1) 超出链表的容量, 则剔除最末尾元素
*/
var hashMap map[int]LRUCache

type LRUCache struct {
	capacity int
	size int
	key int
	val int
	head *LRUCache
	tail *LRUCache
	prev *LRUCache
	next *LRUCache
}

func constructorAno(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size: 0,
		head: &LRUCache{
			key: -1,
			val: -1,
			prev: nil,
			next: nil,
		},
		tail: &LRUCache{
			key: -1,
			val: -1,
			prev: nil,
			next: nil,
		},
	}
}

/*
- 先删除
- 在头部添加
- 返回结果值
*/
func (this *LRUCache) Get(key int) int {
	if this.size == 0 {
		return -1
	}
	node, ok := hashMap[key]
	if !ok {
		return -1
	}
	this.removeNode(node)
	this.addNodeToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int)  {
	node, ok := hashMap[key]
	if ok {
		node.val = value
		this.removeNode(node)
		this.addNodeToHead(node)
		return
	}
	if this.size == this.capacity {
		delete(hashMap, this.tail.prev.key)
		this.removeNode(*this.tail.prev)
		this.size--
	}
	newNode := LRUCache{
		key: key,
		val: value,
	}
	this.addNodeToHead(newNode)
	hashMap[key] = newNode
	this.size++
}

// 新增方法
func (this *LRUCache) remove(key int) {
	if node, ok := hashMap[key]; ok {
		this.removeNode(node)
		delete(hashMap, key)
		this.size--
	}
}
func (this *LRUCache) removeNode(node LRUCache) {
	node.next.prev = node.prev
	node.next.next = node.next
}
func (this *LRUCache) addNodeToHead(node LRUCache) {
	node.next = this.head.next
	this.head.next.prev = &node
	this.head.next = &node
	node.prev = this.head
}