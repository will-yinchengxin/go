package seven
/*
https://leetcode-cn.com/problems/diameter-of-binary-tree/

二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树

	  1
	 / \
	2   3
   / \
  4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
*/
var result int
func diameterOfBinaryTree(root *TreeNode) int {
	/*
		看到二叉树， 即可以联想到使用递归来解决问题
			- 左子树的直径
			- 右子树的直径
			二叉树的直径？
		发现子问题的解不能构造出原问题（左子树 + 右子树的 max直径 不能够求解出 最大的直径值）

		将问题转换为：求解树中的每个节点的最大高度
			- 我们寻找的最长路径一定有一个转折点的 /\
			- 经过转折点的最长路径值： LH + RH + 1
			- 计算出 以任意节点为转折点 的最长路径，计算出 max 值
			- 最后总结类似求树的最大高度的题目
	*/
	result = 0
	diameterOfBinaryTree_R(root)
	return result
}
func diameterOfBinaryTree_R(root *TreeNode) int {
	if root == nil {
		return 0
	}

	LHMax := diameterOfBinaryTree_R(root.Left)
	RHMax := diameterOfBinaryTree_R(root.Right)

	diameter := LHMax + RHMax

	if diameter > result {
		result = diameter
	}

	return getMax(LHMax, RHMax) + 1
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
