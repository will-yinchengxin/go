package six

/*
https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
	给定二叉树:[3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7

返回其层次遍历结果：
	[
	  [3],
	  [20,9],
	  [15,7]
	]
*/
// 取巧方式, 奇数不变, 偶数反转
func levelOrderSnake(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	tmp := []int{}
	res := [][]int{}

	arrList := make([]*TreeNode, 0)
	arrList = append(arrList, root)
	size := 1

	tag := 1
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
		if tag % 2 == 1 {
			tmp = rev(tmp)
		}
		arrList = arrList[size:]
		res = append(res, tmp)
		tmp = []int{}
		size = currentSize
		tag++
	}
	return res
}
func rev(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

//-----------------------------------------------------------------------------------
// tag % 2 == 1 从左到右
// tag % 2 == 0 从右到左
// 对于这种 一会 从左到右 一会 从右到左 的我们采用 栈 来解决
// 使用两个栈来解决 或者使用 双端队列
//func levelOrderSnakeAno(root *TreeNode) [][]int {}
