package nine
/*
https://leetcode-cn.com/problems/jump-game-iii/

跳跃游戏 III

这里有一个非负整数数组arr，你最开始位于该数组的起始下标start处。当你位于下标i处时，你可以跳到
i + arr[i] 或者 i - arr[i]。
请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
注意，不管是什么情况下，你都无法跳到数组之外。

示例 1：
	输入：arr = [4,2,3,0,3,1,2], start = 5
	输出：true
	解释：
	到达值为 0 的下标 3 有以下可能方案：
	下标 5 -> 下标 4 -> 下标 1 -> 下标 3
	下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3

示例 2：
	输入：arr = [4,2,3,0,3,1,2], start = 0
	输出：true
	解释：
	到达值为 0 的下标 3 有以下可能方案：
	下标 0 -> 下标 4 -> 下标 1 -> 下标 3

示例 3：
	输入：arr = [3,0,2,1,2], start = 2
	输出：false
	解释：无法到达值为 0 的下标 1 处。

提示：
	1 <= arr.length <= 5 * 10^4
	0 <= arr[i] < arr.length
	0 <= start < arr.length
*/
var reached bool
var visitedCanReach []bool
func canReach(arr []int, start int) bool {
	reached = false
	lenArr := len(arr)
	visitedCanReach = make([]bool, lenArr)
	dfsCanReach(arr, start)
	return reached
}
func dfsCanReach(arr []int, curI int) {
	if reached || arr[curI] == 0 {
		reached = true
		return
	}
	visitedCanReach[curI] = true
	// 开始左右移动
	move2Left := curI - arr[curI]
	if move2Left >= 0 && move2Left < len(arr) && !visitedCanReach[move2Left] {
		dfsCanReach(arr, move2Left)
	}
	move2Right := curI + arr[curI]
	if move2Right >= 0 && move2Right < len(arr) && !visitedCanReach[move2Right] {
		dfsCanReach(arr, move2Right)
	}
}