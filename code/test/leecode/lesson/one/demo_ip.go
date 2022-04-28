package one

import (
	"math"
	"strconv"
	"strings"
)

// 判断IP是否合法
func CheckIp(ip string) bool {
	if len(ip) == 0 {
		return false
	}
	ipSlice := strings.Split(ip, ".")
	if len(ipSlice) != 4 {
		return false
	}
	// 循环校验每一段
	for i := 0; i < 4; i++ {
		checkRes := checkSegment(ipSlice[i])
		if !checkRes {
			return false
		}
	}

	return true
}

// 校验每一段ip
func checkSegment(segment string) bool {
	// 校验是否全为空格
	lenSegment := len(segment)
	n := 0
	for i := 0; i < lenSegment; i++ {
		if segment[i] == ' ' {
			n++
		}
	}
	if n == lenSegment {
		return false
	}
	// 将字符串转化为数字, 且不为 2bd 或者 12 1 这种形式
	newSeg := strings.Trim(segment, " ")
	newLen := len(newSeg)
	digt := 0
	for i := 0; i < newLen; i++ {
		val, err := strconv.Atoi(string(newSeg[i]))
		if err != nil {
			return false
		}

		digt = digt + val * int(math.Pow10(newLen - i - 1))
	}


	// ip 每段最大值为 255
	if digt > 255 {
		return false
	}
	return true
}