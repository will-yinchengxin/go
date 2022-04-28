package six
/*
二叉搜索树的第k大节点

给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

示例 1:
	输入: root = [3,1,4,null,2], k = 1
	   3
	  / \
	 1   4
	  \
	   2
	输出: 4

示例 2:
	输入: root = [5,3,6,2,4,null,null,1], k = 3
		   5
		  / \
		 3   6
		/ \
	   2   4
	  /
	 1
	输出: 4

限制：
	1 ≤ k ≤ 二叉搜索树元素个数
*/
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	// 进行 右 -> 根 -> 左
	var res []int
	var pore func(node *TreeNode)
	pore = func(node *TreeNode) {
		if node == nil {
			return
		}
		if len(res) == k {
			return
		}
		pore(node.Right)
		res = append(res, node.Val)
		pore(node.Left)
	}
	pore(root)
	return res[k-1]
}