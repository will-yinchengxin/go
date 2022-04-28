package three

/*
请根据每日 气温 列表 temperatures，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用0 来代替。

示例 1:
	输入: temperatures = [73,74,75,71,69,72,76,73]
	输出:[1,1,4,2,1,1,0,0]

示例 2:
	输入: temperatures = [30,40,50,60]
	输出:[1,1,1,0]

示例 3:
	输入: temperatures = [30,60,90]
	输出: [1,1,0]
*/
func DailyTemperatures(temperatures []int) []int {
	// 双循环回超出时间限制
	lenTem := len(temperatures)
	if lenTem == 0 {
		return nil
	}
	res := make([]int, lenTem)
	// 这里用来记录 temp 的数组下标
	stack := []int{}
	for i := 0; i < lenTem; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {
			res[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		// 最终将新元素插入至 栈 中
		stack = append(stack, i)
	}
	return res
}
