package six
// 后续遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var pore func(node *TreeNode)
	pore = func(node *TreeNode) {
		if node == nil {
			return
		}
		pore(node.Left)
		pore(node.Right)
		res = append(res, node.Val)
	}
	pore(root)
	return
}