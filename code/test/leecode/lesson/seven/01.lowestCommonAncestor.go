package seven
/*
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：
“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：
	输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	输出：3
	解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

示例 2：
	输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
	输出：5
	解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

示例 3：
	输入：root = [1,2], p = 1, q = 2
	输出：1


		 4
	   /   \
	  5     11
	 / \   / \
	1   3 9   12
       / \
       7 8

7, 1 最近公共祖先为 5
*/
var lca *TreeNode
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*
		可以利用回溯来解决问题

		LCA 有这样一些特点， 假设目标为 p q
			1）公共祖先将 p q 分在两边
				- 左边 一个 p 右边一个 q
				- 左边 一个 q 右边一个 p
			2） 特殊情况， q， p 其中一个为公共祖先节点

	*/
	lca = nil
	lowestCommonAncestor_R(root, p, q)
	return lca
}
func lowestCommonAncestor_R(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	leftContains := lowestCommonAncestor_R(root.Left, p, q)
	if lca != nil { // 提前退出
		return 2
	}
	rightContains := lowestCommonAncestor_R(root.Right, p, q)
	if lca != nil { // 提前退出
		return 2
	}
	rootContains := 0
	if root == p || root == q {
		rootContains = 1
	}
	if rootContains == 0 && leftContains == 1 && rightContains == 1 {
		lca = root
	}
	if rootContains == 1 && (leftContains == 1 || rightContains == 1) {
		lca = root
	}
	return rootContains + leftContains + rightContains
}