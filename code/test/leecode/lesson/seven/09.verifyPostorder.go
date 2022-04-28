package seven

/*
https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

二叉搜索树的后序遍历序列

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：
     5
    / \
   2   6
  / \
 1   3

示例 1：
	输入: [1,6,3,2,5]
	输出: false

示例 2：
	输入: [1,3,2,6,5]
	输出: true

提示：
	数组长度 <= 1000
 */
var tag bool
func verifyPostorder(postorder []int) bool {
	/*
		二叉搜索树 后续遍历 有以下特点：
			- 左子树的所有值都小于根节点
			- 右子树的所有值都大于根节点
			- 最后一个节点为根节点
			- 向左遍历，第一个小于根节点的值，即为左子树的开始，右侧即为右子树的结束
			- 一次遍历判别两个区间 与 根节点 的大小关系即可
	*/
	if len(postorder) <= 1 {
		return true
	}
	tag = true
	myVerify(postorder, 0, len(postorder) - 1)
	return tag
}
func myVerify(postorder []int, start int, end int)  {
	if start >= end {
		return
	}
	if tag == false {
		return
	}
	K := start
	for K < end && postorder[K] < postorder[end] { 	// 查找左子树的终点
		K++
	}
	// 验证右子树区间是否均大于 root
	for i := K; i <= end-1; i++ {
		if postorder[i] < postorder[end] {
			tag = false
			return
		}
	}

	// 递归校验左右子树是否均满足条件
	myVerify(postorder, start, K-1)
	myVerify(postorder, K, end-1)
}