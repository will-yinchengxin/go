package z1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)

// 0 开头为八进制
// 0x 或 0X 开头为十六进制
// syscall/ztypes_openbsd_arm64.go
const (
	S_IRUSR = 0x100
	S_IWUSR = 0x80
	S_IXUSR = 0x40

	S_IROTH = 0x4 //其它可读
	S_IWOTH = 0x2 //其它可写
	S_IXOTH = 0x1 //其它可执行
)

func TestFMT(t *testing.T) {
	name := struct {
		Name string
	}{
		"will",
	}
	t.Logf("%v", name)  // {will}
	t.Logf("%+v", name) // {Name:will}

	//%b	表示为二进制
	//%d	表示为十进制

	// 宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。
	// 精度通过（可选的）宽度后跟点号后跟的十进制数指定。
	// 如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0
	/*
		%f:    默认宽度，默认精度
		%9f    宽度9，默认精度
		%.2f   默认宽度，精度2
		%9.2f  宽度9，精度2
	*/
	// 对于整数，宽度和精度都设置输出总长度。采用精度时表示右对齐并用0填充，而宽度默认表示用空格填充。
	t.Logf("%9b %3d %s", S_IRUSR, S_IRUSR, "用户读") // 100000000 256 用户读

	// [-][rwx][r-x][r--]
	// 解释 r w x
	t.Logf("%9b %3d %s", S_IROTH, S_IROTH, "其它读")
	t.Logf("%9b %3d %s", S_IWOTH, S_IWOTH, "其它写")
	t.Logf("%9b %3d %s", S_IXOTH, S_IXOTH, "其它可执行")
	/*
		100   4 其它读
		 10   2 其它写
		  1   1 其它可执行

		每个权限位标志符都是 1，也就理解 4 2 1 的由来
	*/
}

// save mutilpart month
func TestMonth(t *testing.T) {
	t.Log(MonthFormatToInt("1,2,3")) // 14
	t.Log(MonthFormatToInt("2,3"))   // 12
	t.Log(MonthFormatToInt("1"))
	t.Log(MonthFormatToInt("2"))
	t.Log(MonthFormatToInt("3"))
	t.Log(MonthFormatToInt("4"))
	t.Log(MonthFormatToInt("5"))
	t.Log(MonthFormatToInt("6"))
	t.Log(MonthFormatToInt("7"))
	t.Log(MonthFormatToInt("8"))
	t.Log(MonthFormatToInt("9"))
	t.Log(MonthFormatToInt("10"))
	t.Log(MonthFormatToInt("11"))
	t.Log(MonthFormatToInt("12"))
}

func MonthFormatToInt(month string) int {
	/*
		1000000000000 4096  12月
		 100000000000 2048  11月
		  10000000000 1024  10月
		   1000000000 512   9 月
			100000000 256   8 月
			 10000000 128   7 月
			  1000000  64   6 月
			   100000  32   5 月
				10000  16   4 月
				 1000   8   3 月
				  100   4   2 月
				   10   2   1 月
	*/
	var Month int = 0
	monthList := strings.Split(month, ",")
	for _, m := range monthList {
		e, _ := strconv.ParseFloat(m, 64)
		fmt.Println(Month, int(math.Pow(2, e))) // 2^n
		Month = Month | int(math.Pow(2, e))
	}
	return Month
}

func TestFormatToList(t *testing.T) {
	Month := 12
	monthList := make([]string, 0)
	for m := 1; m <= 12; m++ {
		if Month&int(math.Pow(2, float64(m))) > 0 {
			monthList = append(monthList, strconv.Itoa(m))
		}
	}
	t.Log(monthList)
}

// 当前时间是否在维护时间段内
func TestIsOnDuty(t *testing.T) {
	t.Log(IsOnDuty("9:00", "12:00"))
	t.Log(IsOnDuty("19:00", "12:00"))
	t.Log(IsOnDuty("14:00", "18:00"))
}

func IsOnDuty(StartTime, EndTime string) bool {
	now := time.Now().Format("15:04")
	return (StartTime <= EndTime && StartTime <= now && EndTime >= now) || // 不跨 00:00
		(StartTime > EndTime && (StartTime <= now || now <= EndTime)) // 跨 00:00
}
