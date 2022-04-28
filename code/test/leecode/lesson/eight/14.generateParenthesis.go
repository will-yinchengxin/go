package eight
/*
https://leetcode-cn.com/problems/generate-parentheses/

括号生成

数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
	输入：n = 3
	输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
	输入：n = 1
	输出：["()"]

提示：
	1 <= n <= 8
*/
var resultGenerateParenthesis []string
func generateParenthesis(n int) []string {
	/*
		这个题目有些 规律
		- 左括号可以一直放置
		- 右括号必须在左括号放置后才可以放置
		- 右括号数量 >= 左括号数量
	*/
	resultGenerateParenthesis = []string{}
	path := make([]string, 2*n) // 括号的数量是确定的
	backtrackGenerateParenthesis(n, n, 0, path, n)
	//backtrackGenerateParenthesis(0, 0, 0, path, n)
	return resultGenerateParenthesis
}
/*
leftBrackets, rightBrackets 左右括号的数量, 推断出可选列表
k 阶段(n*2)
path 路径
n 帮助参数
*/
func backtrackGenerateParenthesis(leftBrackets, rightBrackets int, k int, path []string, n int) {
	if k == 2*n { // 终止条件
		str := ""
		for i := 0; i < len(path); i++ {
			str += path[i]
		}
		resultGenerateParenthesis = append(resultGenerateParenthesis, str)
		return
	}
	if leftBrackets < rightBrackets {
		path[k] = ")"
		backtrackGenerateParenthesis(leftBrackets, rightBrackets-1, k+1, path, n)
	}
	if leftBrackets > 0 {
		path[k] = "("
		backtrackGenerateParenthesis(leftBrackets-1, rightBrackets, k+1, path, n)
	}
	/*
		if leftBrackets < n {
			path[k] = "("
			backtrackGenerateParenthesis(leftBrackets+1, rightBrackets, k+1, path, n)
		}
		if leftBrackets > rightBrackets {
			path[k] = ")"
			backtrackGenerateParenthesis(leftBrackets, rightBrackets+1, k+1, path, n)
		}
	*/
}