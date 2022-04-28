package seven
/*
BiNode

二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，
转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。

返回转换后的单向链表的头节点。

注意：本题相对原题稍作改动

示例：
	输入： [4,2,5,1,3,null,6,0]
	输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]

提示：
	节点数量不会超过 100000。
*/
// 中序遍历转换为单链表
func convertBiNode(root *TreeNode) *TreeNode {
	dummyHead = new(TreeNode)
	tail = dummyHead
	convertBiNode_R(root)
	return dummyHead.Right
}
func convertBiNode_R(node *TreeNode) {
	if node == nil {
		return
	}
	left := node.Left
	right := node.Right

	convertBiNode_R(left)

	tail.Right = node
	tail = node
	tail.Left = nil

	convertBiNode_R(right)
}