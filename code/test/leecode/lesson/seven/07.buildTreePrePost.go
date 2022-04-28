package seven
/*
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

根据前序和后序遍历构造二叉树

返回与给定的前序和后序遍历匹配的任何二叉树。

pre 和 post 遍历中的值是不同的正整数。

示例：
	输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
	输出：[1,2,3,4,5,6,7]

提示：
	1 <= pre.length == post.length <= 30
	pre[]和post[]都是1, 2, ..., pre.length的排列
	每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个
*/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	/*
		拆解为最小子问题
			1） 左子树得 前后 遍历结果 =》 左子树
			1） 右子树得 前后 遍历结果 =》 右子树

			root.left = leftRoot
			root.right = rightRoot
		但是拆解到最后发现，除了能够快速确定根节点以外， 不能够准确的区分左右子树

		那么另辟蹊径， 发现前序遍历的 root 节点后的一个节点很神奇，它可能是左子树的根节点， 也可能是右子树的根节点
		- 存在 左子树 或者 左子树和右子树 都存在 为 左子树
		- 只存在 右子树 的时候 为右子树

		存在一种特殊情况:
			  1
			 /
			2
		   / \
		   3  4
		   前： 1 2 3 4
		   后： 3 4 2 1

			1
			 \
			 2
			/ \
			3  4
			前： 1 2 3 4
			后： 3 4 2 1
		解决办法：这种情况会有两种解答，那么可以统一使用左子树输出即可
	*/
	return myBuildTree_Ano(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}
func myBuildTree_Ano(preorder []int, i int, j int, postorder[]int, p int, r int) *TreeNode {
	if i > j {
		return nil
	}
	root := &TreeNode{
		Val: preorder[i],
	}
	if i == j {  // 与前中遍历内容不一致的位置, 避免后续逻辑造成 数组越界 的错误
		return root
	}

	q := p
	for postorder[q] != preorder[i+1] {
		q++
	}
	leftTreeSize := q - p + 1// 左子树得长度范围（左子树内容）

	leftNode := myBuildTree_Ano(preorder, i+1, i+leftTreeSize, postorder, p, q)
	rightNode := myBuildTree_Ano(preorder, i+1+leftTreeSize, j, postorder, q+1, r-1)
	root.Left = leftNode
	root.Right = rightNode
	return root
}