package map_set

import (
	"fmt"
	"sort"
)

/*
map 类型的零值是nil。map 值的长度表示了其中的键-元素对的数量，其零值的长度总是0。

ipSwitches["192.168.0.1"] = true
delete（ipSwitches，"192.168.0.1"）
无论ipSwitches中是否存在键“192.168.0.1”，内建函数delete都会默默地执行完毕。
*/

// 利用map实现一个不重复的set集合
func mapSet()  {
	hashSet := make(map[string]struct{})
	data := []string{"Hello", "World", "213", "3213", "213", "World"}
	for _, datum := range data {
		hashSet[datum] = struct{}{}
	}
	for s, _ := range hashSet {
		fmt.Sprintf(s)
	}
}

// 对 map 进行排序, 如果值不同, 对值排序, 如果值相同, 对键进行排序
func SortMap() {
	wordFrequency := map[string]int{"banana": 3, "america": 2, "abb": 2, "test": 2, "car": 1, "did" : 7}

	vec := mapToSlice(wordFrequency)
	fmt.Println( vec)
	sort.Slice(vec, func(i, j int) bool {
		// 1. value is different - sort by value (in reverse order)
		if vec[i].value != vec[j].value {
			return vec[i].value > vec[j].value
		}
		// 2. only when value is the same - sort by key
		return vec[i].key < vec[j].key // 正序
		return vec[i].key > vec[j].key // 倒序
	})

	fmt.Printf("%v", vec)
}

func mapToSlice(in map[string]int) []KV {
	vec := make([]KV, len(in))
	i := 0
	for k, v := range in {
		vec[i].key = k
		vec[i].value = v
		i++
	}
	return vec
}

type KV struct {
	key   string
	value int
}