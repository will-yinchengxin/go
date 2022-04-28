package six
/*
翻转二叉树

翻转一棵二叉树。

示例：
	输入：
			 4
		   /   \
		  2     7
		 / \   / \
		1   3 6   9

	输出：
			 4
		   /   \
		  7     2
		 / \   / \
		9   6 3   1
*/
func invertTree(root *TreeNode) *TreeNode {
	/*
		1) 可以拆解为最小子问题
			递归交换每个节点的左右节点
	*/
	if root == nil {
		return nil
	}
	newNode := new(TreeNode)
	newNode.Val = root.Val
	newNode.Left = invertTree(root.Right)
	newNode.Right = invertTree(root.Left)
	return newNode
}