# Go Modules

## 1. 初期

Go语言在1.5之前没有依赖管理工具，若想引入依赖库，需要执行go get命令将代码拉取放入GOPATH/src目录下，作为GOPATH下的全局依赖，这也就意味着没有版本控制及隔离项目的包依赖

为了解决隔离项目的包依赖问题，Go1.5版本推出了vendor机制，环境变量中有一个GO15VENDOREXPERIMENT需要设置为1，该环境变量在Go1.6版本时变成默认开启，目前已经退出了历史舞台

vendor其实就是将原来放在GOPATH/src的依赖包放到工程的vendor目录中进行管理，不同工程独立地管理自己的依赖包，相互之间互不影响，原来是包共享的模式，通过vendor这种机制进行隔离，在项目编译的时候会先去vendor目录查找依赖，如果没有找到才会再去GOPATH目录下查找

- 优点：保证了功能项目的完整性，减少了下载依赖包，直接使用vendor就可以编译
- 缺点：仍然没有解决版本控制问题，go get仍然是拉取最新版本代码

## 2. Go Modules 闪亮登场

go modules是Russ Cox推出来的，发布于Go1.11，成长于Go1.12，丰富于Go1.13，正式于Go1.14推荐在生产上使用，几乎后续的每个版本都或多或少的有一些优化，在Go1.16引入go mod retract、在Go1.18引入go work工作区的概念

### 2.1 常用变量

#### 2.1.1 GO111MODULE环境变量

GO111MODULE 是 Go Modules的开关，主要有以下参数：

- auto：只在项目包含了go.mod文件时启动go modules，在Go1.13版本中是默认值
- on：无脑启动Go Modules，推荐设置，Go1.14版本以后的默认值
- off：禁用Go Modules，一般没有使用go modules的工程使用；

Go1.19.3中默认 GO111MODULE=on，也许未来的某一天 GO111MODULE 也会被取消, 默认置为 on

#### 2.1.2 GOPROXY

GOPROXY 用于设置Go模块代理，Go后续在拉取模块版本时能够脱离传统的VCS方式从镜像站点快速拉取，GOPROXY的值要以英文逗号分割，默认值是`https://proxy.golang.org,direct`，但是该地址在国内无法访问，所以可以使用`goproxy.cn`来代替(七牛云配置)，设置命令：

```
go env -w GOPROXY=GOPROXY=https://goproxy.cn,direct
```

也可以使用其他配置，例如阿里配置：

```
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
```

该环境变量也可以关闭，可以设置为"off"，禁止Go在后续操作中使用任何Go module proxy；

**关于 direct**

> `direct`为特殊指示符，因为我们指定了镜像地址，默认是从镜像站点拉取，但是有些库可能不存在镜像站点中，`direct`可以指示Go回源到模块版本的源地址去抓取，比如 github，当go module proxy返回404、410这类错误时，其会自动尝试列表中的下一个，遇见`direct`时回源地址抓取

#### 2.1.3 GOSUMDB

GOSUMDB 是一个 Go checksum database，用于保证Go在拉取模块版本时拉取到的模块版本数据未经篡改，若发现不一致会中止，也可以将值设置为`off`即可以禁止Go在后续操作中校验模块版本

**什么是Go checksum database?**

[Go checksum database主要用于保护Go不会从任何拉到被篡改过的非法Go模块版本](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)

GOSUMDB的默认值是`sum.golang.org`，默认值与自定义值的格式不一样，默认值在国内是无法访问，这个值我们一般不用动，因为我们一般已经设置好了GOPROXY，goproxy.cn支持代理`sum.golang.org`;

GOSUMDB的值自定义格式如下：

- 格式 1：`<SUMDB_NAME>+<PUBLIC_KEY>`
- 格式 2：`<SUMDB_NAME>+<PUBLIC_KEY> <SUMDB_URL>`

#### 2.1.4 GONOPROXY / GONOSUMDB / GOPRIVATE

这三个变量一般在项目中很少使用, 主要用于私有模块的拉取，在GOPROXY、GOSUMDB中无法访问到模块的场景中，例如拉取git上的私有仓库

GONOPROXY、GONOSUMDB的默认值是GOPRIVATE的值，所以我们一般直接使用GOPRIVATE即可，其值也是可以设置多个，以英文逗号进行分割  `；` 例如

```
go env -w GOPRIVATE="github.com/will-yinchengxin/****,gitlab.xxxx.com"
```

也可以使用通配符的方式进行设置，对域名设置通配符号，这样子域名就都不经过 Go module proxy 和 Go checksum database;

#### 2.1.5 全局缓存 

`go mod download / tidy` 会将依赖缓存到本地，缓存的目录是`GOPATH/pkg/mod/cache`  或者 `GOPATH/pkg/sum`，这些缓存依赖可以被多个项目使用，未来可能会迁移到`$GOCACHE`下面；

可以使用 `go clean -modcache` 清理所有已缓存的模块版本数据；

### 2.2 Go Modules命令

