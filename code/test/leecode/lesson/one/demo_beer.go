package one

/*
现有x瓶啤酒，每3个空瓶子换一瓶啤酒，每7个瓶盖子也可以换一瓶啤酒，问最后可以喝多少瓶啤酒。
*/
func Beer(beerCount int) {
	// count 啤酒 k1 盖子 k2 空瓶
	count, k1, k2 := beerCount, beerCount, beerCount
	for {
		if k1 >= 3 {
			k1 = k1 - 3

			count++
			k1++
			k2++
		}
		if k2 >= 7 {
			k2 = k2 - 7

			count++
			k1++
			k2++
		}
		if k1 < 3 && k2 < 7 {
			break
		}
	}
	print(count)
}