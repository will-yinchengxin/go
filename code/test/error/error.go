package err

import (
	"database/sql"
	"github.com/pkg/errors"
	"fmt"
)

/*
当前 error 的问题有两点:
	1) 无法 wrap 更多的信息，比如调用栈，比如层层封装的 error 消息
	2) 无法很好的处理类型信息，比如我想知道错误是 io 类型的，还是 net 类型的

	go的痛点不是缺少泛型, 不是 error 功能太差, 是 GC 太弱了,尤其对大内存不是很友好
*/


/*
问题一:
	解决问题一: 使用 https://github.com/pkg/errors
*/
//func foo() error {
//	return errors.Wrap(sql.ErrNoRows, "foo failed")
//}
//
//func bar() error {
//	return errors.WithMessage(foo(), "bar failed")
//}
//
//func TestError() {
//	err := bar()
//	fmt.Println(err, "\r\n")
//	fmt.Println(errors.Cause(err), "\r\n")
//	if errors.Cause(err) == sql.ErrNoRows {
//		fmt.Println(12)
//	}
//
//	if errors.Cause(err) == sql.ErrNoRows {
//		fmt.Printf("data not found, %v\n", err)
//		fmt.Printf("%+v\n", err)
//		return
//	}
//	if err != nil {
//		// unknown error
//	}
//}


/*
问题二:
	func bar() error {
	   if err := foo(); err != nil {
		  return fmt.Errorf("bar failed: %w", foo())
	   }
	   return nil
	}

	func foo() error {
	   return fmt.Errorf("foo failed: %w", sql.ErrNoRows)
	}

	多层 error 嵌套, 更改了error的原始信息, 就不能准确获取error的预定义信息,

解决办法:
	1) 1.13 版本新增的 Is 和 As
	2) 或者使用 https://github.com/pkg/errors 库提供的 errors.Cause
		switch err := errors.Cause(err).(type) {
			case *MyError:
					// handle specifically
			default:
					// unknown error
		}
*/
func bar() error {
	if err := foo(); err != nil {
		return fmt.Errorf("bar failed: %w", foo())
	}
	return nil
}

func foo() error {
	return fmt.Errorf("foo failed: %w", sql.ErrNoRows)
}

func TestError() {
	e := errors.New("this is a error")
	w := fmt.Errorf("more info about it %w", e)
	fmt.Println(w) // more info about it this is a error

	err := bar()

	// 这里个小问题，Is 是做的指针地址判断，如果错误 Error() 内容一样，但是根 error 是不同实例，那么 Is 判断也是 false, 这点就很扯
	//if errors.Is(err, errors.New("sql: no rows in result set")) {  // 这里因为被判断为false, 原因是 error 是不同实例
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("err is,  %+v\n", err)
		return
	}
	if errors.As(err, sql.ErrNoRows) {
		fmt.Printf("err as,  %+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}

/*
其中 github.com/pkg/errors 库 Cause 实现:

type withStack struct {
 error
 *stack
}

func (w *withStack) Cause() error { return w.error }

func Cause(err error) error {
 type causer interface {
  Cause() error
 }

 for err != nil {
  cause, ok := err.(causer)
  if !ok {
   break
  }
  err = cause.Cause()
 }
 return err
}

Cause 递归调用，如果没有实现 causer 接口，那么就返回这个 err
*/

/*
- 源码很多地方写 panic, 但是工程实践，尤其业务代码不要主动写 panic
- 理论上 panic 只存在于 server 启动阶段，比如 config 文件解析失败，端口监听失败等等， !!!!! 所有业务逻辑禁止主动 panic // error 与 panic
- 所有异步的 goroutine 都要用 recover 去兜底处理 // error 与 panic
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				log.Errorc(ctx, "[dao] ping panic: %v, stack: %v", r, string(debug.Stack()))
			}
		}()

        ... code ...
	}()
- 数据传输和退出控制，需要用单独的 channel 不能混, 我们一般用 context 取消异步 goroutine, 而不是直接 close channels // 错误处理与资源释放
- error 级联使用问题
	type myError struct {
	 string
	}

	func (i *myError) Error() string {
	 return i.string
	}

	func Call1() error {
	 return nil
	}

	func Call2() *myError {
	 return nil
	}

	func main() {
	 err := Call1()
	 if err != nil {
	  fmt.Printf("call1 is not nil: %v\n", err)
	 }

	 err = Call2()
	 if err != nil {
	  fmt.Printf("call2 err is not nil: %v\n", err)
	 }
	}
	// 报出错误: error is not nil, but value is nil
	解决方法就是 Call2 err 重新定义一个变量，当然最简单就是统一 error 类型, 当然统一 error 类型这种解决方案有点难

- 不要并发对 error 赋值(error 接口不是并发安全的)
- err 是否可以忽略: 不能保证以后还是兼容的逻辑，一定要处理 error，至少要打日志
- err.Wrap(err, "faild"), 如果 err 为 nil, 也会返回nil, 所以wrap前最好判断一下
*/

// 参考连接: https://mp.weixin.qq.com/s/XojOIIZfKm_wXul9eSU1tQ