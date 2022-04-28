package loadgen

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"load/loadGen/lib"
	"load/log"
)

/*
载荷发生器：

QPS：每秒查询量，服务器数据读取
TPS：每秒事务处理量， 服务器数据写入

对于面向终端用户的所有 API，其处理超时时间都是200ms。如果某个 API在承受不高于最高每
秒载荷量 80%的负载情况下造成了处理超时，那么这个 API在性能上就是不合格的

我们在初始化载荷发生器的时候就应该给定上述3个参数，即 「每秒载荷量」、「负载持续时间」、「响应超时时间」。
载荷发生器根据这些参数自行计算出载荷发生以及发送的频率，并控制好并发量。

lps			uint32			每秒载荷量（loads per second）
durationNS	time.Duration	负载持续时间
timeOutNS 	time.Duration	响应超时时间

数组，切片，map 都不是并发安全的，原生的数据类型只有 channel 是并发安全的
*/

// 日志记录器。
var logger = log.DLogger()

// 载荷发生器
// myGenerator 代表载荷发生器的实现类型。
type myGenerator struct {
	resultCh chan *lib.CallResult // 调用结果通道。

	caller lib.Caller // 调用器。

	timeoutNS  time.Duration // 处理超时时间，单位：纳秒。
	lps        uint32        // 每秒载荷量。
	durationNS time.Duration // 负载持续时间，单位：纳秒。

	concurrency uint32        // 载荷并发量：决定 goroutine 的并发数量
	tickets     lib.GoTickets // Goroutine票池。

	// 随时取消载荷发生器
	ctx        context.Context    // 上下文。
	cancelFunc context.CancelFunc // 取消函数。

	callCount int64 // 调用计数。

	status uint32 // 载荷发生器的状态，与 lib/base.go 中状态对应
}

// NewGenerator 会新建一个载荷发生器。
func NewGenerator(pset ParamSet) (lib.Generator, error) {
	logger.Infoln("New a load generator...")

	if err := pset.Check(); err != nil {
		return nil, err
	}
	gen := &myGenerator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.LPS,
		durationNS: pset.DurationNS,
		status:     lib.STATUS_ORIGINAL,
		resultCh:   pset.ResultCh,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}
	return gen, nil
}

// 初始化载荷发生器。
func (gen *myGenerator) init() error {
	var buf bytes.Buffer
	buf.WriteString("Initializing the load generator...")

	// 1e9 表示 一秒对应的纳秒数
	// (1e9 / 1ps)的含义就是：在响应超时时间代表的某一个时间周期内的并发量的最大值。
	// 而最后与之相加的 1 则代表了在某一个时间周期之初，向被测软件发送的那个载荷。
	var total64 = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	if total64 > math.MaxInt32 {
		total64 = math.MaxInt32
	}
	/*
		// concurrency 字段的值代表相关调用过程的并发执行数量。
		// 一个调用过程代表了载荷发生器通过一个载荷与被测软件进行的一次交互。
		// 因此，这一过程的并发执行数量可以比较真实地反映出被测软件的负载程度。

		调用过程的并发执行数量（以下简称并发量）需要根据 timeoutNS（处理超时时间） 字段和 lps（每秒载荷量） 字段的值以及如下公式计算得出：
		并发量 ≈ 单个载荷的响应超时时间 /载荷的发送间隔时间

		性能测试软件要做的就是，先通过若干个预设的限定值模拟出一定程度的负载，然后再以此来测试并得到被测试软件实际能承受的最大负载。
		前面讲到的响应超时时间、每秒载荷发送量和负载持续时间，以及经过计算得出的并发量都属于预设的限定值
	*/
	gen.concurrency = uint32(total64) // 得出并发执行量

	tickets, err := lib.NewGoTickets(gen.concurrency)
	if err != nil {
		return err
	}
	gen.tickets = tickets

	buf.WriteString(fmt.Sprintf("Done. (concurrency=%d)", gen.concurrency))
	logger.Infoln(buf.String())
	return nil
}

// callOne 会向载荷承受方发起一次调用。
func (gen *myGenerator) callOne(rawReq *lib.RawReq) *lib.RawResp {
	atomic.AddInt64(&gen.callCount, 1)
	if rawReq == nil {
		return &lib.RawResp{ID: -1, Err: errors.New("Invalid raw request.")}
	}

	// 这里使用 调用器的 Call 方法，计算请求时间
	start := time.Now().UnixNano()
	resp, err := gen.caller.Call(rawReq.Req, gen.timeoutNS)
	end := time.Now().UnixNano()
	elapsedTime := time.Duration(end - start)

	var rawResp lib.RawResp
	if err != nil {
		errMsg := fmt.Sprintf("Sync Call Error: %s.", err)
		rawResp = lib.RawResp{
			ID:     rawReq.ID,
			Err:    errors.New(errMsg),
			Elapse: elapsedTime}
	} else {
		rawResp = lib.RawResp{
			ID:     rawReq.ID,
			Resp:   resp,
			Elapse: elapsedTime}
	}
	return &rawResp
}

