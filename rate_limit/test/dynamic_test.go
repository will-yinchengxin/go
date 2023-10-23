package test

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 动态速率限制
func TestDynamic(t *testing.T) {
	dynamicRateLimiting()
}

func dynamicRateLimiting() {
	// 每秒100次
	limiter := rate.NewLimiter(rate.Limit(100), 1)

	// 每10秒调整 limiter, 将 limiter 提升到每秒 200
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("---adjust limiter---")
		limiter.SetLimit(rate.Limit(200))
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}
