package context

/*
context 主要使用场景
	1) 传递超时信息
	2) 传递信号, 用于消息通知
	3) 传递数据
*/

// --------------------etcd 例子-------------------------------
//func Watch(ctx context.Context, revision int64) {
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel
//	for {
//		rch := watcher.Watch(ctx, watchPath, clientv3.WithRev(revision))
//		for wresp := range rch {
//			// ..... do something
//		}
//
//		select {
//		case <-ctx.Done():
//			// server closed, return
//		default:
//
//		}
//	}
//}

/*
1) 除了框架层不要使用 WithValue 携带业务数据，这个类型是 interface{}, 编译期无法确定，运行时 assert 有开销。如果真要携带也要用 thread-safe 的数据
2) 一定不要打印 Context, 尤其是从 http 标准库派生出来的，谁知道里面存了什么
3) Context 通常做为第一个参数传给函数，但如果 Context 生命周期等同于结构体，当成结构体成员也可以
4) 尽可能不要自定义用户层 Context, 除非收益巨大
5) 异步 goroutine 逻辑使用 Context 时要清楚谁还持有，会不会提前超时，尤其调 rpc, db, redis 时
6) 派生出来的 child ctx 一定要配合 defer cancel() 使用，释放资源
*/