package six
/*
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
给你二叉树的根节点 root ，返回其节点值的 层序遍历。（即逐层地，从左到右访问所有节点）。

示例 1：
	输入：root = [3,9,20,null,null,15,7]
	输出：[[3],[9,20],[15,7]]

示例 2：
	输入：root = [1]
	输出：[[1]]

示例 3：
	输入：root = []
	输出：[]
*/
func levelOrderByNil(root *TreeNode) [][]int {
	// 使用 nil 方式
	/*
		1) 遍历每一层给一个 nil 标识
		2) 通过 nil 标识进行切片的拆分
	*/
	if root == nil {
		return nil
	}
	tmp := []int{}
	res := [][]int{}

	arrList := make([]*TreeNode, 0)
	arrList = append(arrList, root)
	arrList = append(arrList, nil)
	for len(arrList) != 0 {
		if arrList[0] == nil {
			break
		}
		for arrList[0] != nil {
			elem := arrList[0]
			arrList = arrList[1:]
			tmp = append(tmp, elem.Val)

			if elem.Left != nil {
				arrList = append(arrList, elem.Left)
			}
			if elem.Right != nil {
				arrList = append(arrList, elem.Right)
			}
		}
		arrList = arrList[1:]
		arrList = append(arrList, nil)
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func levelOrderBySize(root *TreeNode) [][]int {
	// 使用 size 方式
	/*
		1) 记录每一层的size
		2) 遍历过程中 < size
		3) 同时拓展出新的size
	*/
	if root == nil {
		return nil
	}
	tmp := []int{}
	res := [][]int{}
	arrList := make([]*TreeNode, 0)
	arrList = append(arrList, root)
	size := 1
	for len(arrList) != 0 {
		currentSize := 0
		for i := 0; i < size; i++ {
			tmp = append(tmp, arrList[i].Val)
			if arrList[i].Left != nil {
				arrList = append(arrList, arrList[i].Left)
				currentSize++
			}
			if arrList[i].Right != nil {
				arrList = append(arrList, arrList[i].Right)
				currentSize++
			}
		}
		arrList = arrList[size:]
		res = append(res, tmp)
		tmp = []int{}
		size = currentSize
	}
	return res
}