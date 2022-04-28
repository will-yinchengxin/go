package seven
/*
https://leetcode-cn.com/problems/list-of-depth-lcci/

特定深度节点链表

给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}
func listOfDepth(tree *TreeNode) []*ListNode {
	/*
		   按层遍历

		   使用 size 方式
				1) 记录每一层的size
				2) 遍历过程中 < size
				3) 同时拓展出新的size
	*/
	if tree == nil {
		return []*ListNode{new(ListNode)}
	}
	res := []*ListNode{}

	tmp := []*TreeNode{}
	tmp = append(tmp, tree)

	arrList := new(ListNode)
	tail := arrList
	size := 1

	for len(tmp) != 0 {
		currentSize := 0
		for i := 0; i < size; i++ {
			tail.Val = tmp[i].Val
			if i + 1 != size {
				tail.Next = &ListNode{}
			}
			tail = tail.Next
			if tmp[i].Left != nil {
				tmp = append(tmp, tmp[i].Left)
				currentSize++
			}
			if tmp[i].Right != nil {
				tmp = append(tmp, tmp[i].Right)
				currentSize++
			}
		}
		tmp = tmp[size:]
		res = append(res, arrList)
		arrList = new(ListNode)
		tail = arrList
		size = currentSize
	}
	return res
}