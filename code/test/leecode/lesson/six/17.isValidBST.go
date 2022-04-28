package six

/*
https://leetcode-cn.com/problems/validate-binary-search-tree/

验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
	节点的左子树只包含 小于 当前节点的数。
	节点的右子树只包含 大于 当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
	输入：root = [2,1,3]
	输出：true

示例 2：
	输入：root = [5,1,4,null,null,3,6]
	输出：false
	解释：根节点的值是 5 ，但是右子节点的值是 4 。

提示：
	树中节点数目范围在[1, 104] 内
	-231 <= Node.val <= 231 - 1
*/
var IsValid bool

func isValidBST(root *TreeNode) bool {
	/*
		1) 二叉查找树中序遍历的结果是有序的(可以利用这一点)
		2) 这里使用查询左右子树中 最大值 和 最小值 方式解决
	*/

	// 方式一
	//res := []int{}
	//var pore func(node *TreeNode)
	//pore = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	pore(node.Left)
	//	res = append(res, node.Val)
	//	pore(node.Right)
	//}
	//pore(root)
	// 判断 res 的有序性

	IsValid = true
	// 方式二
	// 查找左右子树的 最大值 和 最小值
	if root == nil {
		return true
	}
	dfs(root)
	return IsValid
}
func dfs(node *TreeNode) []int {
	max := node.Val
	min := node.Val
	/*
		1) 左边的最大值 小于 node.Val, 整体一定小于 node.Val
		2) 右边的最小值 大于 node.Val, 整体一定大于 node.Val
		3) 最终找到 左边的 最小值 与 右边的 最大值
	*/
	// 寻找左边的最大值
	if node.Left != nil {
		LeftData := dfs(node.Left)
		if IsValid == false { // 提前退出
			return nil
		}
		if LeftData[1] >= node.Val { // 左边的最大值 还大于 node
			IsValid = false
			return nil
		}
		min = LeftData[0] // 将左边的最小值存储
	}
	// 寻找右边的最小值
	if node.Right != nil {
		RightData := dfs(node.Right)
		if IsValid == false { // 提前退出
			return nil
		}
		if RightData[0] <= node.Val { // 右边最小值 还小于 node
			IsValid = false
			return nil
		}
		max = RightData[1] // 将右边的最大值存储
	}
	return []int{min, max}
}