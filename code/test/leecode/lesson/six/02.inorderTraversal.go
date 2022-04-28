package six
// 二叉树中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var pore func(node *TreeNode)
	pore = func(node *TreeNode) {
		if node == nil {
			return
		}
		pore(node.Left)
		res = append(res, node.Val)
		pore(node.Right)
	}
	pore(root)
	return res
}