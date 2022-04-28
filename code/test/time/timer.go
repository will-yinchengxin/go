package time

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// 使用timer的几种方式
// 1) NewTimer/timer.After/timer.AfterFun
func TestTimer() {
	// 程序这里会死锁, Stop() 操作并没将 channel C 关闭
	//timerOne := time.NewTimer(2 * time.Second)
	//go func() {
	//	timerOne.Stop()
	//}()
	//<- timerOne.C
	//print("done")

	timeTwo := time.NewTimer(2 * time.Second)
	go func() {
		if !timeTwo.Stop() {
			<-timeTwo.C
		}
	}()

	select {
	case <-timeTwo.C:
		print("expired")
	default:
	}
	print("done")
}

func Tick() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

// 设置定时器执行任务
type TestJobs struct {
}

func (j *TestJobs) start(ctx context.Context) {
	t := time.NewTicker(30) //每30秒执行一次
	for {
		select {
		case <-ctx.Done(): //监听进程
			// 日志内容记录或者打点操作
			logrus.Fatal("newTimerError")
			//关闭定时器
			t.Stop()
			return
		case <-t.C: // 执行任务
			j.execute(ctx, t)
		}
	}
}

func (j *TestJobs) execute(ctx context.Context, t *time.Ticker) {
	fmt.Println(time.Now().String())
	var err error
	defer func() {
		// 日志内容记录或者打点操作
		logrus.Fatal("newTimerError")
	}()

	// 这里可以执行定时业务的逻辑代码
	//err = dao.NewFirmwareDao().WithContext(ctx).TimePublish()
	if err != nil {
		return
	}
}