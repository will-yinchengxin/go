package six
/*
https://leetcode-cn.com/problems/successor-lcci/
后继者

设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。

如果指定节点没有对应的“下一个”节点，则返回null。

示例 1:
	输入: root = [2,1,3], p = 1

	  2
	 / \
	1   3

	输出: 2

示例 2:
	输入: root = [5,3,6,2,4,null,null,1], p = 6

		  5
		 / \
		3   6
	   / \
	  2   4
	 /
	1

	输出: null
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var tagInorder *TreeNode
	var dfs func(*TreeNode)
	flag := false

	dfs = func(node *TreeNode) {
		if node != nil && tagInorder == nil {
			dfs(node.Left)
			if node == p {
				flag = true
			} else if flag {
				// 这里使用标记记录完成后,一定及时修正 flag,不然这个值会被变更为根节点
				tagInorder, flag = node, false
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return tagInorder
}