package eight

import (
	"strconv"
	"strings"
)

/*
https://leetcode-cn.com/problems/restore-ip-addresses/

复原 IP 地址

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成
你 不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

示例 1：
	输入：s = "25525511135"
	输出：["255.255.11.135","255.255.111.35"]

示例 2：
	输入：s = "0000"
	输出：["0.0.0.0"]

示例 3：
	输入：s = "101023"
	输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

提示：
	0 <= s.length <= 20
	s 仅由数字组成
*/
var resultRestoreIpAddresses []string
func restoreIpAddresses(s string) []string {
	resultRestoreIpAddresses = []string{}
	path := []string{}
	backtrackRestoreIpAddresses(s, 0, path)
	return resultRestoreIpAddresses
}
/*
originalStr 原始字串
k 阶段
path 路径
*/
func backtrackRestoreIpAddresses(originalStr string, k int, path []string) {
	if k == len(originalStr) && len(path) == 4 { // 结束条件
		str := ""
		for i := 0; i < 4; i++ {
			str = str + path[i] + "."
		}
		str = strings.Trim(str, ".")
		resultRestoreIpAddresses = append(resultRestoreIpAddresses, str)
		return
	}
	if len(path) > 4 || k == len(originalStr) { // 提前终止
		return
	}
	val := 0
	/*
		决策阶段, 这里范围为 0 ~ 255
		- 放置一位数字
		- 前导零直接跳过
		- 放置两位数字
		- 放置三位数字
	*/
	if k < len(originalStr) { // 每次放置一位数
		val = int(originalStr[k]-'0')
		path = append(path, strconv.Itoa(val))
		backtrackRestoreIpAddresses(originalStr, k+1, path)
		path = path[:len(path)-1]
	}
	if originalStr[k] == '0' {
		return
	}
	if k+1 < len(originalStr) { // 每次放置两位数字
		val = val*10 + int(originalStr[k+1]-'0')
		path = append(path, strconv.Itoa(val))
		backtrackRestoreIpAddresses(originalStr, k+2, path)
		path = path[:len(path)-1]
	}
	if k+2 < len(originalStr) { // 每次放置三位数字
		val = val*10 + int(originalStr[k+2]-'0')
		if val <= 255 {
			path = append(path, strconv.Itoa(val))
			backtrackRestoreIpAddresses(originalStr, k+3, path)
			path = path[:len(path)-1]
		}
	}
}