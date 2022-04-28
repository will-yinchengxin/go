package three

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，
但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
	输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
	输出：true
	解释：我们可以按以下顺序执行：
	push(1), push(2), push(3), push(4), pop() -> 4,
	push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
	输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
	输出：false
	解释：1 不能在 2 之前弹出。

长度 < 1000
元素大小 < 1000
*/

func ValidateStackSequences(pushed []int, popped []int) bool {
	lenPush := len(pushed)
	lenPop := len(popped)
	if lenPop != lenPush || lenPop == 0 || lenPush == 0 {
		return true
	}
	// 辅助栈用来记录
	// 如果 push 和 pop 的首元素相等, 那么 pop 弹出首元素
	// 不相等, 则将 push 中的元素放置到临时栈 stack 中
	// pop 再次与 临时栈中的结果对比
 	stack := []int{}
	for i := 0; i < lenPush; i++ {
		if pushed[i] != popped[0] {
			stack = append(stack, pushed[i])
		} else {
			popped = popped[1:]
		}
		for len(stack) > 0 && popped[0] == stack[len(stack) - 1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}

	for len(popped) > 0 {
		return false
	}
	return true
}

// 5 个数对应了 10 次操作, 5 次 pop 5 次 push
// |__|__|__|__|__|__|__|__|__|__| , 10 个位置, 每次不是 pop, 就是 push,
// 但是不能随便 pop 或者 随便 push, 要看与弹出队列中的数值是否有对应关系,
// pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 第一个位置, pop 是不可以的, 因为没有元素, 只能 push
// 第二个位置, pop 是不可以的, 因为栈顶元素不为 4, 只能 push ..... 依次类推
// 这里的入栈出栈都是操作的 stack
func ValidateStackSequencesAno(pushed []int, popped []int) bool {
	lenPush := len(pushed)
	lenPop := len(popped)
	if lenPop != lenPush || lenPop == 0 || lenPush == 0 {
		return true
	}
	stack := []int{}
	i := 0
	j := 0

	for i < lenPush || j < lenPop {
		// 出栈
		if len(stack) > 0 && j < lenPop && stack[len(stack) - 1] == popped[i] {
			stack = stack[:len(stack) - 1]
			j++
			continue
		}
		if i < lenPush {
			stack = append(stack, pushed[i])
			i++
			continue
		}
		break
	}
	return i == lenPush && j == lenPop
}