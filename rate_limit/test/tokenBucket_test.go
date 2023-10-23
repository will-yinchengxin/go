package test

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 令牌桶算法
func TestTokenBucket(t *testing.T) {
	tokenBucketRateLimiting()
}

func tokenBucketRateLimiting() {
	// 比漏斗算法优势在用可以应对突发流量

	//  每秒10个请求, 5个请求可以被放行
	limiter := rate.NewLimiter(rate.Limit(10), 5)
	ctx, _ := context.WithTimeout(context.TODO(), time.Millisecond)
	for i := 0; i < 200; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}