/*
	一个调用过程分为5个操作步骤，即生成载荷、发送载荷并接收响应、检查载荷响应、生成调用结果和发送调用结果。
	前 3 个操作步骤都会由使用方在初始化载荷发生器时传人的那个调用器来完成

	asyncCall 方法在一开始就会启用一个专用的 goroutine，因为对 asyncCall 方法的每一次调用都意味着会有一个专用 goroutine 被启用。
	这里的专用 goroutine 总数会由 goroutine 票池控制，后者由载荷发生器的 tickets 字段代表。
	因此，在该方法中，我们需要在适当的时候对 goroutine 票池中的票进行“获得〞和“归还”操作


	在启用专用 goroutine 之前，从 goroutine 票池获得了一张 goroutine 票。当 goroutine
	票池中已无票可拿时，asyncCall 方法所在的 goroutine 会被阻塞于此。只有存在多余的
	goroutine票时，专用 goroutine 才会被启用，从而当前的调用过程才会执行。另一方面，
	在这个 go 两数的 defer 语句中，会及时地把票归还给 goroutine 票池。这个归还的时机很
	重要，不可或缺，也要恰到好处。
*/
// asyncSend 会异步地调用承受方接口。
func (gen *myGenerator) asyncCall() {
	gen.tickets.Take() // 票池中削减 goroutine 数量

	go func() {
		/*
			为了把调用器可能引发的运行时恐慌转变为错误，需要确保在 asyncCall 方法中的go
			函数的开始处有一条 defer 语句
		*/
		defer func() {
			if p := recover(); p != nil {
				err, ok := interface{}(p).(error) // 断言表达式判断其是否是 error
				var errMsg string
				if ok {
					errMsg = fmt.Sprintf("Async Call Panic! (error: %s)", err)
				} else {
					errMsg = fmt.Sprintf("Async Call Panic! (clue: %#v)", p)
				}
				logger.Errorln(errMsg)
				result := &lib.CallResult{
					ID:   -1,
					Code: lib.RET_CODE_FATAL_CALL,
					Msg:  errMsg}
				gen.sendResult(result)
			}
			gen.tickets.Return()
		}()

		/*
			这里调用 bytes, err := json.Marshal(sreq)，产生的 panic 由 defer 来保底
		*/
		rawReq := gen.caller.BuildReq() // 设置请求内容

		// 调用状态：0-未调用或调用中；1-调用完成；2-调用超时。
		var callStatus uint32

		// 负载发生器接收响应的部分
		// timeoutNS 用来判断软件处理单一载荷处请求是否超时
		timer := time.AfterFunc(gen.timeoutNS, func() {
			/*
				atomic.CompareAndSwapUint32 两数检查并设置 callStatus 变量的值，该两数会返回一个 bool 类型值，用以表示比较并交换是否成功。
				如果未成功，就说明载荷响应接收操作已先完成，忽略超时处理
			*/
			if !atomic.CompareAndSwapUint32(&callStatus, 0, 2) {
				return
			}
			result := &lib.CallResult{
				ID:     rawReq.ID,
				Req:    rawReq,
				Code:   lib.RET_CODE_WARNING_CALL_TIMEOUT,
				Msg:    fmt.Sprintf("Timeout! (expected: < %v)", gen.timeoutNS),
				Elapse: gen.timeoutNS,
			}
			gen.sendResult(result)
		})

		/*
			无论何时，一旦代表调用操作的 cal10ne 方法返回，就要先检查并设置 callstatus 变
			量的值。如果 CAS操作不成功，就说明调用操作已超时，前面传人 time.AfterFunc 函数
			的那个匿名函数已经先执行了，需要忽略后续的响应处理。当然，在响应处理之前先要
			停掉前面启动的定时器。
		*/
		rawResp := gen.callOne(&rawReq)

		if !atomic.CompareAndSwapUint32(&callStatus, 0, 1) {
			return
		}
		timer.Stop()

		// 检查是否存在调用错误
		var result *lib.CallResult
		if rawResp.Err != nil {
			result = &lib.CallResult{
				ID:     rawResp.ID,
				Req:    rawReq,
				Code:   lib.RET_CODE_ERROR_CALL,
				Msg:    rawResp.Err.Error(),
				Elapse: rawResp.Elapse}
		} else {
			result = gen.caller.CheckResp(rawReq, *rawResp)
			result.Elapse = rawResp.Elapse
		}
		gen.sendResult(result)
	}()
}

// sendResult 用于发送调用结果。
func (gen *myGenerator) sendResult(result *lib.CallResult) bool {
	/*
		为了确保调用结果发送的正确性，sendResult 方法必须先检查载荷发生器的状态。
		如果它的状态不是启动的，那就不用执行发送操作了
	*/
	if atomic.LoadUint32(&gen.status) != lib.STATUS_STARTED {
		gen.printIgnoredResult(result, "stopped load generator")
		return false
	}
	select {
	// 调用结果通道已满不行执行结果发送了
	case gen.resultCh <- result:
		return true
	default:
		gen.printIgnoredResult(result, "full result channel")
		return false
	}
}

