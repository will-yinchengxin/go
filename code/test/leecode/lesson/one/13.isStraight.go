package one

/*
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。
A 不能视为 14。

示例 1:
	输入: [1,2,3,4,5]
	输出: True

示例 2:
	输入: [0,0,1,2,5]
	输出: True

限制：
	数组长度为 5
	数组的数取值为 [0, 13] .


0重复可以 别的不行, 0是赖子
*/
func IsStraight(nums []int) bool {
	/*
		1) 是零就跳过
		2) 看最大值减去最小值是不是小于5
	*/
	//sort.Ints(nums) 不使用 sort
	if len(nums) <= 0 {
		return false
	}
	tagMap := make(map[int]struct{})
	minNum := 0
	maxNum := 14

	for i := 0; i < 5; i++ {
		// 癞子跳过
		if nums[i] == 0 {
			continue
		}
		// 重复直接 false
		if _, ok := tagMap[nums[i]]; ok {
			return false
		}
		if nums[i] < maxNum {
			maxNum = nums[i]
		}
		if nums[i] > minNum {
			minNum = nums[i]
		}
		tagMap[nums[i]] = struct{}{}
	}
	if minNum - maxNum > 0 {
		return minNum - maxNum < 5
	}
	return maxNum - minNum < 5
}
