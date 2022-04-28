package four

/*
实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

示例 1：
	输入：x = 2.00000, n = 10
	输出：1024.00000

示例 2：
	输入：x = 2.10000, n = 3
	输出：9.26100

示例 3：
	输入：x = 2.00000, n = -2
	输出：0.25000
	解释：2-2 = 1/2 * 2 = 1/4 = 0.25

提示：
	-100.0 < x < 100.0
	-2^31 <= n <= 2^31-1
	-10^4 <= x^n <= 10^4
*/
// 因为指数 n 量级在 10 亿次左右，所以利用循环一定会报错
// 递归实现（快速幂）
// 核心逻辑：X^n = X^(2/n) + X^(2/n)
func myPow(x float64, n int) float64 {
	// 指数可为正可为负
	if (n >= 0) {
		return myPow_r(x, n)
	} else {
		// 指数 n 为负数， 正常情况 1/x^n
		// 特殊情况 x=2.000，n=-2147483648（-2147483648 ~ 2147483647）
		// return 1/myPow_r(x,-1*n), 此时就越界了
		// 使用 + 1 技巧， 在结果输出后在 * x 就能保证不发生越界问题
		return 1/(myPow_r(x, -1*(n+1))*x)
	}
}

func myPow_r(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	halfPow := myPow_r(x, n/2)
	if n % 2 == 1 {
		return halfPow * halfPow * x
	} else {
		return halfPow * halfPow
	}
}