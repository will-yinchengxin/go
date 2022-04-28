package four

/*
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：
	输入： arr = [1,3,5,7,2,4,6,8], k = 4
	输出： [1,2,3,4]

提示：
	0 <= len(arr) <= 100000
	0 <= k <= min(100000, len(arr))
*/
func smallestK(arr []int, k int) []int {
	lenN := len(arr)
	if lenN == 0 {
		return nil
	}

	//sort.Ints(arr)
	// 使用归并排序
	smallestKMergeSort(arr, 0, lenN-1)
	return arr[0:k]
}
func smallestKMergeSort(a []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	smallestKMergeSort(a, start, mid)
	smallestKMergeSort(a, mid+1, end)

	smallestKMerge(a, start, mid, end)
}
func smallestKMerge(a []int, start, mid, end int) {
	left := start
	right := mid + 1
	size := end - start + 1

	// 结果
	res := make([]int, 0, size)
	for left <= mid && right <= end {
		if a[left] < a[right] {
			res = append(res, a[left])
			left++
		} else {
			res = append(res, a[right])
			right++
		}
	}
	if left <= mid {
		res = append(res, a[left:mid+1]...)
	}
	if right <= end {
		res = append(res, a[right:end+1]...)
	}
	for i := 0; i < size; i++ {
		a[start+i] = res[i]
	}
	return
}