```shell
WillYin: willyin$ go mod help
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

```
命令							 作用
go mod init				生成go.mod文件
go mod download		下载go.mod文件中指明的所有依赖放到全局缓存
go mod tidy				整理现有的依赖，添加缺失或移除不使用的modules
go mod graph			查看现有的依赖结构
go mod edit				编辑 go.mod 文件
go mod vendor			导出项目所有的依赖到vendor目录
go mod verify			校验一个模块是否被篡改过
go mod why				解释为什么需要依赖某个模块
```

**go.mod 文件**

go.mod是启用Go modules的项目所必须且最重要的文件，其描述了当前项目的元信息，每个go.mod文件开头符合包含如下信息：

```go
module willTest

go 1.18

require (
	github.com/astaxie/beego v1.12.2
	github.com/go-sql-driver/mysql v1.5.0
)

exclude (
	github.com/willchengxin/gotest
)

replace (
	golang.org/x/image@v0.0.0-20180708004352-c73c2afc3b81 => github.com/golang/image@v0.0.0-20190101..
)

retract v0.2.0
```

- **module**：用于定义当前项目的模块路径（突破 `$GOPATH` 路径）

  ```
  go.mod文件的第一行是module path，采用仓库 + module name 的方式定义
  
  例如: module willTest
  ```

  ```
  go module 拉取依赖包本质也是 go get 行为，go get 主要提供了以下命令：
  
  命令									 作用
  go get								拉取依赖，会进行指定性拉取（更新），并不会更新所依赖的其它模块。
  go get -u							更新现有的依赖，会强制更新它所依赖的其它全部模块，不包括自身。
  go get -u -t ./...		更新所有直接依赖和间接依赖的模块版本，包括单元测试中用到的。
  ```

- **go**：当前项目Go版本，目前只是标识作用

  ```
  go.mod文件的第二行是go version，其是用来指定你的代码所需要的最低版本：
  
  例如: go 1.19.3
  
  这一行不是必须的，目前也只是标识作用，可以不写
  ```

- **require**：用设置一个特定的模块版本

  ```
  require用来指定该项目所需要的各个依赖库以及他们的版本, 有些时候我们还会看到注释, 那注释是用来做什么的呢?
  
  例子:	
  	google.golang.org/appengine v1.6.5 // indirect
  	github.com/dgrijalva/jwt-go v3.2.0+incompatible
  ```

  - **indirect 注释**

    ```
    以下场景才会添加indirect注释：
    
    - 当前项目依赖包A，但是A依赖包B，但是A的go.mod文件中缺失B，所以在当前项目go.mod中补充B并添加indirect注释
    - 当前项目依赖包A，但是依赖包A没有go.mod文件，所以在当前项目go.mod中补充B并添加indirect注释
    - 当前项目依赖包A，依赖包A又依赖包B，当依赖包A降级不在依赖B时，这个时候就会标记indirect注释，可以执行 go mod tidy 移除该依赖
    
    Go1.17版本对此做了优化，indirect 的 module 将被放在单独 require 块的，这样看起来更加清晰明了。
    ```

  - **incompatible 标记**

    ```
    jwt-go这个库就是这样的，这是因为jwt-go的版本已经大于2了，但是他们的module path仍然没有添加v2、v3这样的后缀，不符合Go的 module 管理规范，所以 go module 把他们标记为 incompatible，不影响引用
    ```

- **exclude**：用于从使用中排除一个特定的模块版本

  ```
  用于跳过某个依赖库的版本，使用场景一般是我们知道某个版本有bug或者不兼容，为了安全起可以使用exclude跳过此版本；
  
  exclude (
   go.etcd.io/etcd/client/v2 v2.305.0-rc.0
  )
  ```

- **replace**：用于将一个模块版本替换为另外一个模块版本，例如chromedp使用 `golang.org/x/image` 这个package一般直连是获取不了的，但是它有一个 `github.com/golang/image` 的镜像，所以我们要用replace来用镜像替换它

  ```
  开发中我们常使用第三方库，大部分场景都可以满足我们的需要，但是有些时候我们需要对依赖库做一些定制修改，依赖库修改后，我们想引起最小的改动，就可以使用 replace 命令进行重新引用，调试也可以使用 replace 进行替换
  ```

- **restract**：用来声明该第三方模块的某些发行版本不能被其他模块使用，在 Go1.16 引入

**go.sum 文件**

go.sun文件也是在go mod init阶段创建, go.sum主要是记录了所有依赖的module的校验信息, 组成大致可以理解成 `module + version + hash`

```
github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 h1:Hs82Z41s6SdL1CELW+XaDYmOH4hkBN4/N9og/AsOv7E=
cloud.google.com/go v0.38.0/go.mod h1:990N+gfupTy94rShfmMCWGDn0LpTmnzTp2qbd1dvSRU=
// ......
```

主要有两种形式

- h1:
- /go.mod h1:

```
其中 module 是依赖的路径，version 是依赖的版本号。hash 是以h1:开头的字符串，hash 是 Go modules 将目标模块版本的 zip 文件开包后，针对所有包内文件依次进行 hash，然后再把它们的 hash 结果按照固定格式和算法组成总的 hash 值。

h1 hash 和 go.mod hash 两者要不同时存在，要不就是只存在 go.mod hash，当 Go 认为肯定用不到某个版本的时候就会省略它的 h1 hash，就只有 go.mod hash；
```



