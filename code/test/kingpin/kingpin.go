package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/alertmanager/notify/webhook"
	"github.com/prometheus/alertmanager/template"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

//go:generate sh -c "echo will"

func main() {
	// Base returns the last element of path.
	fmt.Println(filepath.Base("/usr/local/go/")) // go
	// 必须为 IP 格式
	ip := kingpin.Flag("ip", "IP addr").IP()

	// 必传
	pwd := kingpin.Flag("pwd", "Password").String()

	// Int 类型
	hostID := kingpin.Flag("hostID", "HostID").Int()

	// time
	time := kingpin.Flag("time", "Time").Duration()

	// 解析
	/*
		方法清单
			输出信息
				kingpin.Version(): 输出版本信息
				kingpin.FatalIfError(): 如果有报错, 打印错误信息, 并退出
				kingpin.Fatalf(): 打印错误信息, 并退出
				kingpin.Errorf(): 打印错误信息, 不退出
				kingpin.FatalUsage(): 如果有报错, 打印帮助信息
				kingpin.Usage(): 打印帮助信息
			创建参数
				kingpin.Arg(): 创建固定参数(按顺序传入, 不需要 --flag 指定)
				kingpin.Flag(): 创建可选参数
				解析参数
				kingpin.Parse(): 用法同 flag
				kingpin.MustParse(): Parse() 底层调用的它
			其他:
				kingpin.New()
				kingpin.ExpandArgsFromFile()
				kingpin.UsageTemplate()
				kingpin.Command()
				kingpin.HelpFlag.Short('h'): 启动 -h

		接收类型
			按接收方式分(以string类型为例)
				kingpin.Flag().String(): 直接指针接收
				kingpin.Flag().StringVar(): 先创建变量, 用该变量指针接收
			按类型分(不同类型有不同的方法, 以string为例)
				kingpin.Flag().String()
				kingpin.Flag().Strings(): 以 []string 接收, 接收值为多个时, 必选参数只能有一个, 否则无法区分
				kingpin.Flag().StringMap(): 以 map[string]string 类型接收
				其他类型可能没有 map, 如 Bool() 和 BoolList()

		限制
			kingpin.Flag().Required().String(): 必传
			kingpin.Flag().IP(): ip 格式
			kingpin.Flag().Duration(): 时间格式, 10s, 2m, 3h
			kingpin.Flag().Short(): 设置短参数
			kingpin.Flag().Default(): 设置默认值
			kinpin.Flag().Envar(): 使用环境变量
	*/
	kingpin.Parse()

	fmt.Println(*ip, *pwd, *hostID, *time)
	/*
		MacBook-Pro:test yinchengxin$ go run main.go --ip="172.16.27.99" --pwd="/usr/local" --hostID="123" --time="12m"
		172.16.27.99 /usr/local 123 12m0s
	*/
}

/*
func ruleEngineMain() {

	cfg := ruleEngineModules.Config{
		PromlogConfig: promlog.Config{},
	}

	// filepath.Base: Base returns the last element of path.
	// kingpin: 类似 flag parse
	a := kingpin.New(filepath.Base(os.Args[0]), "Rule Engine")

	a.HelpFlag.Short('h')
	logs.Error("start rule engine.......")

	a.Flag("gateway.url", "alert gateway url").
		Default("http://localhost:32000").StringVar(&cfg.GatewayURL)
	a.Flag("gateway.path.rule", "alert gateway rule url").
		Default("/api/v1/rules").StringVar(&cfg.GatewayPathRule)
	a.Flag("gateway.path.prom", "alert gateway prom url").
		Default("/api/v1/proms").StringVar(&cfg.GatewayPathProm)
	a.Flag("gateway.path.notify", "alert gateway notify url").
		Default("/api/v1/alerts").StringVar(&cfg.GatewayPathNotify)

	a.Flag("notify.retries", "notify retries").
		Default("3").IntVar(&cfg.NotifyReties)
	a.Flag("evaluation.interval", "rule evaluation interval").
		Default("15s").SetValue(&cfg.EvaluationInterval)
	a.Flag("reload.interval", "rule reload interval").
		Default("5m").SetValue(&cfg.ReloadInterval)
	a.Flag("auth.token", "http auth token").
		Default("96smhbNpRguoJOCEKNrMqQ").StringVar(&cfg.AuthToken)
}		
*/
