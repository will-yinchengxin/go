package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

// 生成随机token
func main() {
	tmp := make([]byte, 16)
	io.ReadFull(rand.Reader, tmp)
	fmt.Println(base64.RawURLEncoding.EncodeToString(tmp))
}
