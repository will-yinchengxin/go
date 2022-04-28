package four
/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
	输入：n = 2
	输出：2

示例 2：
	输入：n = 7
	输出：21

示例 3：
	输入：n = 0
	输出：1
*/
// 该题目的所有结论与上一题完全相同
var hashMap = make(map[int]int)
func NumWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	if val, ok := hashMap[n]; ok {
		return val
	}
	hashMap[n] = (NumWays(n-1) + NumWays(n-2)) % mod
	return hashMap[n]
}
// 变形, 但是不能连续跳两次2阶
// https://blog.csdn.net/qq_41512783/article/details/109711869
//func NumWaysAno(n int) int {}