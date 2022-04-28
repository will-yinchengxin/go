package six

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7

返回它的最大深度3
*/
func maxDepth(root *TreeNode) int {
	/*
		1) 使用递归解决
		2) max(左边的最大深度, 右边的最大深度)
		3) 时间复杂度O(n) 空间复杂度O(h)
	*/
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return getMax(left, right) + 1
}
func getMax(left ,right int) int {
	if left > right  {
		return left
	}
	return right
}