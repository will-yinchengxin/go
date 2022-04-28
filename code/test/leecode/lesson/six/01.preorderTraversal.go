package six
/*
给你二叉树的根节点 root ，返回它节点值的前序遍历。

示例 1：
	输入：root = [1,null,2,3]
	输出：[1,2,3]

示例 2：
	输入：root = []
	输出：[]

示例 3：
	输入：root = [1]
	输出：[1]

示例 4：
	输入：root = [1,2]
	输出：[1,2]

示例 5：
	输入：root = [1,null,2]
	输出：[1,2]
*/
// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var preorderTraversal_X func(*TreeNode)
	preorderTraversal_X = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorderTraversal_X(node.Left)
		preorderTraversal_X(node.Right)
	}
	preorderTraversal_X(root)
	return
}
//-------------------------------------------------
//var res []int
//func preorderTraversal(root *TreeNode) []int {
//	res = []int{}
//	preorderTraversal_X(root)
//	return res
//}
//func preorderTraversal_X(node *TreeNode) {
//	if node == nil {
//		return
//	}
//	res = append(res, node.Val)
//	preorderTraversal_X(node.Left)
//	preorderTraversal_X(node.Right)
//}