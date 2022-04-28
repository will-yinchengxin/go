package six

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:
	给定二叉树 [3,9,20,null,null,15,7]

		3
	   / \
	  9  20
		/  \
	   15   7
  返回 true 。

示例 2:
	给定二叉树 [1,2,2,3,3,null,null,4,4]

		   1
		  / \
		 2   2
		/ \
	   3   3
	  / \
	 4   4
   返回false 。
*/
var tag = true
func isBalanced(root *TreeNode) bool {
	/*
		1) 递归实现
		2) 不能通过 左子树是 右子树是 整个树都是 这样来判断
		3) 可以通过变形, 求左右子树的最大深度, 差值如果 > 1 就不是平衡树
	*/
	if root == nil {
		return true
	}
	getDepth(root)
	return tag
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if tag == false {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if getSub(left,  right) > 1 {
		tag = false
	}
	return getMax(left, right) + 1
}
func getSub(left ,right int) int {
	if left > right  {
		return left - right
	}
	return right - left
}