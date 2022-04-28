package six
/*
合并二叉树

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
否则不为NULL 的节点将直接作为新二叉树的节点。

示例1:
	输入:
		Tree 1                     Tree 2
			  1                         2
			 / \                       / \
			3   2                     1   3
		   /                           \   \
		  5                             4   7
	输出:
	合并后的树:
			 3
			/ \
		   4   5
		  / \   \
		 5   4   7

注意:合并必须从两个树的根节点开始。
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	/*
		1) 可以拆分为子问题进行求解
			两颗树的左子树的和作为新的左子树的val
			两颗树的右子树的和作为新的右子树的val
			其中一个为 nil 直接返回即可
	*/
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	newNode := new(TreeNode)
	newNode.Val = root1.Val + root2.Val
	newNode.Left = mergeTrees(root1.Left, root2.Left)
	newNode.Right = mergeTrees(root1.Right, root2.Right)

	return newNode
}