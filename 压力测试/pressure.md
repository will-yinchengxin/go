# [vegeta](https://github.com/tsenart/vegeta#report-command)

压力测试也是后端工作中比较多见的需求, 一般来说用得比较多的是：[ab - Apache HTTP server benchmarking tool](https://httpd.apache.org/docs/2.4/programs/ab.html) & [wg/wrk](https://github.com/wg/wrk)。这两款都是使用C语言进行编写的工具

存在的问题:

- 压测的目标不一定是单个的API
- 压测的时长和速率需要能调节
- 压测的结果要能快速展示，包括文字和graph两种形式
- 压测的结果要能进行存储，以使用其他软件进行拓展研究
- 压测要能以client端分布式的方式进行，以保证在制造巨量的请求时，可以使用多台client进行联合测试，且测试结果可以融合归一进行分析

这里就需要我们的 vegeta 登场了

## 1. 安装

Mac 安装

```shell
brew update && brew install vegeta
```

编译安装

```shell
git clone https://github.com/tsenart/vegeta
cd vegeta
make vegeta
mv vegeta ~/bin # Or elsewhere, up to you.
```

## 2. Api 准备

```go
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	g := gin.Default()
	gin.SetMode("release")
	g.GET("/get", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is gin test")
	})
	g.POST("/post", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Body)
		ctx.String(http.StatusOK, "this is gin /")
	})
	g.Handle("GET", "/get/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "test handle")
	})
	http.ListenAndServe(":8083", g)
}
```

## 3. 常用命令

vegeta命令包含多个功能，在运行的时候，需要额外提供一个子命令，来明确目前运行的时候具体要做什么：

- attack：压测命令
- encode：转码命令，将压测输出的结果集从默认格式转成其他的格式（json等）
- plot：graph命令，将压测输出的结果生成graph html，以供查看
- report：报告命令，将压测输出的结果生成文字形式的report

这里着重介绍一下 attack 和  report 命令

### 3.1 [`attack` command](https://github.com/tsenart/vegeta#attack-command)

- [`-body`](https://github.com/tsenart/vegeta#-body):  该option用来指定post body的数据文件，注意，如果在targets里也提供了`@/body/file/path`则以target里的为准，当前option提供的body会被忽略
- [`-max-connections`](https://github.com/tsenart/vegeta#-max-connections): 指定每个目标主机的最大连接数。
- [`-duration`](https://github.com/tsenart/vegeta#-duration): 该option用来控制测试时长，使用字符串形式来提供数值：`5s`
- [`-format`](https://github.com/tsenart/vegeta#-format): 指定要解码的目标格式, 该option一般使用`http`
- [`-header`](https://github.com/tsenart/vegeta#-header): 该option可以提供多个，每个option表示额外提供一个header
- [`-keepalive`](https://github.com/tsenart/vegeta#-keepalive): 指定是否在HTTP请求之间是否使用持久连接（默认为true）
- [`-output`](https://github.com/tsenart/vegeta#-output): 指定将向其中写入二进制结果的输出文件。使其通过管道传输到报表命令输入。默认为stdout。
- [`-rate`](https://github.com/tsenart/vegeta#-rate): 该option用来控制请求并发速率
  - `5/s`表示每秒vegeta会一共发出5个请求（vegeta会启动补充worker，如果速度设置很高的话）
  - 设置为0表示无限制，vegeta会尽可能快地发出请求
  - 设置为0的时候需要和`max-workers`配合使用，来行成一个固定的并发速率，否则可能会快速耗尽系统资源
- [`-targets`](https://github.com/tsenart/vegeta#-targets):  该option指定attack的目标是什么，默认读取stdin的输入，也可以提供一个文件路径，让vegeta读取其文本内容来获取目标
- [`-timeout`](https://github.com/tsenart/vegeta#-timeout): 该option表示每个http请求的超时时长，默认为0表示无超时

### 3.2 [`report` command](https://github.com/tsenart/vegeta#report-command)

#### [`report -type=text`](https://github.com/tsenart/vegeta#report--typetext)

```
[root@vrg-01 ~]#  vegeta attack \
        -targets=./targets.txt \
        -duration=5s \
        -rate=2000/s \
        -format=http | \
        vegeta report -type=text
Requests      [total, rate, throughput]         10000, 2000.18, 1984.99
Duration      [total, attack, wait]             5.038s, 5s, 38.246ms
Latencies     [min, mean, 50, 90, 95, 99, max]  2.987ms, 86.19ms, 49.213ms, 222.07ms, 279.693ms, 357.469ms, 449.597ms
Bytes In      [total, mean]                     570000, 57.00
Bytes Out     [total, mean]                     960000, 96.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:10000
Error Set:

------------------------------------------------------------------------------------------------

“请求”行显示：
	已发出的请求总数。
	攻击期间持续的实际请求率。
	总期间内成功请求的吞吐量。

“持续时间”行显示：
	发出所有请求所花费的攻击时间（总等待时间）
	等待对上次发出的请求的响应的等待时间（总攻击）
	攻击所花费的总时间（攻击+等待）

Latency是读取对请求的响应所花费的时间（包括响应正文中的-max正文字节）。
	min是攻击中所有请求的最小延迟。
	mean是攻击中所有请求延迟的算术平均值。
	50、90、95、99分别是攻击中所有请求的延迟的第50、第90、第95和第99个百分位数。为了进一步了解为什么这些有用，我推荐@tylertreat的这篇文章。
	max是攻击中所有请求的最大延迟。

Bytes In和Bytes Out行显示：
	与请求或响应主体一起发送（传出）或接收（传入）的字节总数。
	与请求或响应主体一起发送（输出）或接收（输入）的平均字节数。
	Success ratio显示响应未出错且状态代码在200到400之间（不包括200和400）的请求的百分比。


状态代码行显示状态代码的直方图。0状态代码表示请求发送失败。


“错误集”显示由所有发出的请求返回的唯一错误集。其中包括获得非成功响应状态代码的请求。
```

#### [`report -type=json`](https://github.com/tsenart/vegeta#report--typejson)

```json
[root@vrg-01 ~]#  vegeta attack \
        -targets=./targets.txt \
        -duration=5s \
        -rate=2000/s \
        -format=http | \
        vegeta report -type=json

{
  "latencies": {
    "total": 237119463,
    "mean": 2371194,
    "50th": 2854306,
    "90th": 3228223,
    "95th": 3478629,
    "99th": 3530000,
    "max": 3660505,
    "min": 1949582
  },
  "buckets": {
    "0": 9952,
    "1000000": 40,
    "2000000": 6,
    "3000000": 0,
    "4000000": 0,
    "5000000": 2
  },
  "bytes_in": {
    "total": 606700,
    "mean": 6067
  },
  "bytes_out": {
    "total": 0,
    "mean": 0
  },
  "earliest": "2015-09-19T14:45:50.645818631+02:00",
  "latest": "2015-09-19T14:45:51.635818575+02:00",
  "end": "2015-09-19T14:45:51.639325797+02:00",
  "duration": 989999944,
  "wait": 3507222,
  "requests": 100,
  "rate": 101.01010672380401,
  "throughput": 101.00012489812,
  "success": 1,
  "status_codes": {
    "200": 100
  },
  "errors": []
}
```

> 这里需要注意, 多个 command 之间要使用 `|`  符号来间隔开, 类似于管道符操作

## 4 小试牛刀

### 4.1 从标准输入中执行命令

只适用于 query string 类型的请求

```shell
[root@vrg-01 ~]# echo 'GET http://172.16.27.143:18899/v1/publish_mgmt2/list/isUpdate' | \
vegeta attack \
    -duration=5s \
    -rate=5/s \
    -format=http \
    -header 'Cache-Control: no-cache' | \
vegeta report -type=text

Requests      [total, rate, throughput]         25, 5.21, 5.18
Duration      [total, attack, wait]             4.825s, 4.8s, 24.548ms
Latencies     [min, mean, 50, 90, 95, 99, max]  3.075ms, 10.947ms, 10.748ms, 20.536ms, 24.292ms, 24.548ms, 24.548ms
Bytes In      [total, mean]                     1325, 53.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:25
Error Set:
```

通过管道符将请求的 url 传递给 vegeta, 每秒 5 个请求, 持续 5s, http 格式请求, 报告以 text 格式输出

### 4.2 从文件中行命令

当请求的 body 为  json/xml/html 或者其他文本类格式的时候, 就要使用这种请求

```shell
[root@vrg-01 ~]# ls
data.json   targets.txt
```

```shell
[root@vrg-01 ~]# cat targets.txt
POST http://172.16.27.50:18899/v1/client/platformStrategy
Content-Type: application/json
@./data.json
```

targets.txt 中也可以包含多个请求

```shell
GET http://localhost:5000/get?get_key=get_val
X-Add-Get-ID1: 78
X-Add-Get-ID2: 88

POST http://localhost:5000/post?post_key1=post_val
X-Add-Post-ID: 199
Content-Type: application/json
@./vegeta/postdata.json

POST http://localhost:5000/post?post_key2=post_val
X-Add-Post-ID: 299
Content-Type: application/json
@./vegeta/postdata.json
```

data 文件

```shell
[root@vrg-01 ~]# cat data.json
{
    "url":"http://origin.url.23.com",
    "ip":"172.16.27.95",
    "device_type":"HongMeng"
}
```

vegeta发送的请求会以targets.txt内列的target一个个请求发送下去, 每个target下面可以附带额外的header，不限数量，每行一个header, POST的target可以在下面带上一行`@/body/file/path`格式的文本，表示读取该位置的文件，作为post的body使用

```shell
[root@vrg-01 ~]#  vegeta attack \
        -targets=./targets.txt \
        -duration=5s \
        -rate=2000/s \
        -timeout=5s \
        -format=http | \
        vegeta report -type=text
        
Requests      [total, rate, throughput]         10000, 2000.18, 1984.99
Duration      [total, attack, wait]             5.038s, 5s, 38.246ms
Latencies     [min, mean, 50, 90, 95, 99, max]  2.987ms, 86.19ms, 49.213ms, 222.07ms, 279.693ms, 357.469ms, 449.597ms
Bytes In      [total, mean]                     570000, 57.00
Bytes Out     [total, mean]                     960000, 96.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:10000
Error Set:

```

## 5. vegeta plot

该命令用来将测试结果输出成graph的html，[官方文档](https://github.com/tsenart/vegeta#plot-command)。option基本不需要调整。这里需要注意，输出的graph只含`latency`，不含其它信息。

```shell
# 将加过输出到 bin 文件中
[root@vrg-01 ~]# echo 'GET http://172.16.27.143:18899/v1/publish_mgmt2/list/isUpdate' | \
vegeta attack \
    -duration=5s \
    -rate=5/s \
    -timeout=5s \
    -format=http \
    -header 'Cache-Control: no-cache'  > results.qps.bin
    
  
  
# 将 bin 文件转换为 html
[root@vrg-01 ~]# vegeta plot results.qps.bin > plot.html  


[root@vrg-01 ~]# ls
data.json     plot.html     results.qps.bin    targets.txt
```
