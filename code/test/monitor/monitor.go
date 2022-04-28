package monitor

import (
	"github.com/shirou/gopsutil/process"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func Monitor()  {
	// 创建进程对象
	p, _ := process.NewProcess(int32(os.Getpid()))
	fmt.Println(os.Getpid(), p) // 12904 {"pid":12904}

	// 进程cpu使用率
	cpuPercent, _ := p.Percent(time.Second) // 进程的CPU使用率需要通过计算指定时间内的进程的CPU使用时间变化计算出来
	fmt.Println(cpuPercent)
	// 占单个核心的比例
	fmt.Println(runtime.NumCPU()) // 获取cpu的逻辑核心数
	cp := cpuPercent / float64(runtime.NumCPU())
	fmt.Println(cp)

	// 内存使用率, 线程数 goroutine 数
	mp, _ := p.MemoryPercent() // 后去进程占用内存的比例
	fmt.Println(mp)
	threadCount := pprof.Lookup("threadcreate").Count()// 创建的线程数
	fmt.Println(threadCount)
	gNum := runtime.NumGoroutine() // 启动的goroutine的数量
	fmt.Println(gNum)
	// 以上获取的指标, 只有再虚拟机和物理机环境下才能精准获取,类似docker容器这样的是依靠Linux的Namespace和Cgroup计数实现的隔离和资源限制的, 是不能通过以上方式获取数据的
}

/*
在Linux中，Cgroups给用户暴露出来的操作接口是文件系统，它以文件和目录的方式组织在操作系统的/sys/fs/cgroup路径下，
在 /sys/fs/cgroup下面有很多诸cpuset、cpu、 memory这样的子目录，每个子目录都代表系统当前可以被Cgroups进行限制的资源种类。

针对我们监控Go进程内存和CPU指标的需求，我们只要知道cpu.cfs_period_us、cpu.cfs_quota_us 和memory.limit_in_bytes 就行。
前两个参数需要组合使用，可以用来限制进程在长度为cfs_period的一段时间内，只能被分配到总量为cfs_quota的CPU时间，
可以简单的理解为容器能使用的核心数 = cfs_quota / cfs_period。
*/
func MonitorDocker() {
	// docker 容器中获取最大核心数
	pp, _ := process.NewProcess(int32(os.Getpid()))
	cpuPeriod, err := readUint("/sys/fs/cgroup/cpu/cpu.cfs_period_us")
	if err != nil {
		fmt.Println(err)
		return
	}
	cpuQuota, err := readUint("/sys/fs/cgroup/cpu/cpu.cfs_quota_us")
	cpuNum := float64(cpuQuota) / float64(cpuPeriod)
	cpuPercentAno, err := pp.Percent(time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	// cp := cpuPercent / float64(runtime.NumCPU())
	// 调整为
	cpAno := cpuPercentAno / cpuNum
	fmt.Println(cpAno)
	memLimit, err := readUint("/sys/fs/cgroup/memory/memory.limit_in_bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	memInformation := pp.MemoryInfo
	mem, err := memInformation()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mem.RSS * 100 / memLimit)
	mpAno := mem.RSS * 100 / memLimit
	fmt.Println(mpAno)
}

func readUint(path string) (uint64, error) {
	v, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return parseUint(strings.TrimSpace(string(v)), 10, 64)
}

func parseUint(s string, base, bitSize int) (uint64, error) {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		intValue, intErr := strconv.ParseInt(s, base, bitSize)
		// 1. Handle negative values greater than MinInt64 (and)
		// 2. Handle negative values lesser than MinInt64
		if intErr == nil && intValue < 0 {
			return 0, nil
		} else if intErr != nil &&
			intErr.(*strconv.NumError).Err == strconv.ErrRange &&
			intValue < 0 {
			return 0, nil
		}
		return 0, err
	}
	return v, nil
}