package six

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树:[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回：
	[3,9,20,15,7]
*/
func levelOrder(root *TreeNode) []int {
	/*
		1) 按层遍历本质就是广度优先搜索(前中后 序遍历本质是树上的深度优先搜索)
			- 可以采用数组
			- 可以采用队列(头部插入节点, 尾部导出节点)
		2) 这里使用 list 会超时, 我们使用切片来模拟 list
		3) helpList 我们从尾部插入, 从头部取出
	*/
	if root == nil {
		return nil
	}
	helpList := make([]*TreeNode, 0)
	helpList = append(helpList, root)
	res := []int{}
	for len(helpList) != 0 {
		elem := helpList[0]
		helpList = helpList[1:]
		res = append(res, elem.Val)
		if elem.Left != nil {
			helpList = append(helpList, elem.Left)
		}
		if elem.Right != nil {
			helpList = append(helpList, elem.Right)
		}
	}
	return res
}