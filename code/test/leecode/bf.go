package leecode

/*
BF 算法寻找字符串
	stringA := "abvfh"
	lenA := len(stringA)
	stringB := "fh"
	lenB := len(stringB)
	fmt.Println(BF(stringA, lenA, stringB, lenB))
*/
func BF(stringA string, lenA int, stringB string, lenB int) int {
	// BF 算法中 n-m+1 代表子串的个数
	for i := 0; i < (lenA - lenB + 1); i++ {
		j := 0
		for j < lenB {
			if stringA[i+j] != stringB[j] {
				break
			}
			j++
		}
		if j == lenB {
			return i
		}
	}
	return -1
}
