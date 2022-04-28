package loadgen

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"load/loadGen/lib"
)

/*
NewCenerator 函数的参数声明列表与之前展示的不太一样，这里把所有参数都
内置到了一个名为 Paramset 的结构体类型中。如此一来，在需要变动 NewGenerator 两数的
参数时，就无需改变它的声明了，变动只会影响 ParamSet 类型。另外，还为 Paramset 类
型添加了一个名为 Check 的指针方法，它会检查当前值中所有字段的有效性。一旦发现无
效字段，它就会返回一个非 nil 的 error 类型值。这样做使得 NewCenerator 两数的调用方
可以先行检查待传入的参数集合的有效性。当然，无论怎样，在 NewGenerator 函数内部还
是需要调用这个方法。Paramset 类型及其方法的声明就不展示了，你可以在 loadgen 子包
*/
type ParamSet struct {
	Caller     lib.Caller           // 调用器。
	TimeoutNS  time.Duration        // 响应超时时间，单位：纳秒。
	LPS        uint32               // 每秒载荷量。
	DurationNS time.Duration        // 负载持续时间，单位：纳秒。
	ResultCh   chan *lib.CallResult // 调用结果通道。
}

// Check 会检查当前值的所有字段的有效性。
// 若存在无效字段则返回值非nil。
func (pset *ParamSet) Check() error {
	var errMsgs []string

	if pset.Caller == nil {
		errMsgs = append(errMsgs, "Invalid caller!")
	}
	if pset.TimeoutNS == 0 {
		errMsgs = append(errMsgs, "Invalid timeoutNS!")
	}
	if pset.LPS == 0 {
		errMsgs = append(errMsgs, "Invalid lps(load per second)!")
	}
	if pset.DurationNS == 0 {
		errMsgs = append(errMsgs, "Invalid durationNS!")
	}
	if pset.ResultCh == nil {
		errMsgs = append(errMsgs, "Invalid result channel!")
	}
	var buf bytes.Buffer
	buf.WriteString("Checking the parameters...")
	if errMsgs != nil {
		errMsg := strings.Join(errMsgs, " ")
		buf.WriteString(fmt.Sprintf("NOT passed! (%s)", errMsg))
		logger.Infoln(buf.String())
		return errors.New(errMsg)
	}
	buf.WriteString(
		fmt.Sprintf("Passed. (timeoutNS=%s, lps=%d, durationNS=%s)",
			pset.TimeoutNS, pset.LPS, pset.DurationNS))
	logger.Infoln(buf.String())
	return nil
}
