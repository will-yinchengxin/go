package rand

import (
	"fmt"
	"math/rand"
	"time"
)

// Code 生成6位数随机码int
func GetRoundString() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机字串介于0，1000000之间，取六位
	fmt.Println(rnd.Int31n(1000000))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(code)
	return code
}