// printIgnoredResult 打印被忽略的结果。
func (gen *myGenerator) printIgnoredResult(result *lib.CallResult, cause string) {
	resultMsg := fmt.Sprintf("ID=%d, Code=%d, Msg=%s, Elapse=%v", result.ID, result.Code, result.Msg, result.Elapse)
	logger.Warnf("Ignored result: %s. (cause: %s)\n", resultMsg, cause)
}

/*
	gen.ctx 字段的 Err 方法会返回一个 error 类型值，该值会体现“信号”被“发出”的缘由。
	所以，gen.prepareTostop 方法接受这样一个值，并把它作为日志的一部分记录下来，以便做调试和告知之用。
	该方法会先仅在载荷发生器的状态为已启动时，把它变为正在停止状态
*/
// prepareStop 用于为停止载荷发生器做准备。
func (gen *myGenerator) prepareToStop(ctxError error) {
	logger.Infof("Prepare to stop load generator (cause: %s)...", ctxError)

	atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STARTED, lib.STATUS_STOPPING)

	logger.Infof("Closing result channel...")

	close(gen.resultCh)

	atomic.StoreUint32(&gen.status, lib.STATUS_STOPPED)
}

/*
	在进人已启动状态之后，载荷发生器才真正开始生成并发送载荷。包含了载荷发送
	操作和载荷响应接收操作的调用操作是异步执行的，因为只有这样载荷发生器才能做到并发运行
*/
// genLoad 会产生载荷并向承受方发送。
func (gen *myGenerator) genLoad(throttle <-chan time.Time) {
	for {
		select {
		case <-gen.ctx.Done():
			gen.prepareToStop(gen.ctx.Err())
			return
		default:
		}

		gen.asyncCall() // 进行异步调用

		// 这里 case： <-gen.ctx.Done() 重复一遍的原因是：
		// 由于 select 语句在多个满足条件的 case 之间做伪随机选择时的不确定性，当节流网的到期通知和上下文的“信号”同时到达时，
		// 后者代表的 case 不一定会被选中。这也是为了保险起见，以使载荷发生器总能及时地停止
		if gen.lps > 0 {
			select {
			case <-throttle: // 周期性的长短由节阀流控制
			case <-gen.ctx.Done():
				gen.prepareToStop(gen.ctx.Err())
				return
			}
		}
	}
}

// Start 会启动载荷发生器。
func (gen *myGenerator) Start() bool {
	logger.Infoln("Starting load generator...")

	// 检查是否具备可启动的状态，顺便设置状态为正在启动
	if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_ORIGINAL, lib.STATUS_STARTING) {
		if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STOPPED, lib.STATUS_STARTING) {
			return false
		}
	}

	// 设定节流阀。
	var throttle <-chan time.Time
	if gen.lps > 0 {
		interval := time.Duration(1e9 / gen.lps)
		logger.Infof("Setting throttle (%v)...", interval)
		throttle = time.Tick(interval)
	}

	// 初始化上下文和取消函数。运行一段时间自己停止下来
	gen.ctx, gen.cancelFunc = context.WithTimeout(context.Background(), gen.durationNS)

	// 初始化调用计数。
	gen.callCount = 0

	/*
		注意，这里在改变状态时使用了原子操作。简单来说，原子操作就是一定会一次做
		完的操作。在操作过程中不允许任何中断，操作所在的 goroutine 和内核线程也绝不会被
		切换下 CPU。这是由 CPU、操作系统和 Go 运行时系统多级保证的。当然，从根本上讲，
		还是 CPU 的原语在起作用。Go标准库中的 atomic 包提供了很多原子操作方法。
	*/
	// 设置状态为已启动。
	atomic.StoreUint32(&gen.status, lib.STATUS_STARTED)

	// 启动一个 goroutine 来执行并发送载荷的流程，所以 Start 是非阻塞的
	go func() {
		// 生成 并 发送载荷。
		logger.Infoln("Generating loads...")
		gen.genLoad(throttle)
		logger.Infof("Stopped. (call count: %d)", gen.callCount)
	}()
	return true
}

/*
	手动停止 载荷发生器 的方法
*/
func (gen *myGenerator) Stop() bool {
	if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STARTED, lib.STATUS_STOPPING) {
		return false
	}

	gen.cancelFunc()

	for {
		// 如果状态为 STATUS_STOPPED 说明 prepareToStop 执行结束
		if atomic.LoadUint32(&gen.status) == lib.STATUS_STOPPED {
			break
		}
		time.Sleep(time.Microsecond)
	}
	return true
}

func (gen *myGenerator) Status() uint32 {
	return atomic.LoadUint32(&gen.status)
}

func (gen *myGenerator) CallCount() int64 {
	return atomic.LoadInt64(&gen.callCount)
}
