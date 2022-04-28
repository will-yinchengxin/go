package four

import (
	"fmt"
)

/*
	1) 递归
		- 一些题目的题解, 要使用 备忘录 的功能, 避免重复调用, 造成堆栈溢出

		- 每个函数都对应一个栈帧, 进入函数便会生成, 栈帧包括: 1)参数 2)局部变量 3)返回地址

		- 怎样确定一个问题可以使用递归解决
			- 规模大的问题和规模小的问题解决思路一样, 仅规模不同,
			- 利用子问题组合可以得到原问题的解
			- 存在最小子问题, 可以直接返回结果

		- 使用递归的正确姿势: 假设子问题 B C 已经解决, 然后考虑父问题 A, 找出 递推公式 及 终止条件
			模板:
				func f(param) {
					if .. return // 终止条件
					.... 		 // 前置逻辑
					递归函数		 // 子问题
					....         // 是否需要手动恢复
				}

		- 切记不要试图梳理清楚递归执行的整个过程, 这实际上进入了思维误区

		- 递归 分析时间复杂度方式:
			- 递推公式(可化简的式子, 适合用递推公式)
			- 递归树(具有普适性)
		- 向下为递的过程, 向上为归的过程
	2) 排序:
		排序指标
			- 时间复杂度
			- 空间复杂度
			- 原地排序
			- 稳定性

		原地 & 非原地:
			能不能在原始数组上, 通过"位移"实现排序, 能为 原地 排序, 不能则为 非原地
			是不是需要申请新的空间, 来储蓄排序数据

		空间复杂度与原地性的关系:
			原地排序 算法的空间复杂度并不一定是 O(1) -- 快排
			空间复杂度为O(1) 的排序算法肯定是 原地排序

		递归 分析时间复杂度方式:
			- 递推公式(可化简的式子, 适合用递推公式)
			- 递归树(具有普适性)

		平均时间复杂度:
			最好/坏 时间复杂度都不能更好的表示时间复杂度, 所以我们引入了平均时间复杂度
			例: 查找 x 在数组中位置, 有 n + 1 种情况, 在 0 ~ n-1 和 不在数组种
				我们把每种需要遍历的元素个数列加起来, 除以 n+1
				1 + 2 + 3 + ...+n + n(不存在数组种也需要遍历所有)/n+1 = ((1+n)n/2  + n) / n+1 = n(n+3)/2(n+1)
				时间复杂度: 省略系数, 低阶, 常量, 即时间复杂度为 O(n)

		tips:
			1) 快速排序, 空间复杂度需要考虑递归函数调用栈的消耗, 原地不需要考虑
			2) 等差数列求和公式: (首数+尾数)×个数÷2
			3) 等比数列求和公式: 2^n - 1 (1 2 4 8 = 2^4 - 1)
			4) go 中排序函数 sort.Slice
					sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]}) // 二维数组进行排序
						sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] }) // 字符串转切片进行排序
			5) https://www.cs.usfca.edu/~galles/visualization/Algorithms.html  数据结构可视化界面
*/

// 冒泡排序(升序)
// 包含两个过程, 比较 和 交换, 冒泡排序中只会交换两个相邻的元素, 每交换一次, 有序度 + 1
// 所以冒泡排序中, 总的交换次数是确定的, 即为 逆序度
// 方式是通过交换, 每次遍历将最大元素置为末尾
// 平均交换次数: k1 = n*(n-1)/4 (取最大值 与 最小值 的 中间值 作为 平均值)
// 平均比较次数: k2 < n^2 && k2 > k1=n*(n-1)/4
// 冒泡排序算法的平均时间复杂度: O(k1 + k2) = O(n^2)
// 空间复杂度 O(1)
func bubbleSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		flag := false                // 提前退出冒泡的标志
		for j := 0; j < n-i-1; j++ { // 将最大值放置 末尾
			if a[j] > a[j+1] {
				tmp := a[j+1] // j 与 j+1 交换, 循环条件为 n - j - 1
				a[j+1] = a[j]
				a[j] = tmp
				flag = true // 表示有数据交换
			}
		}
		if !flag { // 没有交换, 证明有序, 直接退出
			break
		}
	}
	return a
}

