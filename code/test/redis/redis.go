package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

/*
   127.0.0.1:6379> set name 123
   OK
   127.0.0.1:6379> get name
   "123"
   127.0.0.1:6379> eval "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0  end" 1 name 123
   (integer) 1
   127.0.0.1:6379> get name
   (nil)
*/
/*
	redis := RS.InitRedis()
	lockKey := "redisLock"
	clientID := uuid.NewV4().String()
	ok := redis.SetWitLock(lockKey, clientID, 10)
	if !ok {
		return
	}

	script := "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0  end"
	s := LUA.NewScript(1, script)
	val, err := s.Do(redis.RS, lockKey, clientID)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Println(val)
**/
type RedisResoce struct {
	RS redis.Conn
}

func InitRedis() *RedisResoce {
	host := "172.16.252.99:6379"

	rs, err := redis.Dial("tcp", host)
	//defer rs.Close()

	if err != nil {
		log.Println(err)
		return &RedisResoce{}
	}
	return &RedisResoce{
		rs,
	}
}

func (rs *RedisResoce) SetWitLock(key string, val string, time int) bool { //SET test 1 EX 10 NX
	intVal, err := rs.RS.Do("set", key, val, "EX", time, "NX")
	if err != nil {
		log.Fatal(err)
		return false
	}
	if intVal != nil {
		return true
	}
	return false
}

func (rs *RedisResoce) StringSetAndGet() {
	_, err := rs.RS.Do("Set", "test", 100)
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	val, err := redis.Int64(rs.RS.Do("Get", "test"))
	if err != nil {
		fmt.Println("get test val failed,", err)
		return
	}
	fmt.Println(val)

	_, err = rs.RS.Do("MSet", "b", 10, "c", "test")
	if err != nil {
		fmt.Println("mset val failed,", err)
	}
	//value, err := redis.Strings(rs.RS.Do("MGet", "b", "c"))
	//value, err := redis.Int64s(rs.RS.Do("MGet", "b", "c"))
	//if err != nil {
	//	fmt.Println("mget val failed,", err)
	//}
	//for _, v := range value {
	//	fmt.Println(v)
	//}

}

func (rs *RedisResoce) Expire() {
	_, err := rs.RS.Do("expire", "yin", 10)
	if err != nil {
		fmt.Println("expire err, ", err)
		return
	}
}

func (rs *RedisResoce) SETNX() {
	int, err := redis.Int(rs.RS.Do("setnx", "fadsf", 10))
	if err != nil {
		fmt.Println("expire err, ", err)
		return
	}
	fmt.Println(int > 0)
}

func (rs *RedisResoce) List() {
	_, err := rs.RS.Do("lpush", "testList", "test1", "test2", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	//val, err := redis.String(rs.RS.Do("lpop", "testList"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(val)

	vall, err := redis.Strings(rs.RS.Do("lrange", "testList", 0, -1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vall)
}

func (rs *RedisResoce) Hash() {
	_, err := rs.RS.Do("Hset", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := redis.Int(rs.RS.Do("Hget", "books", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func SendAndFlush() {
	c := InitRedis()
	c.RS.Send("SET", "foo", "bar")
	c.RS.Send("SET", "dai", "bar")
	c.RS.Send("SET", "fa", "bar")
	c.RS.Send("GET", "foo")
	//c.RS.Send("GET", "dai")
	c.RS.Flush()              // Flush 将连接的输出缓冲区刷新到服务器
	v1, err := c.RS.Receive() // 从服务器读取单个回复
	v2, err := c.RS.Receive()
	v3, err := c.RS.Receive()
	v4, err := c.RS.Receive()
	v5, err := c.RS.Receive() // Receive 超出数量会陷入等待
	if err != nil {
		fmt.Println(v1, v2, v3, v4, v5)
	}
}

func TestRedis() {
	//rd, _ := initRedis()
	// 操作redis时调用Do方法，第一个参数传入操作名称（字符串），然后根据不同操作传入key、value、数字等
	// 返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	//rd.Do("SET", "name", "will")

	// send 结合 flush 可以批量执行多个命令
	//rd.Send("SET", "age", 18)
	//rd.Send("SET", "sex", "man")
	//rd.Flush()

	/*事务执行*/
	//rd.Send("MULTI")
	//rd.Send("INCR", "foo")
	//rd.Send("INCR", "bar")
	//r, _ := rd.Do("EXEC")
	//fmt.Println(r) // prints [1, 1]

	// Exists
	// fmt.Println(redis.Bool(rd.Do("EXISTS", "name")))

	// mget
	//reply, _ := redis.Values(rd.Do("MGET", "name", "sex"))
	//var name string
	//var sex string
	//redis.Scan(reply, &name, &sex)
	//fmt.Println(name, sex)
}
