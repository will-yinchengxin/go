package four

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

示例 1:
	输入: [7,5,6,4]
	输出: 5

[4,5,6,7]
9

限制：
	0 <= 数组长度 <= 50000
*/
var Cnt = 0
func ReversePairs(nums []int) int {
	lenN := len(nums)
	if lenN <= 1 {
		return 0
	}

	reversePairsMergeSort(nums, 0, lenN-1)
	print(Cnt)
	return Cnt
}
func reversePairsMergeSort(a []int, begin int, end int)  {
	if begin >= end { // 数组只有单个元素或者没有元素时返回
		return
	}

	mid := begin + (end-begin)/2
	// 先将左边排序好
	reversePairsMergeSort(a, begin, mid)
	// 再将右边排序好
	reversePairsMergeSort(a, mid+1, end)

	reversePairsMerge(a, begin, mid, end) // 两个有序数组进行合并
}

// 归并操作
func reversePairsMerge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := begin
	rightSize := mid + 1
	size := end - begin + 1

	resultAno := make([]int, 0, size)
	for leftSize <= mid && rightSize <= end {
		// 小的元素先放进辅助数组里
		if array[leftSize] <= array[rightSize] {
			resultAno = append(resultAno, array[leftSize])
			Cnt += rightSize - (mid+1)
			leftSize++
		} else {
			resultAno = append(resultAno, array[rightSize])
			rightSize++
		}
	}
	// 将剩下的元素追加到辅助数组后面

	for ; leftSize <= mid; leftSize++ {
		Cnt += end - (mid + 1) + 1
		resultAno = append(resultAno, array[leftSize])
	}
	for ; rightSize <= end; rightSize++ {
		resultAno = append(resultAno, array[leftSize])
	}
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < size; i++ {
		array[begin+i] = resultAno[i]
	}
}