// 插入排序(升序)
// 从无序的序列中 找出元素放置在有序序列中, 每次那无序序列数据与有序序列的数据依次对比插入
// 时间复杂度: 最好O(n), 最坏O(n^2) 平均时间复杂度为 O(n^2)
// 空间复杂度 O(1)
// 是原地排序: 是的, 没有申请多余的内存空间, 空间复杂度为O(1) 的排序算法肯定是 原地排序
// 是稳定排序: a[end] > num 才交换
func insertSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ { // [0, i} 已经排序
		num := a[i]
		end := i - 1
		for end >= 0 && a[end] > num { // 查找插入的位置
			a[end+1] = a[end]  // 数据移动
			end--
		}
		a[end+1] = num // 插入数据
	}
	return a
}

// 选择排序(升序)
// 类似 插入排序, 将整个数组分为已排序区间和未排序区间, 不同在于每次从未排序区间找最小元素将其放置已排序的末尾
// 时间复杂度: 最好/坏/平均 时间复杂度 都是 O(n^2)
// 空间复杂度 O(1)
// 是否为原地排序: 是
// 是否为稳定排序: 否, 每次都要执行交换 a[i], a[index] = a[index], a[i]
func selectSort(a []int, n int) []int{
	var index int
	//外层循环，每次得到 data[i] ~ data[size-1] 之间的最小值，存放在 data[i]
	for i := 0; i < n; i++ {
		index = i
		/*
			内存循环，暂且将data[i]当做 data[i] ~ data[size-1] 之间的最小值
			使用index记录这个最小值的标号
			依次使用data[index] 和 data[i+1] ~ data[size-1] 之间的所有元素作比较
			如果 data[index] > data[j]，则 index = j
			一趟内存循环执行完，index 中就已经记录了 data[i] ~ data[size-1] 之间最小元素的标号
		*/
		for j := i + 1; j < n; j++ {
			if a[index] > a[j] {
				index = j
			}
		}
		//将data[i] 与 data[index] 交换，从而将 data[i] ~ data[size-1] 中的最小值存放在 data[i] 中
		a[i], a[index] = a[index], a[i]
	}
	return a
}

// 归并排序
// 采用了分治的思想, 采用递归来实现
// 将待排序数组分成 前后 两个部分, 然后对前后两个部分分别排序, 再将排序好的两个部分合并在一起, 整个数组就有序了
// merge_sort(p, r) = merge( merge_sort(p,q), merge_sort(q+1,r) ), q = (p+r)/2
// 终止条件 p >= r, 不再进行分解

// 时间复杂度:
	// 1)递推公式分析:
	// T(n) = T(n/2) + T(n/2) + n, T(1) = C, 将 T(n) 化解为只包含 n, C 的表达式
	//
	// 			 T(n) = 2*T(n/2) + n = 2*(2*T(n/4)+n/2) + n = 4*T(n/4)+2*n ...
	//                = 2^k*T(n/2^k) + k*n ...
	// 当 n/2^k 接近 1 的时候, 得到 k = logn, 得到 (2^logn)T(C) + logn*n = nT(c) + logn*n ≈ n*logn
	//
	// 2) 递归树进行分析:
	// 	  递的过程不耗时, 归的过程比较耗时, 递归树的每次层归起来的耗时都是 n, 那么时间复杂度也就与 二叉树的高度 有关
	//    因为二叉树的时满二叉树, 高度也就为 logn, 所以时间复杂度就为 n * logn

// 空间复杂度:
//   空间复杂度 = 函数调用栈 + tmp数组长度, 调用栈最大值也就是二叉树的高度 logn, tmp最大值为 n, 也就是 n + logn = O(n)

