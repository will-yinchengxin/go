package six
/*
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
N 叉树的前序遍历

输入：root = [1,null,3,2,4,null,5,6]
输出：[1,3,5,6,2,4]

输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
*/
func preorder(root *Node) []int {
	res := []int{}

	var preo func(node *Node)
	preo = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for i := 0; i < len(node.Children); i++ {
			preo(node.Children[i])
		}
	}
	preo(root)
	return res
}