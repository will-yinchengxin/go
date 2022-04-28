package project

import (
	"github.com/garyburd/redigo/redis"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

// 使用Redis实现积分排行榜，并支持同积分按时间排序
/*
使用zset(浮点数存储):
	1) score 为整数位
	2) 结束时间 - 当前时间 作为小数位
	3) 1) + 2) 的浮点数作为 score 的最终值
	4) 使用命令:
		zrevrangebyscore list 1 100 // 获取倒叙排名
		zscore key member // 某个用户的值
		zcard key  // 获取参与排名的总人数
		zrevrank key member // 获取某个用户的当前排名
*/

type RedisResoce struct {
	RS redis.Conn
}

func InitRedis() *RedisResoce {
	host := "127.0.0.1:16379"

	rs, err := redis.Dial("tcp", host)
	//defer rs.Close()

	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
	}
	//redisDb := 0
	//rs.Do("SELECT", redisDb)

	return &RedisResoce{
		rs,
	}
}

func Integral() {
	rs := InitRedis()
	rs.RS.Do("")
}