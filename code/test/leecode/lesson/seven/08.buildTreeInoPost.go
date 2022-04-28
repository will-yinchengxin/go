package seven
/*
https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

从中序与后序遍历序列构造二叉树

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
	你可以假设树中没有重复的元素。

例如，给出
	中序遍历 inorder =9（left）,3（root）,6, 15, 8 ,20,7
	后序遍历 postorder = 9（left）, 6 , 8 , 15, 7, 20, 3（root）

返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
  /  \
 6   8
*/
func buildTreeInoPost(inorder []int, postorder []int) *TreeNode {
	/*
		如破口为找 root
		中序遍历的 左子树 右子树
		后序遍历的 左子树 右子树

		中序遍历 inorder =9（left）,3（root）,6, 15, 8 ,20,7
		后序遍历 postorder = 9（left）, 6 , 8 , 15, 7, 20, 3（root）
	*/
	return myBuildTreeAno(inorder, 0, len(inorder) - 1, postorder, 0, len(inorder) - 1)
}
func myBuildTreeAno(inorder []int, inorder_start int, inorder_end int, postorder []int, postorder_start int, postorder_end int) *TreeNode {
	if inorder_start > inorder_end {
		return nil
	}
	root := &TreeNode{
		Val: postorder[postorder_end],
	}

	// 在中序遍历的结果中， 查看 preorder[preorder_start](前序遍历中第一个元素为root) 所在位置
	p := inorder_start
	for inorder[p] != postorder[postorder_end] {
		p++
	}
	leftTreeSize := p - inorder_start // 左子树得长度范围 3

	leftNode := myBuildTreeAno(inorder, inorder_start, p-1, postorder, postorder_start, postorder_start+leftTreeSize-1)
	rightNode := myBuildTreeAno(inorder, p+1, inorder_end, postorder, leftTreeSize+postorder_start, postorder_end-1)
	root.Left = leftNode
	root.Right = rightNode
	return root
}