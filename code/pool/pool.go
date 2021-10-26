package pool

import "sync"

type user struct {
	Name string
	Email string
}

func (u *user) reset(name string, email string)  {
	u.Name = name
	u.Email = email
}

// sync.pool 的一般使用规则
func Pool() {
	pool := sync.Pool{
		New: func() interface{} {
			return &user{}
		},
	}
	u := pool.Get().(*user)
	// 将 u 归还
	defer pool.Put(u)

	//重置u
	u.reset("tom", "2132423@qq.com")
}
