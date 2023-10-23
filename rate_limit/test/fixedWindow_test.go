package test

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

// 固定窗口限流
func TestFixedWindow(t *testing.T) {
	fixedWindowRateLimiting()
}

func fixedWindowRateLimiting() {
	// 允许每秒100次
	limiter := rate.NewLimiter(rate.Limit(100), 1)
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}

func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond)
}
