package six
/*
https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

N 叉树的层序遍历

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

示例 1：
	输入：root = [1,null,3,2,4,null,5,6]
	输出：[[1],[3,2,4],[5,6]]

示例 2：
	输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
	输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

提示：
	树的高度不会超过1000
	树的节点总数在 [0,10^4] 之间。
*/
func levelOrderN(root *Node) [][]int {
	if root == nil {
		return nil
	}
	tmp := []int{}
	res := [][]int{}
	arrList := make([]*Node, 0)
	arrList = append(arrList, root)
	size := 1
	for len(arrList) != 0 {
		currentSize := 0
		for i := 0; i < size; i++ {
			tmp = append(tmp, arrList[i].Val)
			for j := 0; j < len(arrList[i].Children); j++ {
				if arrList[i].Children[j] != nil {
					arrList = append(arrList, arrList[i].Children[j])
					currentSize++
				}
			}

		}
		arrList = arrList[size:]
		res = append(res, tmp)
		tmp = []int{}
		size = currentSize
	}
	return res
}