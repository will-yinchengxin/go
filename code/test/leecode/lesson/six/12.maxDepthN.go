package six

/*
N 叉树的最大深度

给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

示例 1：
	输入：root = [1,null,3,2,4,null,5,6]
	输出：3
*/
func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}
	dep := 0
	for i := 0; i < len(root.Children); i++ {
		val :=  maxDepthN(root.Children[i])
		if val > dep {
			dep = val
		}
	}
	if dep == 0 {
		return 1
	}
	return dep + 1
}