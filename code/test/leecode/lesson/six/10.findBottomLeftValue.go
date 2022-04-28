package six
/*
https://leetcode-cn.com/problems/find-bottom-left-tree-value/

找树左下角的值

给定一个二叉树的 根节点 root，请找出该二叉树的最底层最左边节点的值。
假设二叉树中至少有一个节点。

示例 1:
	输入: root = [2,1,3]
	输出: 1

示例 2:
	输入: [1,2,3,4,null,5,6,null,null,7]
	输出: 7

提示:
	二叉树的节点个数的范围是 [1,104]
	-231<= Node.val <= 231- 1
*/
// 按层遍历, 最后一层的第一个元素
func findBottomLeftValueAno(root *TreeNode) int {
	// 按 层 从右向左遍历, 最后一个元素, 就是要找的元素
	helpList := make([]*TreeNode, 0)
	helpList = append(helpList, root)
	res := []int{}
	for len(helpList) != 0 {
		elem := helpList[0]
		helpList = helpList[1:]
		res = append(res, elem.Val)
		if elem.Right != nil {
			helpList = append(helpList, elem.Right)
		}
		if elem.Left != nil {
			helpList = append(helpList, elem.Left)
		}

	}
	return res[len(res)-1]
}