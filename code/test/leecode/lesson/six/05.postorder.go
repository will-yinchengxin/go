package six
/*
https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
N 叉树的后序遍历

输入：root = [1,null,3,2,4,null,5,6]
输出：[5,6,3,2,4,1]

输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
*/
func postorder(root *Node) []int {
	res := []int{}
	var post func(node *Node)
	post = func(node *Node) {
		if node == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			post(node.Children[i])
		}
		res = append(res, node.Val)
	}
	post(root)
	return res
}