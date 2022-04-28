package seven
/*
从前序与中序遍历序列构造二叉树

https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

给定一棵树的前序遍历preorder 与中序遍历 inorder。请构造二叉树并返回其根节点。

示例 1:
	Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	Output: [3,9,20,null,null,15,7]

示例 2:
	Input: preorder = [-1], inorder = [-1]
	Output: [-1]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	/*
		拆解为最小子问题
			1） 左子树的 前中 序遍历结果 =》 左子树 leftRoot
			2） 右子树的 前中 序遍历结果 =》 右子树 rightRoot
			root.left  = leftRoot
			root.right = rightRoot
		问题的关键是找到交界点， 并能够准确的划分区域
		pre: 3(root) 9(left) 20 15 7
		ino: 9(left) 3(root) 15 20 7
	*/
	return myBuildTree(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1)
}
func myBuildTree(preorder []int, preorder_start int, preorder_end int, inorder []int, inorder_start int, inorder_end int) *TreeNode {
	if preorder_start > preorder_end {
		return nil
	}
	root := &TreeNode{
		Val: preorder[preorder_start],
	}

	// 在中序遍历的结果中， 查看 preorder[preorder_start](前序遍历中第一个元素为root) 所在位置
	p := inorder_start
	for inorder[p] != preorder[preorder_start] {
		p++
	}
	leftTreeSize := p - inorder_start // 左子树得长度范围

	leftNode := myBuildTree(preorder, preorder_start+1, preorder_start+leftTreeSize, inorder, inorder_start, p -1)
	rightNode := myBuildTree(preorder, preorder_start+1+leftTreeSize, preorder_end, inorder, p+1, inorder_end)
	root.Left = leftNode
	root.Right = rightNode
	return root
}