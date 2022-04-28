package six
/*
对称二叉树

给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
		 4
	   /   \
	  2     2
	 / \   / \
	3   4 4   3
	输入：root = [1,2,2,3,4,4,3]
	输出：true

示例 2：
		 1
	   /   \
	  2     2
	   \     \
	    3     3
	输入：root = [1,2,2,null,3,null,3]
	输出：false

提示：
	树中节点数目在范围 [1, 1000] 内
	-100 <= Node.val <= 100
*/
func isSymmetric(root *TreeNode) bool {
	/*
		1) 不能分解为最小问题
		2) 以镜像树来思路考虑问题,
				 4
			   /   \
			  2     2
			 / \   / \
			1   3 3   1
		- 最左边 VS 最右边
		- 左二  VS 右二

	*/
	if root == nil {
		return true
	}
	return isSymmetric_R(root.Left, root.Right)
}
func isSymmetric_R(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil && right == nil) || (right != nil && left == nil) {
		return false
	}
	if (left != nil && right != nil) && (left.Val == right.Val) {
		return isSymmetric_R(left.Left, right.Right) && isSymmetric_R(left.Right, right.Left)
	}
	return false // (left != nil && right != nil) && (left.Val != right.Val)
}