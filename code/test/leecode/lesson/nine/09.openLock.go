package nine

import "fmt"

/*
https://leetcode-cn.com/problems/open-the-lock/

打开转盘锁

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：
例如把 '9' 变为'0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadEnds 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

示例 1:
	输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
	输出：6
	解释：
	可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
	注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
	因为当拨动到 "0102" 时这个锁就会被锁定。

示例 2:
	输入: deadends = ["8888"], target = "0009"
	输出：1
	解释：把最后一位反向旋转一次即可 "0000" -> "0009"。

示例 3:
	输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
	输出：-1
	解释：无法旋转到目标数字且不被锁定。

提示：
	1 <=deadends.length <= 500
	deadends[i].length == 4
	target.length == 4
	target 不在 deadends 之中
	target 和 deadends[i] 仅由若干位数字组成
*/
func openLock(deadends []string, target string) int {
	/*
		为题的类型是一个位图, 且寻找最优解, 实质是一个图上的 bfs 的题目
		图上的 bfs 可以类比 树 上的 按层遍历,
	*/
	deadset := make(map[string]bool, 0) // 将 deadEnds 重置为 hash 表方便快速查找
	for _, d := range deadends {
		deadset[d] = true
	}
	if deadset["0000"] { // 初始即为爆炸情况
		return -1
	}
	queue := make([]string, 0)
	visited := make(map[string]bool, 0)
	queue = append(queue, "0000")
	visited["0000"] = true
	depth := 0 // 记录最下次数也就是层级
	for len(queue) > 0 {
		size := len(queue)
		k := 0
		for k < size {
			node := queue[0]
			queue = queue[1:]
			k++
			if node == target {
				return depth
			}
			newnodes := genNewNode(node) // 获取拓展出的所有新密码
			for _, newnode := range newnodes {
				if visited[newnode] || deadset[newnode] {
					continue
				}
				queue = append(queue, newnode)
				visited[newnode] = true
			}
		}
		depth++
	}
	return -1
}
func genNewNode(node string) []string {
	newnodes := make([]string, 0)
	change := []int{-1, 1}
	for i := 0; i < 4; i++ {
		for k := 0; k < 2; k++ {
			newNode := make([]byte, 4)
			for j := 0; j < i; j++ {
				newNode[j] = node[j]
			}
			for j := i + 1; j < 4; j++ {
				newNode[j] = node[j]
			}
			newC := fmt.Sprintf("%d", (int(node[i]-'0')+change[k]+10)%10)
			newNode[i] = newC[0]
			newnodes = append(newnodes, string(newNode))
		}
	}
	return newnodes
}