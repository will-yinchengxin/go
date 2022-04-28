package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Pool struct {
	pool *redis.Pool
}

func NewPool() *Pool {
	return &Pool{
		&redis.Pool{
			MaxIdle: 16, // 最初的连接数量
			MaxActive: 0, // 最大的连接数量, 不确定给0(0表示自动定义),按需分配
			IdleTimeout: 300, // 连接关闭时间 300 秒(300 不使用自动关闭)
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "127.0.0.1:16379")
			},
		},
	}
}

func TestPool() {
	rs := NewPool()
	c := rs.pool.Get()
	defer rs.pool.Close()
	defer c.Close()

	_, err := c.Do("Set", "Name", "will")
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := redis.String(c.Do("Get", "Name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}