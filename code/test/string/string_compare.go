package string

import "strings"

func String(a, b string) bool {
	if a == b {
		return true
	}
	return false
}

func Compare(a, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	}
	return false
}

func EqualFold(a, b string) bool {
	if strings.EqualFold(a, b) {
		return true
	}
	return false
}