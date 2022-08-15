package limit

import (
	"fmt"
	"time"
)

func main() {
	bucket := &BucketLimit{
		cap:            10,
		interval:       time.Second,
		lastAccessTime: time.Now(),
	}
	bucket.AccessBucket()
}

type BucketLimit struct {
	cap            int           //漏桶的容量
	interval       time.Duration // 允许访问的时间间隔 1 s
	recCount       int           // 最近访问的次数
	lastAccessTime time.Time
}

func (b *BucketLimit) AccessBucket() bool {
	now := time.Now()
	pastTime := now.Sub(b.lastAccessTime)
	leaks := int(float64(pastTime) / float64(b.interval))
	fmt.Println(leaks)
	if leaks > 0 {
		if leaks > b.recCount {
			b.recCount = 0
		} else {
			b.recCount -= leaks
		}
	}
	if b.recCount > b.cap {
		return false
	}
	// 超过了五分钟，全部至零
	if pastTime > time.Second*5 {
		b.recCount = 0
	} else {
		b.recCount++
	}
	b.lastAccessTime = time.Now()
	return true
}
// help page: https://www.cnblogs.com/failymao/p/15228406.html