// -------------------- 传递的为 len -----------------------------
func MergeSort(a []int, n int) {
	// a = {5, 4, 6, 2, 3, 1}, n = 6
	mergeSort_r(a, 0, n)
	fmt.Println(a)
}
// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func mergeSort_r(a []int, begin int, end int) {
	if end - begin <= 1{ // 数组只有单个元素或者没有元素时返回
		return
	}
	// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
	mid := begin + (end-begin)/2

	// 先将左边排序好
	mergeSort_r(a, begin, mid) // 时间复杂度为 n/2
	// 再将右边排序好
	mergeSort_r(a, mid, end)   // 时间复杂度为 n/2

	merge(a, begin, mid, end) 	// 两个有序数组进行合并 时间复杂度为 n
}
// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := end - begin // 辅助数组的长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
}
// -------------------- 传递的为 len -1-----------------------------
func MergeSortAno(a []int, n int) {
	// a = {5, 4, 6, 2, 3, 1}, n = len(a)-1
	mergeSort_rAno(a, 0, 5)
	fmt.Println(a)
}
// 自顶向下归并排序，排序范围在 [begin,end) 的数组
func mergeSort_rAno(a []int, begin int, end int) {
	if  begin >= end { // 数组只有单个元素或者没有元素时返回
		return
	}

	mid := begin + (end-begin)/2
	// 先将左边排序好
	mergeSort_rAno(a, begin, mid)
	// 再将右边排序好
	mergeSort_rAno(a, mid+1, end)

	mergeAno(a, begin, mid, end) 	// 两个有序数组进行合并
}
// 归并操作
func mergeAno(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := begin
	rightSize := mid + 1
	size := end - begin + 1
	result := make([]int, 0, size)

	for leftSize <= mid && rightSize <= end {
		// 小的元素先放进辅助数组里
		if array[leftSize] < array[rightSize] {
			result = append(result, array[leftSize])
			leftSize++
		} else {
			result = append(result, array[rightSize])
			rightSize++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	if leftSize <= mid {
		result = append(result, array[leftSize:mid+1]...)
	}
	if rightSize <= end {
		result = append(result, array[rightSize:end+1]...)
	}
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < size; i++ {
		array[begin+i] = result[i]
	}
	return
}

// 快速排序
// 原理与 归并排序 相类似, 归并是递的过程比较简单, 归的过程比较繁琐, 快排是递的过程比较繁琐, 归的过程比较简单
// 分区点普遍选用最后一个点(分区点的选择直接影响时间复杂度)
// 1	5	6	2	3	4
//                      ^ (分区间)
// [0,i] < 4 | (i+1, j-1) | [j, len-1] > 4
// 小于4区间   | 未排序区间   | 大于4区间

// 时间复杂度:
//		1) 分区及其平均: 正好分成了大小接近的两个小区间, 快排的递推公式: T(1) = C, T(n) = 2*T(n/2)+n, 与归并排序完全相同, 所以为 nlogn
//		2) 分区及其不平均: 数组原本已经有序了 1 3 5 6 8, 如果我们选择最后一个元素作为 分区点(pivot), 分区大小很不均匀, 我们需要进行大约 n 次操作
// 			才能完成快排的整个过程, 每次扫描 n/2 的元素, 那么时间复杂度就从 nlogn 降到了 n^2

// 空间复杂度与函数调用栈的 max 值有关, 也就是 logn
// 为原地排序
// 非稳定排序: 例: 6 8 7 6 3 5 9 4 , 两个 6 的顺序会发生变化
func QuickSort(a []int, n int) {
	// []int{5, 4, 6, 2, 3, 1}, 5
	quickSort_R(a, 0, n) // 数组长度 - 1
	fmt.Println(a)
}
func quickSort_R(a []int, begin, end int) {
	if begin >= end {
		return
	}
	j := partition(a, begin, end)
	quickSort_R(a, begin, j-1)
	quickSort_R(a, j+1, end)
}
func partition(arr []int, low int, high int) int {
	i, j := low, high-1 // 双指针, 选择最后一个元素作为分区点
	for {
		for arr[i] <= arr[high] {
			if i == high {
				break
			}
			i++
		}
		// <= 处理 1,1,1 情况
		for arr[high] <= arr[j] {
			if j == low {  // 比较在前, -- / ++ 在后, 避免数组越界
				break
			}
			j--
		}
		if i >= j {
			break
		}
		swap(arr, i, j)
	}
	swap(arr, high, i)
	//swap(arr, low, j)
	return i
	//return j
}
func swap(arr []int, a int, b int) {
	if arr[a] == arr[b] {
		return
	}
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}