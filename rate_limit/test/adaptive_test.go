package test

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 自适应调整限流策略
func TestAdaptive(t *testing.T) {
	adaptiveRateLimiting()
}

func adaptiveRateLimiting() {
	// 每秒十次
	limiter := rate.NewLimiter(rate.Limit(10), 1)

	// 自适应调整
	go func() {
		for {
			time.Sleep(time.Second * 10)
			// 测量之前请求的响应时间
			responseTime := measureResponseTime()
			if responseTime > 500*time.Millisecond {
				fmt.Println("---adjust limiter 50---")
				limiter.SetLimit(rate.Limit(50))
			} else {
				fmt.Println("---adjust limiter 100---")
				limiter.SetLimit(rate.Limit(100))
			}
		}
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go func() {
			process()
		}()
	}
}

// 测量以前请求的响应时间
// 执行自己的逻辑来测量响应时间
func measureResponseTime() time.Duration {
	return time.Millisecond * 100
}
