package eight
/*
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

电话号码的字母组合

给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
	输入：digits = "23"
	输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
	输入：digits = ""
	输出：[]

示例 3：
	输入：digits = "2"
	输出：["a","b","c"]
*/
var resultLC []string
func letterCombinations(digits string) []string {
	if len(digits) == 0 { // 字串为空
		return []string{}
	}
	// 制作 hashMap
	mappings := make(map[int]string, 8)
	mappings[2] = "abc"
	mappings[3] = "def"
	mappings[4] = "ghi"
	mappings[5] = "jkl"
	mappings[6] = "mno"
	mappings[7] = "pqrs"
	mappings[8] = "tuv"
	mappings[9] = "wxyz"

	resultLC = []string{}
	path := make([]string, len(digits)) // 这里使用 零 切片来保存数据，方便后续重新赋值（var slice []int nil切片，slice := []int{} 空切片）
	backtrackLetterCombinations(mappings, digits, 0 , path)
	return resultLC
}
/*
k 当前阶段
path 路径
digits[k] + mappings 确定当前阶段的可选列表
*/
func backtrackLetterCombinations(mappings map[int]string, digits string, k int, path []string) {
	if k == len(digits) { // 终止条件
		str := ""
		for i := 0; i < k; i++ {
			str += path[i]
		}
		resultLC = append(resultLC, str)
		return
	}

	mapping := mappings[int(digits[k] - '0')] // 在 mappings 中获取当前遍历到的路径，如 “23” 遍历至 3，找到 3 对应的所有字母
	for i := 0; i < len(mapping); i++ {
		path[k] = string(mapping[i])
		backtrackLetterCombinations(mappings, digits, k+1, path)
	}
}