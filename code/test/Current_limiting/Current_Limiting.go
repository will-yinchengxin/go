package Current_limiting

import (
	"github.com/juju/ratelimit"
	"fmt"
	"time"
)
// 漏斗/令牌桶作为限流常用的工具,开源界流行的限流器大多数都是基于令牌桶实现的
// github.com/juju/ratelimit 提供了几种不同特色的令牌桶方式

func GetCurrentLimit() {
	// 参数一代表时间间隔, 参数二代表容量, 参数三代表放入频率
	//b := ratelimit.NewBucketWithQuantum(time.Millisecond, 10, 1)
	b := ratelimit.NewBucketWithQuantum(time.Second, 10, 1)
	/*
	before := b.Available()
	tokenGet := b.TakeAvailable(1)
	if tokenGet != 0 {
		fmt.Println("获取令牌桶 index = ", i+1, "前后数量-> 前: ", before, " 后: ", b.Available(), ", tokenGet = ", tokenGet)
	} else {
		fmt.Println("未获取到令牌桶, 拒绝 ", i+1)
	}

	这里可以通过判断返回值 为 true or false/或者根据ip的不同来 判断允不允许客户端请求
	*/
	for i := 0; i < 100; i++ {
		before := b.Available()
		tokenGet := b.TakeAvailable(1)
		if tokenGet != 0 {
			fmt.Println("获取令牌桶 index = ", i+1, "前后数量-> 前: ", before, " 后: ", b.Available(), ", tokenGet = ", tokenGet)
		} else {
			fmt.Println("未获取到令牌桶, 拒绝 ", i+1)
		}
		time.Sleep(time.Millisecond)
	}
}
