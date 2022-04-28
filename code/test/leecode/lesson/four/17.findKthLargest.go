package four

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
	输入: [3,2,1,5,6,4] 和 k = 2
	输出: 5

示例2:
	输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
	输出: 4

提示：
	1 <= k <= nums.length <= 104
	-104<= nums[i] <= 104
*/
func findKthLargest(nums []int, k int) int {
	lenN := len(nums)
	if lenN == 0 {
		return 0
	}

	//sort.Ints(nums)
	// 使用归并排序
	findKthLargestMergeSort(nums, 0, lenN-1)
	return nums[lenN-k]
}
func findKthLargestMergeSort(a []int, begin int, end int) {
	if begin >= end {
		return
	}
	// 获取中间节点
	mid := begin + (end -  begin) / 2
	findKthLargestMergeSort(a, begin, mid)
	findKthLargestMergeSort(a, mid+1, end)
	// 进行排序合并
	findKthLargestMerge(a, begin, mid, end)
}
func findKthLargestMerge(a []int, begin, mid, end int) {
	leftSize := begin
	rightSize := mid + 1
	size := end - begin + 1
	// 以中间点为分界点, 左右两指针 同时 向末尾 遍历比较
	resultAno := make([]int, 0 , size)
	for leftSize <= mid && rightSize <= end {
		// 小的元素先放进辅助数组里
		if resultAno[leftSize] < resultAno[rightSize] {
			resultAno = append(resultAno, resultAno[leftSize])
			leftSize++
		} else {
			resultAno = append(resultAno, resultAno[rightSize])
			rightSize++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	if leftSize <= mid {
		resultAno = append(resultAno, resultAno[leftSize:mid+1]...)
	}
	if rightSize <= end {
		resultAno = append(resultAno, resultAno[rightSize:end+1]...)
	}
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < size; i++ {
		resultAno[begin+i] = resultAno[i]
	}
	return
}