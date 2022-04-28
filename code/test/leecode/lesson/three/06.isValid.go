package three

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
示例 1：
	输入：s = "()"
	输出：true

示例2：
	输入：s = "()[]{}"
	输出：true

示例3：
	输入：s = "(]"
	输出：false
*/
// 申请两个栈, 来回倒腾, 获取其中的值与已有
func isValid(s string) bool {
	lenS := len(s)
	if lenS == 0 || lenS == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < lenS; i++ {
		if pairs[s[i]] > 0 {
			// 这一步是判断的关键
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
