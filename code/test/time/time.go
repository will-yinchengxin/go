package time

import (
	"fmt"
	"time"
)

var Name string

func Time() {
	now := time.Now().Local()
	// 获取当前时间
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2021-08-27 15:14:49
	// 获取指定间隔后的时间
	//now.Add(time.Hour * 2)
	deadTime := time.Date(now.Year(), now.Month(), now.Day(), now.Add(time.Hour * 1).Hour(), now.Minute(), now.Second(), 0, now.Location()).Format("2006-01-02 15:04:05")
	fmt.Println(deadTime)

	// 时间戳
	timestamp1 := now.Unix()     // 毫秒时间戳 1630048018
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Println(timestamp1, timestamp2)

	// 将时间戳转为时间格式
	timer := time.Unix(1630048018, 0)
	timerNow := time.Unix(1630048018, 0).Format("2006-01-02 15:04:05") // 2021-08-27 15:06:58
	fmt.Println(timerNow)
	fmt.Println(timer.Year(), timer.Month(), timer.Day(), timer.Hour(), timer.Second(), timer.Minute())

	// 定时器
	//for i := range time.Tick(time.Second) {
	//	fmt.Println(i)
	//}
}

// 求程序执行的时间
func TimeCause() {
	// 计算程序执行的时间
	timeStart := time.Now()
	fmt.Println("this is time test")
	time.Sleep(time.Second)
	timeSpend := time.Since(timeStart)
	fmt.Println(timeSpend)
}

// 求规定时常内, 有多少个时间间隔, 加入时间间隔为1秒
func HowManyTick() int64 {
	startTime := time.Now()
	time.Sleep(time.Second)
	return int64(time.Now().Sub(startTime) / time.Millisecond)
}

// 计算两个时间间的时间间隔 时间格式形如: 2006/01/02
func GetTimeArr(start, end string) int64{
	timeLayout  := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout,  start,  loc)
	endUnix, _ := time.ParseInLocation(timeLayout,  end,  loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date :=	(endTime - startTime) / 86400
	return date
}

// 计算两个时间戳之间的所有时间戳
// 请求: time.StatisticsDate(1630290706, 1636945034)
// [1636905600 1636819200 1636732800 1636646400 1636560000]
func StatisticsDate(startTime int64, endTime int64) (dates []int64) {
	dates = []int64{}
	tEnd := time.Unix(endTime, 0)
	tEnd = time.Date(tEnd.Year(), tEnd.Month(), tEnd.Day(), 0, 0, 0, 0, time.Local)

	tStart := time.Unix(startTime, 0)
	tStart = time.Date(tStart.Year(), tStart.Month(), tStart.Day(), 0, 0, 0, 0, time.Local)

	for {
		dates = append(dates, tEnd.Unix())
		tEnd = tEnd.AddDate(0, 0, -1)
		if tEnd.Unix() < tStart.Unix() {
			break
		}

	}
	return
}

// 获取当天零点的时间戳
func TodayZero()  {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber)
}

// 获取指定时间戳 0 点和 23:59 时间戳
func ParseDate(timeInt int64) (start, end int64) {
	timerNow := time.Unix(timeInt, 0).Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timerNow)
	timeStart := t.Unix()
	timeEnd := timeStart + 86400 - 1
	return timeStart, timeEnd
}