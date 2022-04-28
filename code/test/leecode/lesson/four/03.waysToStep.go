package four
/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:
	 输入：n = 3
	 输出：4
	 说明: 有四种走法

示例2:
	 输入：n = 5
	 输出：13

提示:
	n范围在[1, 1000000]之间
*/
// 递归 + 备忘录
/*
	1) 查找 子问题 的求解是否能拼凑出 原问题 的解
	2) 看第一步怎么走，总结出递推公式
		可以 1 步， K1.. 走了 1 个 还剩 F(n-1) 种
		可以 2 步， K2.. 走了 2 个 还剩 F(n-2) 种
		可以 3 步， K3.. 走了 3 个 还剩 F(n-3) 种
		F(n) = F(n-1) + F(n-2) + F(n-3)
	3) 递归终止条件，也就是最小子问题
		F(0).. 显然不可以， 也不存这种情况
		F(1).. 不能通过式子推到， 只能有1种方式
		F(2).. 不能通过式子推到， 只能有2种方式
		F(3).. 不能通过式子推到， 只能有4种方式
		F(4) = F(3) + F(2) + F(1) 符合公式， > 3 这里就可以使用递推公式
*/
// 空间复杂度还是与树的 max(high) 有关也就是函数调用栈，所以时间复杂度为 O(n)
// 时间复杂度为 O(n), 也会超时？？
/*
	加了备忘录为什么会超时？
		数据量在 100 万左右， 那么发起 100 万次的函数调用， 是非常耗时的会引起超时现象的
		即使通过：
			执行用时：388 ms, 在所有 Go 提交中击败了 6.37% 的用户
			内存消耗：163.8 MB, 在所有 Go 提交中击败了 6.37% 的用户
		内存的消耗也是蛮大的
*/
func waysToStep(n int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}
	if n == 3 {
		return 4
	}
	if val, ok := hashMap[n]; ok {
		return val
	}
	hashMap[n] = (waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)) % mod
	return hashMap[n]
}

// 非递归实现， 解决 100 次函数调用的问题
func wayToStepAno(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	helpSel := make([]int, n + 1)
	helpSel[1] = 1
	helpSel[2] = 2
	helpSel[3] = 4
	for i := 4; i <= n ; i++ {
		// 计算公式 ???
		helpSel[i] = (helpSel[i-1]%mod + helpSel[i-2]%mod + helpSel[i-3])%mod
	}
	return helpSeli[n]
}
// 发现每次的 helpSel[i]只与前三个值有关，没有必要依次申请 make([]int, n + 1) 的切片
func wayToStepAnoAno(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	a, b, c, d := 1, 2, 4, 0
	for i := 4; i <= n ; i++ {
		// 这里保证数值不会溢出
		d = (a%mod + b%mod + c)%mod
		a = b
		b = c
		c = d
	}
	return d
}