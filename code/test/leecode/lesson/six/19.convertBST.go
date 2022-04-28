package six
/*
https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
把二叉搜索树转换为累加树

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node的新值等于原树中大于或等于node.val的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：本题和 1038:https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同

示例 1：
	输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
	输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

示例 2：
	输入：root = [0,null,1]
	输出：[1,null,1]

示例 3：
	输入：root = [1,0,2]
	输出：[3,3,2]

示例 4：
	输入：root = [3,2,4,1]
	输出：[7,9,4,10]

提示：
	树中的节点数介于 0和 104之间。
	每个节点的值介于 -104和104之间。
	树中的所有值 互不相同 。
	给定的树为二叉搜索树。
*/
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}
//---------------------------------------------
//var Sum int = 0 // 这种方式报错
var Sum int
func convertBSTAno(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	/*
		关于这里为什么需要对 var 的 变量做一个赋值, 而直接在最外部 赋值声明如: var Sum = 0, 程序会报错啊
	*/
	Sum = 0 // 这里需要内部初始话这个变量值
	convertBST_R(root)
	return root
}

func convertBST_R(node *TreeNode) {
	if node == nil {
		return
	}
	convertBST_R(node.Right)
	Sum += node.Val
	node.Val = Sum
	convertBST_R(node.Left)
}